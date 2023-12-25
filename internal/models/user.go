package models

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Id             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"admin"`
}

func (u *User) Fields() map[string]bool {
	if u == nil {
		u = &User{}
	}
	ty := reflect.TypeOf(*u)
	fields := make(map[string]bool, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		parts := strings.SplitN(ty.Field(i).Tag.Get("json"), ",", 2)
		switch parts[0] {
		case "-":
			continue
		case "":
			fields[ty.Field(i).Name] = true
		default:
			fields[parts[0]] = true
		}
	}
	return fields
}

func (u *User) UnsafeValueAt(field string) any {
	if u == nil {
		u = &User{}
	}
	reflectedValue := reflect.ValueOf(*u)
	return reflect.Indirect(reflectedValue).FieldByName(field)
}

func (u *User) ValueAt(field string) (any, error) {
	if u == nil {
		u = &User{}
	}
	if _, exists := u.Fields()[field]; !exists {
		return nil, fmt.Errorf("%w: %s", ErrFieldNotFound, field)
	}
	return u.UnsafeValueAt(field), nil
}

func (u *User) String() (string, error) {
	if u == nil {
		u = &User{}
	}
	return StringOf(u)
}
