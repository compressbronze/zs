package models

import (
	"fmt"
	"reflect"
	"strconv"
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

func (*Organization) DocumentType() string {
	return "Organization"
}

func (o *Organization) StringID() string {
	if o == nil {
		o = &Organization{}
	}
	return strconv.Itoa(o.ID)
}

func (o *Organization) Fields() *orderedmap.OrderedMap[string, int] {
	if o == nil {
		o = &Organization{}
	}
	ty := reflect.TypeOf(*o)
	fields := orderedmap.NewOrderedMap[string, int]()
	for i := 0; i < ty.NumField(); i++ {
		parts := strings.SplitN(ty.Field(i).Tag.Get("json"), ",", 2)
		switch parts[0] {
		case "-":
			continue
		case "":
			fields.Set(ty.Field(i).Name, i)
		default:
			fields.Set(parts[0], i)
		}
	}
	return fields
}

func (o *Organization) ValueAtIdx(i int) any {
	if o == nil {
		o = &Organization{}
	}
	reflectedValue := reflect.ValueOf(*o)
	return reflect.Indirect(reflectedValue).Field(i).Interface()
}

func (o *Organization) ValueAt(field string) (any, error) {
	if o == nil {
		o = &Organization{}
	}
	// sign, it is O(k) were k is the number of fields in the struct
	// we may be able to cache this in a store, but let's live with this for now
	if i, exists := o.Fields().Get(field); exists {
		return o.ValueAtIdx(i), nil
	}
	return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, field)
}

func (o *Organization) String() (string, error) {
	if o == nil {
		o = &Organization{}
	}
	return StringOf(o)
}

func (*Organization) Contains() []ContainedModel {
	return []ContainedModel{
		{
			Model: &Ticket{},
			Field: "organization_id",
		},
		{
			Model: &User{},
			Field: "organization_id",
		},
	}
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
