package models_test

import (
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
	"gotest.tools/v3/assert"
)

var dummyUUID = uuid.New()

func TestFieldsSlice(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name           string
		model          models.Model
		expectedFields []string
	}{
		{
			name:  "organization_nil",
			model: (*models.Organization)(nil),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"name",
				"domain_names",
				"created_at",
				"details",
				"shared_tickets",
				"tags",
			},
		},
		{
			name:  "organization_zero",
			model: new(models.Organization),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"name",
				"domain_names",
				"created_at",
				"details",
				"shared_tickets",
				"tags",
			},
		},
		{
			name:  "user_nil",
			model: (*models.User)(nil),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"name",
				"alias",
				"created_at",
				"active",
				"verified",
				"shared",
				"locale",
				"timezone",
				"last_login_at",
				"email",
				"phone",
				"signature",
				"organization_id",
				"tags",
				"suspended",
				"role",
			},
		},
		{
			name:  "user_zero",
			model: new(models.User),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"name",
				"alias",
				"created_at",
				"active",
				"verified",
				"shared",
				"locale",
				"timezone",
				"last_login_at",
				"email",
				"phone",
				"signature",
				"organization_id",
				"tags",
				"suspended",
				"role",
			},
		},
		{
			name:  "ticket_nil",
			model: (*models.Ticket)(nil),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"created_at",
				"type",
				"subject",
				"description",
				"priority",
				"status",
				"submitter_id",
				"assignee_id",
				"organization_id",
				"tags",
				"has_incidents",
				"due_at",
				"via",
			},
		},
		{
			name:  "ticket_zero",
			model: new(models.Ticket),
			expectedFields: []string{
				"_id",
				"url",
				"external_id",
				"created_at",
				"type",
				"subject",
				"description",
				"priority",
				"status",
				"submitter_id",
				"assignee_id",
				"organization_id",
				"tags",
				"has_incidents",
				"due_at",
				"via",
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.DeepEqual(t, tc.expectedFields, models.FieldSlice(tc.model))
		})
	}
}

func TestValueAt(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name          string
		model         models.Model
		field         string
		expectedValue any
		expectedError error
	}{
		{
			name:          "organization_nil_fake",
			model:         (*models.Organization)(nil),
			field:         "fake",
			expectedValue: 0,
			expectedError: models.ErrFieldNotFound,
		},
		{
			name:          "organization_nil_id",
			model:         (*models.Organization)(nil),
			field:         "_id",
			expectedValue: 0,
			expectedError: nil,
		},
		{
			name:          "ticket_nil_fake",
			model:         (*models.Ticket)(nil),
			field:         "fake",
			expectedValue: 0,
			expectedError: models.ErrFieldNotFound,
		},
		{
			name:          "ticket_nil_id",
			model:         (*models.Ticket)(nil),
			field:         "_id",
			expectedValue: uuid.UUID{},
			expectedError: nil,
		},
		{
			name:          "user_nil_fake",
			model:         (*models.User)(nil),
			field:         "fake",
			expectedValue: 0,
			expectedError: models.ErrFieldNotFound,
		},
		{
			name:          "user_nil_id",
			model:         (*models.User)(nil),
			field:         "_id",
			expectedValue: 0,
			expectedError: nil,
		},
		{
			name:          "ticket_id",
			model:         &models.Ticket{ID: dummyUUID},
			field:         "_id",
			expectedValue: dummyUUID,
			expectedError: nil,
		},
		{
			name:          "user_id",
			model:         &models.User{ID: 12},
			field:         "_id",
			expectedValue: 12,
			expectedError: nil,
		},
		{
			name:          "user_tags",
			model:         &models.User{ID: 12, Tags: []string{"a", "b", "c"}},
			field:         "tags",
			expectedValue: []string{"a", "b", "c"},
			expectedError: nil,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			value, err := tc.model.ValueAt(tc.field)
			if tc.expectedError != nil {
				assert.ErrorIs(t, err, tc.expectedError)
				return
			}
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expectedValue, value)
		})
	}
}

func toMap(o *orderedmap.OrderedMap[string, int]) map[string]bool {
	m := make(map[string]bool)
	for _, k := range o.Keys() {
		m[k] = true
	}
	return m
}
