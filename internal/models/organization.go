package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/uuid"
)

type Organization struct {
	ID            int       `json:"_id"`
	URL           string    `json:"url"`
	ExternalID    uuid.UUID `json:"external_id"`
	Name          string    `json:"name"`
	DomainNames   []string  `json:"domain_names"`
	CreatedAt     string    `json:"created_at"`
	Details       string    `json:"details"`
	SharedTickets bool      `json:"shared_tickets"`
	Tags          []string  `json:"tags"`
}

func (o *Organization) Fields() *orderedmap.OrderedMap[string, bool] {
	if o == nil {
		o = &Organization{}
	}
	ty := reflect.TypeOf(*o)
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

func (o *Organization) UnsafeValueAt(field string) any {
	if o == nil {
		o = &Organization{}
	}
	reflectedValue := reflect.ValueOf(*o)
	return reflect.Indirect(reflectedValue).FieldByName(field)
}

func (o *Organization) ValueAt(field string) (any, error) {
	if o == nil {
		o = &Organization{}
	}
	if !o.Fields().GetOrDefault(field, false) {
		return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, field)
	}
	return o.UnsafeValueAt(field), nil
}

func (o *Organization) String() (string, error) {
	if o == nil {
		o = &Organization{}
	}
	return StringOf(o)
}

// OrganizationSliceToModelsSlice converts a slice of Organization to a slice of Model.
func OrganizationSliceToModelsSlice(in []Organization) []Model {
	out := make([]Model, 0, len(in))
	for _, v := range in {
		v := v
		out = append(out, &v)
	}
	return out
}
