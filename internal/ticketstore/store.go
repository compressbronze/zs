package ticketstore

import (
	"github.com/satrap-illustrations/zs/internal/models"
)

type TicketStore interface {
	ListFields() []string
	Search(field, query string) ([]models.Ticket, error)
}
