package userstore

import "github.com/satrap-illustrations/zs/internal/models"

type UserStore interface {
	ListFields() []string
	Search(field, query string) ([]models.User, error)
}
