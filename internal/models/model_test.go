package models_test

import (
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/satrap-illustrations/zs/internal/models"
	"gotest.tools/v3/assert"
)

func toMap(o *orderedmap.OrderedMap[string, bool]) map[string]bool {
	m := make(map[string]bool)
	for _, k := range o.Keys() {
		m[k] = true
	}
	return m
}

func TestFields(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name           string
		model          models.Model
		expectedFields map[string]bool
	}{
		{
			name:  "organization_nil",
			model: (*models.Organization)(nil),
			expectedFields: map[string]bool{
				"_id":            true,
				"url":            true,
				"external_id":    true,
				"name":           true,
				"domain_names":   true,
				"created_at":     true,
				"details":        true,
				"shared_tickets": true,
				"tags":           true,
			},
		},
		{
			name:  "organization_zero",
			model: new(models.Organization),
			expectedFields: map[string]bool{
				"_id":            true,
				"url":            true,
				"external_id":    true,
				"name":           true,
				"domain_names":   true,
				"created_at":     true,
				"details":        true,
				"shared_tickets": true,
				"tags":           true,
			},
		},
		{
			name:  "user_nil",
			model: (*models.User)(nil),
			expectedFields: map[string]bool{
				"_id":             true,
				"url":             true,
				"external_id":     true,
				"name":            true,
				"alias":           true,
				"created_at":      true,
				"active":          true,
				"verified":        true,
				"shared":          true,
				"locale":          true,
				"timezone":        true,
				"last_login_at":   true,
				"email":           true,
				"phone":           true,
				"signature":       true,
				"organization_id": true,
				"tags":            true,
				"suspended":       true,
				"role":            true,
			},
		},
		{
			name:  "user_zero",
			model: new(models.User),
			expectedFields: map[string]bool{
				"_id":             true,
				"url":             true,
				"external_id":     true,
				"name":            true,
				"alias":           true,
				"created_at":      true,
				"active":          true,
				"verified":        true,
				"shared":          true,
				"locale":          true,
				"timezone":        true,
				"last_login_at":   true,
				"email":           true,
				"phone":           true,
				"signature":       true,
				"organization_id": true,
				"tags":            true,
				"suspended":       true,
				"role":            true,
			},
		},
		{
			name:  "ticket_nil",
			model: (*models.Ticket)(nil),
			expectedFields: map[string]bool{
				"_id":             true,
				"url":             true,
				"external_id":     true,
				"created_at":      true,
				"type":            true,
				"subject":         true,
				"description":     true,
				"priority":        true,
				"status":          true,
				"submitter_id":    true,
				"assignee_id":     true,
				"organization_id": true,
				"tags":            true,
				"has_incidents":   true,
				"due_at":          true,
				"via":             true,
			},
		},
		{
			name:  "ticket_zero",
			model: new(models.Ticket),
			expectedFields: map[string]bool{
				"_id":             true,
				"url":             true,
				"external_id":     true,
				"created_at":      true,
				"type":            true,
				"subject":         true,
				"description":     true,
				"priority":        true,
				"status":          true,
				"submitter_id":    true,
				"assignee_id":     true,
				"organization_id": true,
				"tags":            true,
				"has_incidents":   true,
				"due_at":          true,
				"via":             true,
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fields := tc.model.Fields()
			assert.DeepEqual(t, tc.expectedFields, toMap(fields))
		})
	}
}
