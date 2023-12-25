package models

import (
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

// Fields returns the fields in the Ticket.
func (u *User) Fields() []string {
	ty := reflect.TypeOf(*u)
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

func (u *User) String() (string, error) {
	return StringOf(u)
}
