package inverted

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/tokeniser"
)

var ErrInvalidField = fmt.Errorf("invalid field for ticket store")

type TicketStore struct {
	models map[uuid.UUID]models.Ticket
	// we aren't counting multiplicity at the moment, we could use it to rank the results.
	index map[tokeniser.Token][]uuid.UUID
}

func NewTicketStore(tickets []models.Ticket) TicketStore {
	s := TicketStore{
		models: map[uuid.UUID]models.Ticket{},
		index:  map[tokeniser.Token][]uuid.UUID{},
	}
	for i, ticket := range tickets {
		s.models[ticket.ID] = ticket
		for _, token := range tokeniser.Tokenise(&tickets[i]) {
			if s.index[token] == nil {
				s.index[token] = []uuid.UUID{ticket.ID}
			} else {
				s.index[token] = append(s.index[token], ticket.ID)
			}
		}
	}
	return s
}

func (TicketStore) ListFields() []string {
	return models.FieldSlice(new(models.Ticket))
}

func (s TicketStore) Search(field, query string) ([]models.Ticket, error) {
	t := new(models.Ticket)
	if _, exists := t.Fields().Get(field); !exists {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == t.StringID() {
		id, err := uuid.Parse(query)
		if err != nil {
			return nil, err
		}
		if ticket, ok := s.models[id]; ok {
			return []models.Ticket{ticket}, nil
		}
		return []models.Ticket{}, nil
	}

	out := []models.Ticket{}
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
