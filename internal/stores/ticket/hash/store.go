package hash

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
)

var ErrInvalidField = fmt.Errorf("invalid field")

type TicketStore map[uuid.UUID]models.Ticket

func NewTicketStore(tickets []models.Ticket) TicketStore {
	out := TicketStore{}
	for _, ticket := range tickets {
		out[ticket.Id] = ticket
	}
	return out
}

func (TicketStore) ListFields() []string {
	return models.FieldSlice(new(models.Ticket))
}

func (s TicketStore) Search(field, query string) ([]models.Ticket, error) {
	if new(models.Ticket).Fields()[field] {
		return nil, fmt.Errorf("%w: %s", ErrInvalidField, field)
	}

	if field == "_id" {
		id, err := uuid.Parse(query)
		if err != nil {
			return nil, err
		}
		if ticket, ok := s[id]; ok {
			return []models.Ticket{ticket}, nil
		}
		return []models.Ticket{}, nil
	}

	out := []models.Ticket{}
	for _, ticket := range s {
		if ticket.UnsafeValueAt(field) == query {
			out = append(out, ticket)
		}
	}

	return out, nil
}
