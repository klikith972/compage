package generator

import (
	"errors"
	"github.com/kube-tarian/compage-core/internal/core"
	"github.com/kube-tarian/compage-core/internal/languages"
	"github.com/kube-tarian/compage-core/internal/languages/golang"
	"github.com/kube-tarian/compage-core/internal/utils"
	log "github.com/sirupsen/logrus"
)

// Generator called from rest as well as gRPC
func Generator(coreProject *core.Project) error {
	// create a directory with project name to contain code generated by compage-core.
	projectDirectory := utils.GetProjectDirectoryName(coreProject.Name)
	if err := utils.CreateDirectories(projectDirectory); err != nil {
		return err
	}

	// Iterate over all nodes and generate code for all nodes.
	compageYaml := coreProject.CompageYaml
	for _, compageNode := range compageYaml.Nodes {
		log.Info("processing node ID : ", compageNode.ID)
		// if language is not set, consider that the node is go project
		if compageNode.ConsumerData.Language == "" || compageNode.ConsumerData.Language == languages.Go {
			goNode, err := golang.NewNode(compageYaml, compageNode)
			if err != nil {
				// return errors like certain protocols aren't yet supported
				return err
			}
			if err := golang.Generator(coreProject.Name, goNode); err != nil {
				return err
			}
			// trigger template runner
			// TODO
		} else if compageNode.ConsumerData.Language == languages.NodeJs {
			return errors.New("unsupported language : " + languages.NodeJs)
		} else {
			return errors.New("unsupported language : " + compageNode.ConsumerData.Language)
		}
	}

	return nil
}
