package hash

import (
	"fmt"
	"strconv"

	"github.com/satrap-illustrations/zs/internal/models"
)

var ErrInvalidField = fmt.Errorf("invalid field")

type UserStore map[int]models.User

func NewUserStore(users []models.User) UserStore {
	out := UserStore{}
	for _, user := range users {
		out[user.Id] = user
	}
	return out
}

func (UserStore) ListFields() []string {
	return models.FieldSlice(new(models.User))
}

func (s UserStore) Search(field, query string) ([]models.User, error) {
	if new(models.Ticket).Fields()[field] {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == "_id" {
		id, err := strconv.Atoi(query)
		if err != nil {
			return nil, err
		}
		if user, ok := s[id]; ok {
			return []models.User{user}, nil
		}
		return []models.User{}, nil
	}

	out := []models.User{}
	for _, user := range s {
		if user.UnsafeValueAt(field) == query {
			out = append(out, user)
		}
	}

	return out, nil
}
