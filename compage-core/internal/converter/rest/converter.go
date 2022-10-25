package rest

import (
	"github.com/kube-tarian/compage-core/internal/converter"
	"github.com/kube-tarian/compage-core/internal/core"
)

func GetProject(input core.GenerateProjectInput) (*core.Project, error) {
	compageYaml, err := converter.GetCompageYaml(input.Yaml)
	if err != nil {
		return nil, err
	}
	return &core.Project{
		CompageYaml: compageYaml,
		Name:        input.Name,
		Repository:  input.Repository,
	}, nil
}
