package implementations_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/stores/implementations"
	"gotest.tools/v3/assert"
)

func TestOrganizations(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		expected []models.Model
	}{
		{
			name: "organization_id_101",
			expected: []models.Model{
				&models.Organization{
					ID:          101,
					URL:         "http://initech.zendesk.com/api/v2/organizations/101.json",
					ExternalID:  uuid.Must(uuid.Parse("9270ed79-35eb-4a38-a46f-35725197ea8d")),
					Name:        "Enthaze",
					DomainNames: []string{"kage.com", "ecratic.com", "endipin.com", "zentix.com"},
					CreatedAt:   "2016-05-21T11:10:28 -10:00",
					Details:     "MegaCorp",
					Tags:        []string{"Fulton", "West", "Rodriguez", "Farley"},
				},
			},
		},
		{
			name: "organization_name_Enthaze",
			expected: []models.Model{
				&models.Organization{
					ID:          101,
					URL:         "http://initech.zendesk.com/api/v2/organizations/101.json",
					ExternalID:  uuid.Must(uuid.Parse("9270ed79-35eb-4a38-a46f-35725197ea8d")),
					Name:        "Enthaze",
					DomainNames: []string{"kage.com", "ecratic.com", "endipin.com", "zentix.com"},
					CreatedAt:   "2016-05-21T11:10:28 -10:00",
					Details:     "MegaCorp",
					Tags:        []string{"Fulton", "West", "Rodriguez", "Farley"},
				},
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			store, err := implementations.NewHashStore("../../../data")
			assert.NilError(t, err)

			models, err := store.Search("Organizations", "name", "Enthaze")
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, models)
		})
	}
}
