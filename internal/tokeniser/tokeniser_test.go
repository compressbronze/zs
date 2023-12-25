package tokeniser_test

import (
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
		expectedTokens []string
	}{
		{
			name:           "Organization_zero",
			model:          &models.Organization{},
			expectedTokens: []string{"0", "false"},
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
			expectedTokens: []string{
				"-11:00",
				"118",
				"2016-02-11T04:24:09",
				"6970300e-f211-4c01-a538-70b4464a1d84",
				"Ferguson",
				"Leon",
				"Limozen",
				"MegaCorp",
				"Olsen",
				"Walsh",
				"false",
				"fishland.com",
				"http://initech.zendesk.com/api/v2/organizations/118.json",
				"otherway.com",
				"rodeomad.com",
				"suremax.com",
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
			expectedTokens: []string{
				"-10:00",
				"-10:00",
				"0ebe753c-9c78-458a-817f-3993780bedbf",
				"118",
				"2016-05-19T12:19:56",
				"2016-08-18T03:33:30",
				"23",
				"537ad752-9056-42c9-86db-f0bdf06d3c10",
				"56",
				"A",
				"Alabama",
				"Consequat",
				"Islands",
				"Lorem",
				"Missouri",
				"Nuisance",
				"Seychelles",
				"Virgin",
				"Virginia",
				"ad",
				"aliqua",
				"aliquip",
				"chat",
				"ea",
				"ea",
				"enim",
				"est",
				"est",
				"high",
				"http://initech.zendesk.com/api/v2/tickets/0ebe753c-9c78-458a-817f-3993780bedbf.json",
				"id",
				"in",
				"magna",
				"magna",
				"mollit",
				"pending",
				"problem",
				"proident",
				"sit",
				"true",
				"velit",
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
			expectedTokens: []string{
				"-10:00",
				"-10:00",
				"118",
				"2014-06-03T02:26:28",
				"2016-04-23T12:00:11",
				"4acd4eb0-9168-4270-b09f-09600a05b0b2",
				"59",
				"8774-883-991",
				"Be",
				"Don't",
				"Happy",
				"Key",
				"Lucile",
				"Masthope",
				"Mendez",
				"Mr",
				"Nigeria",
				"Oceola",
				"Rockingham",
				"Waikele",
				"Worry",
				"agent",
				"false",
				"false",
				"false",
				"false",
				"http://initech.zendesk.com/api/v2/users/59.json",
				"lucilemendez@flotonic.com",
				"zh-CN",
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				tokens := tokeniser.Tokenise(tc.model)
				slices.Sort(tokens)
				slices.Sort(tc.expectedTokens)
				assert.DeepEqual(t, tc.expectedTokens, tokens)
			})
		})
	}
}
