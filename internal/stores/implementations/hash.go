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
	userStore         user.Store
	ticketStore       ticket.Store
}

func NewHashStore(path string) *HashStore {
	orgs := []models.Organization{}
	tickets := []models.Ticket{}
	users := []models.User{}

	return &HashStore{
		organizationStore: organizationhash.NewOrganizationStore(orgs),
		ticketStore:       tickethash.NewTicketStore(tickets),
		userStore:         userhash.NewUserStore(users),
	}
}

func (h *HashStore) ListFields() map[string][]string {
	return map[string][]string{
		"organizations": h.organizationStore.ListFields(),
		"users":         h.userStore.ListFields(),
		"tickets":       h.ticketStore.ListFields(),
	}
}

func (h *HashStore) Search(doctype, field, query string) ([]models.Model, error) {
	switch doctype {
	case "organizations":
		orgs, err := h.organizationStore.Search(field, query)
		return []models.Model(orgs), err
	case "users":
		return h.userStore.Search(field, query)
	case "tickets":
		return h.ticketStore.Search(field, query)
	default:
		return nil, ErrInvalidDocType
	}

	return nil, nil
}
