package stores

import "github.com/satrap-illustrations/zs/internal/models"

type Store interface {
	ListDocumentTypes() []string
	ListFields() map[string][]string
	Search(documentType, field, query string) ([]models.Model, error)
}
