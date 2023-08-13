package handler

import (
	"github.com/khulnasoft-lab/pipeline-parser/pkg/consts"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/enhancers"
	azureEnhancer "github.com/khulnasoft-lab/pipeline-parser/pkg/enhancers/azure"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/loaders"
	azureLoader "github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/azure"
	azureModels "github.com/khulnasoft-lab/pipeline-parser/pkg/loaders/azure/models"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/models"
	"github.com/khulnasoft-lab/pipeline-parser/pkg/parsers"
	azureParser "github.com/khulnasoft-lab/pipeline-parser/pkg/parsers/azure"
)

type AzureHandler struct{}

func (g *AzureHandler) GetPlatform() models.Platform {
	return consts.AzurePlatform
}

func (g *AzureHandler) GetLoader() loaders.Loader[azureModels.Pipeline] {
	return &azureLoader.AzureLoader{}
}

func (g *AzureHandler) GetParser() parsers.Parser[azureModels.Pipeline] {
	return &azureParser.AzureParser{}
}

func (g *AzureHandler) GetEnhancer() enhancers.Enhancer {
	return &azureEnhancer.AzureEnhancer{}
}
