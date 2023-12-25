package implementations

import (
	"fmt"
	"path/filepath"

	"github.com/satrap-illustrations/zs/internal/models"
	"github.com/satrap-illustrations/zs/internal/stores/organization"
	organizationinverted "github.com/satrap-illustrations/zs/internal/stores/organization/inverted"
	"github.com/satrap-illustrations/zs/internal/stores/ticket"
	ticketinverted "github.com/satrap-illustrations/zs/internal/stores/ticket/inverted"
	"github.com/satrap-illustrations/zs/internal/stores/user"
	userinverted "github.com/satrap-illustrations/zs/internal/stores/user/inverted"
)

type InvertedStore struct {
	organizationStore organization.Store
	ticketStore       ticket.Store
	userStore         user.Store
}

func (*InvertedStore) ListDocumentTypes() []string {
	return []string{"Organizations", "Tickets", "Users"}
}

func NewInvertedStore(path string) (*InvertedStore, error) {
	organizations := []models.Organization{}
	if err := readJSONFile(filepath.Join(path, "organizations.json"), &organizations); err != nil {
		return nil, fmt.Errorf("failed to read organizations.json: %w", err)
	}

	tickets := []models.Ticket{}
	if err := readJSONFile(filepath.Join(path, "tickets.json"), &tickets); err != nil {
		return nil, fmt.Errorf("failed to read tickets.json: %w", err)
	}

	users := []models.User{}
	if err := readJSONFile(filepath.Join(path, "users.json"), &users); err != nil {
		return nil, fmt.Errorf("failed to read users.json: %w", err)
	}

	return &InvertedStore{
		organizationStore: organizationinverted.NewOrganizationStore(organizations),
		ticketStore:       ticketinverted.NewTicketStore(tickets),
		userStore:         userinverted.NewUserStore(users),
	}, nil
}

func (h *InvertedStore) ListFields() map[string][]string {
	return map[string][]string{
		"Organizations": h.organizationStore.ListFields(),
		"Tickets":       h.ticketStore.ListFields(),
		"Users":         h.userStore.ListFields(),
	}
}

func (h *InvertedStore) Search(doctype, field, query string) ([]models.Model, error) {
	sameTypeModels := []models.Model{}
	switch doctype {
	case "Organizations":
		organizations, err := h.organizationStore.Search(field, query)
		if err != nil {
			return nil, err
		}
		sameTypeModels = append(sameTypeModels, models.OrganizationSliceToModelsSlice(organizations)...)
	case "Tickets":
		tickets, err := h.ticketStore.Search(field, query)
		if err != nil {
			return nil, err
		}
		sameTypeModels = append(sameTypeModels, models.TicketSliceToModelsSlice(tickets)...)
	case "Users":
		users, err := h.userStore.Search(field, query)
		if err != nil {
			return nil, err
		}
		sameTypeModels = append(sameTypeModels, models.UserSliceToModelsSlice(users)...)
	default:
		return nil, ErrInvalidDocType
	}
	return h.augmentWithRelatedDocuments(sameTypeModels)
}

func (h *InvertedStore) augmentWithRelatedDocuments(in []models.Model) ([]models.Model, error) {
	out := make([]models.Model, 0, len(in))
	for _, m := range in {
		out = append(out, m)
		for _, c := range m.Contains() {
			switch c.Model.(type) {
			case *models.Organization:
				organizations, err := h.organizationStore.Search(c.Field, m.StringID())
				if err != nil {
					return nil, err
				}
				out = append(out, models.OrganizationSliceToModelsSlice(organizations)...)
			case *models.Ticket:
				tickets, err := h.ticketStore.Search(c.Field, m.StringID())
				if err != nil {
					return nil, err
				}
				out = append(out, models.TicketSliceToModelsSlice(tickets)...)
			case *models.User:
				users, err := h.userStore.Search(c.Field, m.StringID())
				if err != nil {
					return nil, err
				}
				out = append(out, models.UserSliceToModelsSlice(users)...)
			}
		}
	}
	return out, nil
}
