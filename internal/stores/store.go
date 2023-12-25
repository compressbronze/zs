package store

import "github.com/satrap-illustrations/zs/internal/models"

type Store interface {
	ListFields() (map[string][]string, error)
	Search(documentType, field, query string) ([]models.Model, error)
}
