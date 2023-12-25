package inverted

import (
	"fmt"
	"strconv"

	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/tokeniser"
)

var ErrInvalidField = fmt.Errorf("invalid field for organization store")

type OrganizationStore struct {
	models map[int]models.Organization
	// we aren't counting multiplicity at the moment, we could use it to rank the results.
	index map[tokeniser.Token][]int
}

func NewOrganizationStore(organizations []models.Organization) OrganizationStore {
	s := OrganizationStore{
		models: map[int]models.Organization{},
		index:  map[tokeniser.Token][]int{},
	}
	for i, organization := range organizations {
		s.models[organization.ID] = organization
		for _, token := range tokeniser.Tokenise(&organizations[i]) {
			if s.index[token] == nil {
				s.index[token] = []int{organization.ID}
			} else {
				s.index[token] = append(s.index[token], organization.ID)
			}
		}
	}
	return s
}

func (OrganizationStore) ListFields() []string {
	return models.FieldSlice(new(models.Organization))
}

func (s OrganizationStore) Search(field, query string) ([]models.Organization, error) {
	o := new(models.Organization)
	if _, exists := o.Fields().Get(field); !exists {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == o.StringID() {
		id, err := strconv.Atoi(query)
		if err != nil {
			return nil, err
		}
		if organization, ok := s.models[id]; ok {
			return []models.Organization{organization}, nil
		}
		return []models.Organization{}, nil
	}

	out := []models.Organization{}
	token := tokeniser.Token{
		Text:  query,
		Field: field,
	}
	if indices, exist := s.index[token]; exist {
		for _, i := range indices {
			out = append(out, s.models[i])
		}
	}

	return out, nil
}
