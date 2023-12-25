package models

import (
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type Ticket struct {
	Id             uuid.UUID      `json:"_id"`
	URL            string         `json:"url"`
	ExternalId     uuid.UUID      `json:"external_id"`
	CreatedAt      string         `json:"created_at"`
	Type           string         `json:"type"`
	Subject        string         `json:"subject"`
	Description    string         `json:"description"`
	Priority       string         `json:"priority"`
	Status         string         `json:"pending"`
	SubmitterId    int            `json:"submitter_id"`
	AssigneeId     int            `json:"assignee_id"`
	OrganizationId int            `json:"organization_id"`
	Tags           []string       `json:"tags"`
	HasIncidents   bool           `json:"has_incidents"`
	DueAt          string         `json:"due_at"`
	Via            string         `json:"web"`
	Data           map[string]any `json:"-"`
}

// Fields returns the fields in the Ticket.
func (t *Ticket) Fields() []string {
	ty := reflect.TypeOf(*t)
	fields := make([]string, 0, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		parts := strings.SplitN(ty.Field(i).Tag.Get("json"), ",", 2)
		if parts[0] == "-" {
			continue
		}
		fields = append(fields, parts[0])
	}
	return fields
}

func (t *Ticket) String() (string, error) {
	return StringOf(t)
}
