package ticket

import (
	"github.com/satrap-illustrations/zs/internal/models"
)

type Store interface {
	ListFields() []string
	Search(field, query string) ([]models.Ticket, error)
}
