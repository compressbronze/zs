package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/uuid"
)

type Ticket struct {
	ID             uuid.UUID `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     uuid.UUID `json:"external_id"`
	CreatedAt      string    `json:"created_at"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Description    string    `json:"description"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	SubmitterID    int       `json:"submitter_id"`
	AssigneeID     int       `json:"assignee_id"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	HasIncidents   bool      `json:"has_incidents"`
	DueAt          string    `json:"due_at"`
	Via            string    `json:"via"`
}

func (t *Ticket) Fields() *orderedmap.OrderedMap[string, bool] {
	if t == nil {
		t = &Ticket{}
	}
	ty := reflect.TypeOf(*t)
	fields := orderedmap.NewOrderedMap[string, bool]()
	for i := 0; i < ty.NumField(); i++ {
		parts := strings.SplitN(ty.Field(i).Tag.Get("json"), ",", 2)
		switch parts[0] {
		case "-":
			continue
		case "":
			fields.Set(ty.Field(i).Name, true)
		default:
			fields.Set(parts[0], true)
		}
	}
	return fields
}

func (t *Ticket) UnsafeValueAt(field string) any {
	if t == nil {
		t = &Ticket{}
	}
	reflectedValue := reflect.ValueOf(*t)
	return reflect.Indirect(reflectedValue).FieldByName(field)
}

func (t *Ticket) ValueAt(field string) (any, error) {
	if t == nil {
		t = &Ticket{}
	}
	if !t.Fields().GetOrDefault(field, false) {
		return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, field)
	}
	return t.UnsafeValueAt(field), nil
}

func (t *Ticket) String() (string, error) {
	if t == nil {
		t = &Ticket{}
	}
	return StringOf(t)
}

// TicketSliceToModelsSlice converts a slice of Ticket to a slice of Model.
func TicketSliceToModelsSlice(in []Ticket) []Model {
	out := make([]Model, 0, len(in))
	for _, v := range in {
		v := v
		out = append(out, &v)
	}
	return out
}
