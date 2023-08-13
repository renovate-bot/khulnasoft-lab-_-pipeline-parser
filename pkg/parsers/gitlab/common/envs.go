package common

import (
	gitlabModels "github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/gitlab/models/common"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
)

func ParseEnvironmentVariables(environmentVariables *gitlabModels.EnvironmentVariablesRef) *models.EnvironmentVariablesRef {
	if environmentVariables == nil {
		return nil
	}

	return &models.EnvironmentVariablesRef{
		FileReference:        environmentVariables.FileReference,
		EnvironmentVariables: (map[string]any)(*environmentVariables.Variables),
	}
}
