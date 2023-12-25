package organizationstore

import "github.com/satrap-illustrations/zs/internal/models"

type OrganizationStore interface {
	ListFields() []string
	Search(field, query string) ([]models.Organization, error)
}
