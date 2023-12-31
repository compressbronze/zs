package hash

import (
	"fmt"
	"strconv"

	"github.com/satrap-illustrations/zs/internal/models"
)

var ErrInvalidField = fmt.Errorf("invalid field for organization store")

type OrganizationStore map[int]models.Organization

func NewOrganizationStore(organizations []models.Organization) OrganizationStore {
	out := OrganizationStore{}
	for _, organization := range organizations {
		out[organization.ID] = organization
	}
	return out
}

func (OrganizationStore) ListFields() []string {
	return models.FieldSlice(new(models.Organization))
}

func (s OrganizationStore) Search(field, query string) ([]models.Organization, error) {
	i, exists := new(models.Organization).Fields().Get(field)
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == "_id" {
		id, err := strconv.Atoi(query)
		if err != nil {
			return nil, err
		}
		if organization, ok := s[id]; ok {
			return []models.Organization{organization}, nil
		}
		return []models.Organization{}, nil
	}

	out := []models.Organization{}
	for _, organization := range s {
		if models.ValueContains(organization.ValueAtIdx(i), query) {
			out = append(out, organization)
		}
	}

	return out, nil
}
