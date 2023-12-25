package stores_test

import (
	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
)

//nolint:dupword,lll
var organization118Results = []models.Model{
	&models.Organization{
		ID:          118,
		URL:         "http://initech.zendesk.com/api/v2/organizations/118.json",
		ExternalID:  uuid.Must(uuid.Parse("6970300e-f211-4c01-a538-70b4464a1d84")),
		Name:        "Limozen",
		DomainNames: []string{"otherway.com", "rodeomad.com", "suremax.com", "fishland.com"},
		CreatedAt:   "2016-02-11T04:24:09 -11:00",
		Details:     "MegaCorp",
		Tags:        []string{"Leon", "Ferguson", "Olsen", "Walsh"},
	},
	&models.Ticket{
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
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("17951590-6a78-49e8-8e45-1d4326ba49cc")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/17951590-6a78-49e8-8e45-1d4326ba49cc.json",
		ExternalID:     uuid.Must(uuid.Parse("f77cae39-867c-4890-9696-b4d5c7748fa3")),
		CreatedAt:      "2016-06-28T03:29:34 -10:00",
		Type:           "incident",
		Subject:        "A Nuisance in Kenya",
		Description:    "Magna est nostrud commodo sint aliqua labore deserunt. Do est dolore enim duis non culpa fugiat laboris exercitation et.",
		Priority:       "normal",
		Status:         "open",
		SubmitterID:    53,
		OrganizationID: 118,
		Tags: []string{
			"District Of Columbia", "Wisconsin", "Illinois",
			"Fédératéd Statés Of Micronésia",
		},
		HasIncidents: true,
		DueAt:        "2016-08-16T09:10:29 -10:00",
		Via:          "chat",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("3d0d0ce2-6d1b-4f8d-a743-3863aeb29aab")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/3d0d0ce2-6d1b-4f8d-a743-3863aeb29aab.json",
		ExternalID:     uuid.Must(uuid.Parse("a6793810-40a3-486c-a002-bd43b384c759")),
		CreatedAt:      "2016-06-07T12:05:31 -10:00",
		Type:           "task",
		Subject:        "A Problem in Pitcairn",
		Description:    "Reprehenderit eiusmod dolore deserunt deserunt nostrud labore amet exercitation laborum. Consequat enim nostrud in id voluptate esse nostrud deserunt quis culpa cillum nulla ullamco nulla.",
		Priority:       "urgent",
		Status:         "pending",
		SubmitterID:    41,
		AssigneeID:     64,
		OrganizationID: 118,
		Tags:           []string{"Guam", "Colorado", "Washington", "Wyoming"},
		DueAt:          "2016-08-07T05:39:40 -10:00",
		Via:            "chat",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("4c5a405d-0805-4d8b-ac48-2a3d7f3816e4")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/4c5a405d-0805-4d8b-ac48-2a3d7f3816e4.json",
		ExternalID:     uuid.Must(uuid.Parse("3e690d6e-b322-4d56-b72a-a18cda46f717")),
		CreatedAt:      "2016-03-08T04:22:35 -11:00",
		Type:           "incident",
		Subject:        "A Drama in Haiti",
		Description:    "Enim commodo officia laborum veniam anim nisi occaecat. Lorem voluptate cupidatat do eu irure reprehenderit culpa.",
		Priority:       "high",
		Status:         "hold",
		SubmitterID:    74,
		AssigneeID:     16,
		OrganizationID: 118,
		Tags:           []string{"Massachusetts", "New York", "Minnesota", "New Jersey"},
		DueAt:          "2016-08-06T09:42:11 -10:00",
		Via:            "voice",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("53867869-0db0-4b8d-9d6c-9d1c0af4e693")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/53867869-0db0-4b8d-9d6c-9d1c0af4e693.json",
		ExternalID:     uuid.Must(uuid.Parse("d3b44197-5e5f-4dee-82de-bda68efb6210")),
		CreatedAt:      "2016-05-14T09:19:56 -10:00",
		Type:           "task",
		Subject:        "A Drama in Gabon",
		Description:    "Eu anim laborum enim voluptate ex minim quis magna culpa occaecat qui amet anim. Consectetur adipisicing sunt est fugiat cillum eiusmod elit nostrud cupidatat culpa esse eiusmod.",
		Priority:       "urgent",
		Status:         "solved",
		SubmitterID:    51,
		AssigneeID:     5,
		OrganizationID: 118,
		Tags:           []string{"Utah", "Hawaii", "Alaska", "Maryland"},
		DueAt:          "2016-08-14T06:11:52 -10:00",
		Via:            "web",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("7382ad0e-dea7-4c8d-b38f-cbbf016f2598")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/7382ad0e-dea7-4c8d-b38f-cbbf016f2598.json",
		ExternalID:     uuid.Must(uuid.Parse("6d3b0e05-6013-4513-9913-0bb6a0f66ef7")),
		CreatedAt:      "2016-03-31T03:16:52 -11:00",
		Type:           "task",
		Subject:        "A Problem in American Samoa",
		Description:    "Excepteur dolor in commodo minim irure laboris. In incididunt mollit veniam pariatur ullamco laborum ullamco aliqua do fugiat Lorem.",
		Priority:       "high",
		Status:         "closed",
		SubmitterID:    35,
		AssigneeID:     64,
		OrganizationID: 118,
		Tags:           []string{"Missouri", "Alabama", "Virginia", "Virgin Islands"},
		HasIncidents:   true,
		DueAt:          "2016-08-06T08:36:17 -10:00",
		Via:            "chat",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("87db32c5-76a3-4069-954c-7d59c6c21de0")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/87db32c5-76a3-4069-954c-7d59c6c21de0.json",
		ExternalID:     uuid.Must(uuid.Parse("1c61056c-a5ad-478a-9fd6-38889c3cd728")),
		CreatedAt:      "2016-07-06T11:16:50 -10:00",
		Type:           "problem",
		Subject:        "A Problem in Morocco",
		Description:    "Sit culpa non magna anim. Ea velit qui nostrud eiusmod laboris dolor adipisicing quis deserunt elit amet.",
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
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("8d7b4d51-ef95-4923-9ab8-42332ab2188d")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/8d7b4d51-ef95-4923-9ab8-42332ab2188d.json",
		ExternalID:     uuid.Must(uuid.Parse("c0a785cf-b0e0-4627-acb6-97adac4b7be6")),
		CreatedAt:      "2016-05-30T02:40:22 -10:00",
		Type:           "question",
		Subject:        "A Catastrophe in Malta",
		Description:    "Est consequat elit do do id laborum ad enim sit nostrud id eiusmod. Labore tempor velit cupidatat aliquip excepteur anim aliquip aliquip.",
		Priority:       "high",
		Status:         "pending",
		SubmitterID:    3,
		AssigneeID:     8,
		OrganizationID: 118,
		Tags:           []string{"Virginia", "Virgin Islands", "Maine", "West Virginia"},
		HasIncidents:   true,
		DueAt:          "2016-08-12T02:41:31 -10:00",
		Via:            "voice",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("92e5d8f0-853a-4f56-b7fb-b0582e6b1c79")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/92e5d8f0-853a-4f56-b7fb-b0582e6b1c79.json",
		ExternalID:     uuid.Must(uuid.Parse("39e2b2fa-9d90-4390-beb5-2bade85ce5ba")),
		CreatedAt:      "2016-01-06T09:27:57 -11:00",
		Type:           "incident",
		Subject:        "A Drama in Nepal",
		Description:    "Et occaecat elit enim tempor ipsum. Sint sit proident sit ipsum cillum voluptate ipsum nostrud officia sint exercitation reprehenderit id eu.",
		Priority:       "high",
		Status:         "pending",
		SubmitterID:    8,
		AssigneeID:     72,
		OrganizationID: 118,
		Tags:           []string{"Kentucky", "North Carolina", "South Carolina", "Indiana"},
		HasIncidents:   true,
		DueAt:          "2016-08-05T09:42:07 -10:00",
		Via:            "voice",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("945ce2d3-3edc-4936-8d51-e59e74cf917a")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/945ce2d3-3edc-4936-8d51-e59e74cf917a.json",
		ExternalID:     uuid.Must(uuid.Parse("5c741d66-cdd4-4d20-bb95-a3948217bf2c")),
		CreatedAt:      "2016-04-23T05:47:03 -10:00",
		Type:           "task",
		Subject:        "A Drama in Guinea",
		Description:    "Esse Lorem qui cillum amet enim sint aute duis veniam non. Esse irure sit qui non amet reprehenderit ullamco tempor duis exercitation excepteur.",
		Priority:       "urgent",
		Status:         "hold",
		SubmitterID:    70,
		AssigneeID:     32,
		OrganizationID: 118,
		Tags:           []string{"American Samoa", "Northern Mariana Islands", "Puerto Rico", "Idaho"},
		HasIncidents:   true,
		DueAt:          "2016-07-31T05:29:05 -10:00",
		Via:            "voice",
	},
	&models.Ticket{
		ID:             uuid.Must(uuid.Parse("ad49f154-2ceb-4052-9129-ddc6d4b7e479")),
		URL:            "http://initech.zendesk.com/api/v2/tickets/ad49f154-2ceb-4052-9129-ddc6d4b7e479.json",
		ExternalID:     uuid.Must(uuid.Parse("8a57c17a-c7bc-4b1c-bfad-eec83f4a791d")),
		CreatedAt:      "2016-05-17T08:32:44 -10:00",
		Type:           "question",
		Subject:        "A Problem in Kyrgyzstan",
		Description:    "Pariatur eu ipsum esse qui. Quis minim ea deserunt enim do cupidatat velit aliqua qui duis pariatur velit consectetur.",
		Priority:       "high",
		Status:         "closed",
		SubmitterID:    3,
		AssigneeID:     31,
		OrganizationID: 118,
		Tags:           []string{"Georgia", "Tennessee", "Mississippi", "Marshall Islands"},
		HasIncidents:   true,
		DueAt:          "2016-07-31T02:59:08 -10:00",
		Via:            "voice",
	},
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
}
