package azure

import (
	azureModels "github.com/khulnasoft-labs/pipeline-parser/pkg/loaders/azure/models"
	"github.com/khulnasoft-labs/pipeline-parser/pkg/models"
)

func parseEnvironmentVariablesRef(envRef *azureModels.EnvironmentVariablesRef) *models.EnvironmentVariablesRef {
	if envRef == nil {
		return nil
	}

	return &models.EnvironmentVariablesRef{
		EnvironmentVariables: envRef.EnvironmentVariables,
		FileReference:        envRef.FileReference,
	}
}
