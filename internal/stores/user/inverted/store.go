package inverted

import (
	"fmt"
	"strconv"

	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/tokeniser"
)

var ErrInvalidField = fmt.Errorf("invalid field for user store")

type UserStore struct {
	models map[int]models.User
	// we aren't counting multiplicity at the moment, we could use it to rank the results.
	index map[tokeniser.Token][]int
}

func NewUserStore(users []models.User) UserStore {
	s := UserStore{
		models: map[int]models.User{},
		index:  map[tokeniser.Token][]int{},
	}
	for i, user := range users {
		s.models[user.ID] = user
		for _, token := range tokeniser.Tokenise(&users[i]) {
			if s.index[token] == nil {
				s.index[token] = []int{user.ID}
			} else {
				s.index[token] = append(s.index[token], user.ID)
			}
		}
	}
	return s
}

func (UserStore) ListFields() []string {
	return models.FieldSlice(new(models.User))
}

func (s UserStore) Search(field, query string) ([]models.User, error) {
	u := new(models.User)
	if _, exists := u.Fields().Get(field); !exists {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == u.StringID() {
		id, err := strconv.Atoi(query)
		if err != nil {
			return nil, err
		}
		if user, ok := s.models[id]; ok {
			return []models.User{user}, nil
		}
		return []models.User{}, nil
	}

	out := []models.User{}
	token := tokeniser.Token{
		Text:  query,
		Field: field,
	}
	if indices, exist := s.index[token]; exist {
		for _, i := range indices {
			out = append(out, s.models[i])
		}
	}

	return out, nil
}
