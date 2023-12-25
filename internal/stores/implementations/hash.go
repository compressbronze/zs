package implementations

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

func readJSONFile(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}

func (*HashStore) ListDocumentTypes() []string {
	return []string{"Organizations", "Tickets", "Users"}
}

func NewHashStore(path string) (*HashStore, error) {
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

	return &HashStore{
		organizationStore: organizationhash.NewOrganizationStore(organizations),
		ticketStore:       tickethash.NewTicketStore(tickets),
		userStore:         userhash.NewUserStore(users),
	}, nil
}

func (h *HashStore) ListFields() map[string][]string {
	return map[string][]string{
		"organizations": h.organizationStore.ListFields(),
		"tickets":       h.ticketStore.ListFields(),
		"users":         h.userStore.ListFields(),
	}
}

func (h *HashStore) Search(doctype, field, query string) ([]models.Model, error) {
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

func (h *HashStore) augmentWithRelatedDocuments(in []models.Model) ([]models.Model, error) {
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
