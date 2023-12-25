package implementations

import (
	"errors"

	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/stores/organization"
	organizationhash "github.com/satrap-illustrations/zs/internal/stores/organization/hash"
	"github.com/satrap-illustrations/zs/internal/stores/ticket"
	tickethash "github.com/satrap-illustrations/zs/internal/stores/ticket/hash"
	"github.com/satrap-illustrations/zs/internal/stores/user"
	userhash "github.com/satrap-illustrations/zs/internal/stores/user/hash"
)

var ErrInvalidDocType = errors.New("invalid document type")

type HashStore struct {
	organizationStore organization.Store
	ticketStore       ticket.Store
	userStore         user.Store
}

func NewHashStore(path string) *HashStore {
	organizations := []models.Organization{}
	tickets := []models.Ticket{}
	users := []models.User{}

	return &HashStore{
		organizationStore: organizationhash.NewOrganizationStore(organizations),
		ticketStore:       tickethash.NewTicketStore(tickets),
		userStore:         userhash.NewUserStore(users),
	}
}

func (h *HashStore) ListFields() map[string][]string {
	return map[string][]string{
		"organizations": h.organizationStore.ListFields(),
		"tickets":       h.ticketStore.ListFields(),
		"users":         h.userStore.ListFields(),
	}
}

func (h *HashStore) Search(doctype, field, query string) ([]models.Model, error) {
	switch doctype {
	case "organizations":
		organizations, err := h.organizationStore.Search(field, query)
		return models.OrganizationSliceToModelsSlice(organizations), err
	case "tickets":
		tickets, err := h.ticketStore.Search(field, query)
		return models.TicketSliceToModelsSlice(tickets), err
	case "users":
		users, err := h.userStore.Search(field, query)
		return models.UserSliceToModelsSlice(users), err
	default:
		return nil, ErrInvalidDocType
	}
}
