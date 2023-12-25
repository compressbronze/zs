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
			name:     "user_organization_id_118",
			docType:  "Users",
			field:    "organization_id",
			query:    "118",
			expected: usersInOrg118Results,
		},
		{
			name:     "user_timezone_Antigua",
			docType:  "Users",
			field:    "timezone",
			query:    "Antigua",
			expected: userTimezoneAntiguaResults,
		},
		{
			name:     "tickets_subject_Latvia",
			docType:  "Tickets",
			field:    "subject",
			query:    "Latvia",
			expected: ticketsSubjectLatviaResults,
		},
	} {
		tc := tc

		const dataDir = "../../data"

		//nolint:staticcheck
		hashStore, err := implementations.NewHashStore(dataDir)
		assert.NilError(t, err)

		invStore, err := implementations.NewInvertedStore(dataDir)
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

					if ts.name == "HashStore" {
						t.Skip("Deprecated. This tests allows checking compatibility with the new implementation.")
					}

					foundModels, err := ts.store.Search(tc.docType, tc.field, tc.query)
					assert.NilError(t, err)
					slices.SortFunc(foundModels, sortFunc)
					slices.SortFunc(tc.expected, sortFunc)
					assert.DeepEqual(t, tc.expected, foundModels)
				})
			})
		}
	}
}

func TestEmptyFieldsAreSerchable(t *testing.T) {
	t.Parallel()

	const dataDir = "../../test/fixtures/empty_fields"

	store, err := implementations.NewInvertedStore(dataDir)
	assert.NilError(t, err)

	foundModels, err := store.Search("Tickets", "description", "")
	assert.NilError(t, err)

	expected := []models.Model{
		&models.Ticket{
			ID:             uuid.Must(uuid.Parse("436bf9b0-1147-4c0a-8439-6f79833bff5b")),
			URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			ExternalID:     uuid.Must(uuid.Parse("9210cdc9-4bee-485f-a078-35396cd74063")),
			CreatedAt:      "2016-04-28T11:19:34 -10:00",
			Type:           "incident",
			Subject:        "A Catastrophe in Korea (North)",
			Priority:       "high",
			Status:         "pending",
			SubmitterID:    38,
			AssigneeID:     24,
			OrganizationID: 116,
			Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
			DueAt:          "2016-07-31T02:37:50 -10:00",
			Via:            "web",
		},
		&models.Ticket{
			ID:             uuid.Must(uuid.Parse("4cce7415-ef12-42b6-b7b5-fb00e24f9cc1")),
			URL:            "http://initech.zendesk.com/api/v2/tickets/4cce7415-ef12-42b6-b7b5-fb00e24f9cc1.json",
			ExternalID:     uuid.Must(uuid.Parse("ef665694-aa3f-4960-b264-0e77c50486cf")),
			CreatedAt:      "2016-02-25T09:12:47 -11:00",
			Type:           "question",
			Subject:        "A Nuisance in Ghana",
			Priority:       "high",
			Status:         "solved",
			SubmitterID:    9,
			AssigneeID:     48,
			OrganizationID: 104,
			Tags:           []string{"Delaware", "New Hampshire", "Utah", "Hawaii"},
			DueAt:          "2016-08-05T10:31:03 -10:00",
			Via:            "web",
		},
		&models.Ticket{
			ID:             uuid.Must(uuid.Parse("87db32c5-76a3-4069-954c-7d59c6c21de0")),
			URL:            "http://initech.zendesk.com/api/v2/tickets/87db32c5-76a3-4069-954c-7d59c6c21de0.json",
			ExternalID:     uuid.Must(uuid.Parse("1c61056c-a5ad-478a-9fd6-38889c3cd728")),
			CreatedAt:      "2016-07-06T11:16:50 -10:00",
			Type:           "problem",
			Subject:        "A Problem in Morocco",
			Priority:       "urgent",
			Status:         "solved",
			SubmitterID:    14,
			AssigneeID:     7,
			OrganizationID: 118,
			Tags:           []string{"Texas", "Nevada", "Oregon", "Arizona"},
			HasIncidents:   true,
			DueAt:          "2016-08-19T07:40:17 -10:00",
			Via:            "voice",
		},
	}

	slices.SortFunc(foundModels, sortFunc)
	slices.SortFunc(expected, sortFunc)

	assert.DeepEqual(t, expected, foundModels)
}

func sortFunc(a, b models.Model) int {
	return cmp.Compare(a.DocumentType()+a.StringID(), b.DocumentType()+b.StringID())
}
