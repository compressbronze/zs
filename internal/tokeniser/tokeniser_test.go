package tokeniser_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/tokeniser"
	"gotest.tools/v3/assert"
)

func TestTokenise(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name           string
		model          models.Model
		expectedTokens []tokeniser.Token
	}{
		{
			name:  "Organization_zero",
			model: &models.Organization{},
			expectedTokens: []tokeniser.Token{
				{Text: "0", Field: "_id"},
				{Text: "false", Field: "shared_tickets"},
			},
		},
		{
			name: "Organization",
			model: &models.Organization{
				ID:          118,
				URL:         "http://initech.zendesk.com/api/v2/organizations/118.json",
				ExternalID:  uuid.Must(uuid.Parse("6970300e-f211-4c01-a538-70b4464a1d84")),
				Name:        "Limozen",
				DomainNames: []string{"otherway.com", "rodeomad.com", "suremax.com", "fishland.com"},
				CreatedAt:   "2016-02-11T04:24:09 -11:00",
				Details:     "MegaCorp",
				Tags:        []string{"Leon", "Ferguson", "Olsen", "Walsh"},
			},
			expectedTokens: []tokeniser.Token{
				{Text: "-11:00", Field: "created_at"},
				{Text: "118", Field: "_id"},
				{Text: "2016-02-11T04:24:09", Field: "created_at"},
				{Text: "6970300e-f211-4c01-a538-70b4464a1d84", Field: "external_id"},
				{Text: "Ferguson", Field: "tags"},
				{Text: "Leon", Field: "tags"},
				{Text: "Limozen", Field: "name"},
				{Text: "MegaCorp", Field: "details"},
				{Text: "Olsen", Field: "tags"},
				{Text: "Walsh", Field: "tags"},
				{Text: "false", Field: "shared_tickets"},
				{Text: "fishland.com", Field: "domain_names"},
				{Text: "http://initech.zendesk.com/api/v2/organizations/118.json", Field: "url"},
				{Text: "otherway.com", Field: "domain_names"},
				{Text: "rodeomad.com", Field: "domain_names"},
				{Text: "suremax.com", Field: "domain_names"},
			},
		},
		//nolint: dupword
		{
			name: "Ticket",
			model: &models.Ticket{
				ID:             uuid.Must(uuid.Parse("0ebe753c-9c78-458a-817f-3993780bedbf")),
				URL:            "http://initech.zendesk.com/api/v2/tickets/0ebe753c-9c78-458a-817f-3993780bedbf.json",
				ExternalID:     uuid.Must(uuid.Parse("537ad752-9056-42c9-86db-f0bdf06d3c10")),
				CreatedAt:      "2016-05-19T12:19:56 -10:00",
				Type:           "problem",
				Subject:        "A Nuisance in Seychelles",
				Description:    "Consequat enim velit magna ad sit. Lorem mollit proident est id aliqua ea ea est aliquip magna.",
				Priority:       "high",
				Status:         "pending",
				SubmitterID:    23,
				AssigneeID:     56,
				OrganizationID: 118,
				Tags:           []string{"Missouri", "Alabama", "Virginia", "Virgin Islands"},
				HasIncidents:   true,
				DueAt:          "2016-08-18T03:33:30 -10:00",
				Via:            "chat",
			},
			expectedTokens: []tokeniser.Token{
				{Text: "-10:00", Field: "created_at"},
				{Text: "-10:00", Field: "due_at"},
				{Text: "0ebe753c-9c78-458a-817f-3993780bedbf", Field: "_id"},
				{Text: "118", Field: "organization_id"},
				{Text: "2016-05-19T12:19:56", Field: "created_at"},
				{Text: "2016-08-18T03:33:30", Field: "due_at"},
				{Text: "23", Field: "submitter_id"},
				{Text: "537ad752-9056-42c9-86db-f0bdf06d3c10", Field: "external_id"},
				{Text: "56", Field: "assignee_id"},
				{Text: "A", Field: "subject"},
				{Text: "Alabama", Field: "tags"},
				{Text: "Consequat", Field: "description"},
				{Text: "Islands", Field: "tags"},
				{Text: "Lorem", Field: "description"},
				{Text: "Missouri", Field: "tags"},
				{Text: "Nuisance", Field: "subject"},
				{Text: "Seychelles", Field: "subject"},
				{Text: "Virgin", Field: "tags"},
				{Text: "Virginia", Field: "tags"},
				{Text: "ad", Field: "description"},
				{Text: "aliqua", Field: "description"},
				{Text: "aliquip", Field: "description"},
				{Text: "chat", Field: "via"},
				{Text: "ea", Field: "description"},
				{Text: "ea", Field: "description"},
				{Text: "enim", Field: "description"},
				{Text: "est", Field: "description"},
				{Text: "est", Field: "description"},
				{Text: "high", Field: "priority"},
				{
					Text:  "http://initech.zendesk.com/api/v2/tickets/0ebe753c-9c78-458a-817f-3993780bedbf.json",
					Field: "url",
				},
				{Text: "id", Field: "description"},
				{Text: "in", Field: "subject"},
				{Text: "magna", Field: "description"},
				{Text: "magna", Field: "description"},
				{Text: "mollit", Field: "description"},
				{Text: "pending", Field: "status"},
				{Text: "problem", Field: "type"},
				{Text: "proident", Field: "description"},
				{Text: "sit", Field: "description"},
				{Text: "true", Field: "has_incidents"},
				{Text: "velit", Field: "description"},
			},
		},
		{
			name: "User",
			model: &models.User{
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
			expectedTokens: []tokeniser.Token{
				{Text: "-10:00", Field: "created_at"},
				{Text: "-10:00", Field: "last_login_at"},
				{Text: "118", Field: "organization_id"},
				{Text: "2014-06-03T02:26:28", Field: "last_login_at"},
				{Text: "2016-04-23T12:00:11", Field: "created_at"},
				{Text: "4acd4eb0-9168-4270-b09f-09600a05b0b2", Field: "external_id"},
				{Text: "59", Field: "_id"},
				{Text: "8774-883-991", Field: "phone"},
				{Text: "Be", Field: "signature"},
				{Text: "Don't", Field: "signature"},
				{Text: "Happy", Field: "signature"},
				{Text: "Key", Field: "name"},
				{Text: "Lucile", Field: "alias"},
				{Text: "Masthope", Field: "tags"},
				{Text: "Mendez", Field: "name"},
				{Text: "Mr", Field: "alias"},
				{Text: "Nigeria", Field: "timezone"},
				{Text: "Oceola", Field: "tags"},
				{Text: "Rockingham", Field: "tags"},
				{Text: "Waikele", Field: "tags"},
				{Text: "Worry", Field: "signature"},
				{Text: "agent", Field: "role"},
				{Text: "false", Field: "active"},
				{Text: "false", Field: "shared"},
				{Text: "false", Field: "suspended"},
				{Text: "false", Field: "verified"},
				{Text: "http://initech.zendesk.com/api/v2/users/59.json", Field: "url"},
				{Text: "lucilemendez@flotonic.com", Field: "email"},
				{Text: "zh-CN", Field: "locale"},
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				tokens := tokeniser.Tokenise(tc.model)
				sorter := func(a, b tokeniser.Token) int {
					if cmpText := cmp.Compare(a.Text, b.Text); cmpText != 0 {
						return cmpText
					}
					return cmp.Compare(a.Field, b.Field)
				}
				slices.SortFunc(tokens, sorter)
				slices.SortFunc(tc.expectedTokens, sorter)
				assert.DeepEqual(t, tc.expectedTokens, tokens)
			})
		})
	}
}
