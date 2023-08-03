package utils

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/khulnasoft-labs/pipeline-parser/pkg/models"
)

// ToDo: change to SetCommonBearerAuthToken and test
func GetHttpClient(credentials *models.Credentials) *req.Client {
	client := req.C()
	if credentials == nil {
		return client
	}

	return client.SetCommonHeader("Authorization", fmt.Sprintf("token %s", credentials.Token))
}

func GetHttpClientWithBasicAuth(credentials *models.Credentials) *req.Client {
	client := req.C()
	if credentials == nil {
		return client
	}

	return client.SetCommonBasicAuth("", credentials.Token)
}
