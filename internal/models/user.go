package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/uuid"
)

type User struct {
	ID             int       `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     uuid.UUID `json:"external_id"`
	Name           string    `json:"name"`
	Alias          string    `json:"alias"`
	CreatedAt      string    `json:"created_at"`
	Active         bool      `json:"active"`
	Verified       bool      `json:"verified"`
	Shared         bool      `json:"shared"`
	Locale         string    `json:"locale"`
	Timezone       string    `json:"timezone"`
	LastLoginAt    string    `json:"last_login_at"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Signature      string    `json:"signature"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	Suspended      bool      `json:"suspended"`
	Role           string    `json:"role"`
}

func (u *User) Fields() *orderedmap.OrderedMap[string, int] {
	if u == nil {
		u = &User{}
	}
	ty := reflect.TypeOf(*u)
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

func (u *User) ValueAtIdx(i int) any {
	if u == nil {
		u = &User{}
	}
	reflectedValue := reflect.ValueOf(*u)
	return reflect.Indirect(reflectedValue).Field(i).Interface()
}

func (u *User) ValueAt(field string) (any, error) {
	if u == nil {
		u = &User{}
	}
	// sign, it is O(k) were k is the number of fields in the struct
	// we may be able to cache this in a store, but let's live with this for now
	if i, exists := u.Fields().Get(field); exists {
		return u.ValueAtIdx(i), nil
	}
	return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, field)
}

func (u *User) String() (string, error) {
	if u == nil {
		u = &User{}
	}
	return StringOf(u)
}

// UserSliceToModelsSlice converts a slice of User to a slice of Model.
func UserSliceToModelsSlice(in []User) []Model {
	out := make([]Model, 0, len(in))
	for _, v := range in {
		v := v
		out = append(out, &v)
	}
	return out
}
