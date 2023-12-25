package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type Organization struct {
	Id            int       `json:"_id"`
	URL           string    `json:"url"`
	ExternalId    uuid.UUID `json:"external_id"`
	Name          string    `json:"name"`
	DomainNames   []string  `json:"domain_names"`
	CreatedAt     string    `json:"created_at"`
	Details       string    `json:"details"`
	SharedTickets bool      `json:"shared_tickets"`
	Tags          []string  `json:"tags"`
}

func (o *Organization) Fields() map[string]bool {
	if o == nil {
		o = &Organization{}
	}
	ty := reflect.TypeOf(*o)
	fields := make(map[string]bool, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		parts := strings.SplitN(ty.Field(i).Tag.Get("json"), ",", 2)
		if parts[0] == "-" {
			continue
		}
		fields[parts[0]] = true
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
	if _, exists := o.Fields()[field]; !exists {
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
