package models

import (
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

// Fields returns the fields in the Organization.
func (o *Organization) Fields() []string {
	ty := reflect.TypeOf(*o)
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

func (o *Organization) String() (string, error) {
	return StringOf(o)
}
