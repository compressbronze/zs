package stores_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/stores"
	"github.com/satrap-illustrations/zs/internal/stores/implementations"
	"gotest.tools/v3/assert"
)

func TestStores(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		docType  string
		field    string
		query    string
		expected []models.Model
	}{
		{
			name:     "organization_id_118",
			docType:  "Organizations",
			field:    "_id",
			query:    "118",
			expected: organization118Results,
		},
		{
			name:     "organization_name_Limozen",
			docType:  "Organizations",
			field:    "name",
			query:    "Limozen",
			expected: organization118Results,
		},
		{
			name:    "user_organization_id_118",
			docType: "Users",
			field:   "organization_id",
			query:   "118",
			expected: []models.Model{
				&models.User{
					ID:             59,
					URL:            "http://initech.zendesk.com/api/v2/users/59.json",
					ExternalID:     uuid.Must(uuid.Parse("4acd4eb0-9168-4270-b09f-09600a05b0b2")),
					Name:           "Key Mendez",
					Alias:          "Mr Lucile",
					CreatedAt:      "2016-04-23T12:00:11 -10:00",
					Locale:         "zh-CN",
					Timezone:       "Nigeria",
					LastLoginAt:    "2014-06-03T02:26:28 -10:00",
					Email:          "lucilemendez@flotonic.com",
					Phone:          "8774-883-991",
					Signature:      "Don't Worry Be Happy!",
					OrganizationID: 118,
					Tags:           []string{"Rockingham", "Waikele", "Masthope", "Oceola"},
					Role:           "agent",
				},
				&models.User{
					ID:             49,
					URL:            "http://initech.zendesk.com/api/v2/users/49.json",
					ExternalID:     uuid.Must(uuid.Parse("4bd5e757-c0cd-445b-b702-ee3ed794f6c4")),
					Name:           "Faulkner Holcomb",
					Alias:          "Miss Jody",
					CreatedAt:      "2016-05-12T08:39:30 -10:00",
					Active:         true,
					Shared:         true,
					Locale:         "zh-CN",
					Timezone:       "Antigua and Barbuda",
					LastLoginAt:    "2014-12-04T12:51:36 -11:00",
					Email:          "jodyholcomb@flotonic.com",
					Phone:          "9255-943-719",
					Signature:      "Don't Worry Be Happy!",
					OrganizationID: 118,
					Tags:           []string{"Hanover", "Woodlake", "Saticoy", "Hinsdale"},
					Suspended:      true,
					Role:           "end-user",
				},
			},
		},
		{
			name:    "user_timezone_Antigua_and_Barbuda",
			docType: "Users",
			field:   "timezone",
			query:   "Antigua and Barbuda",
			expected: []models.Model{
				&models.User{
					ID:             49,
					URL:            "http://initech.zendesk.com/api/v2/users/49.json",
					ExternalID:     uuid.Must(uuid.Parse("4bd5e757-c0cd-445b-b702-ee3ed794f6c4")),
					Name:           "Faulkner Holcomb",
					Alias:          "Miss Jody",
					CreatedAt:      "2016-05-12T08:39:30 -10:00",
					Active:         true,
					Shared:         true,
					Locale:         "zh-CN",
					Timezone:       "Antigua and Barbuda",
					LastLoginAt:    "2014-12-04T12:51:36 -11:00",
					Email:          "jodyholcomb@flotonic.com",
					Phone:          "9255-943-719",
					Signature:      "Don't Worry Be Happy!",
					OrganizationID: 118,
					Tags:           []string{"Hanover", "Woodlake", "Saticoy", "Hinsdale"},
					Suspended:      true,
					Role:           "end-user",
				},
			},
		},
	} {
		tc := tc

		hashStore, err := implementations.NewHashStore("../../data")
		assert.NilError(t, err)

		invStore, err := implementations.NewInvertedStore("../../data")
		assert.NilError(t, err)

		for _, ts := range []struct {
			name  string
			store stores.Store
		}{
			{name: "HashStore", store: hashStore},
			{name: "InvertedStore", store: invStore},
		} {
			ts := ts

			t.Run(ts.name, func(t *testing.T) {
				t.Parallel()

				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					foundModels, err := ts.store.Search(tc.docType, tc.field, tc.query)
					assert.NilError(t, err)
					sortFunc := func(a, b models.Model) int {
						return cmp.Compare(a.DocumentType()+a.StringID(), b.DocumentType()+b.StringID())
					}
					slices.SortFunc(foundModels, sortFunc)
					slices.SortFunc(tc.expected, sortFunc)
					assert.DeepEqual(t, tc.expected, foundModels)
				})
			})
		}
	}
}
