package parsers

import "github.com/khulnasoft-labs/pipeline-parser/pkg/models"

type Parser[T any] interface {
	Parse(*T) (*models.Pipeline, error)
}
