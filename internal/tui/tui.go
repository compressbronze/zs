package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/satrap-illustrations/zs/internal/stores"
	"github.com/satrap-illustrations/zs/internal/tui/doctypelist"
)

type state int

const (
	header state = iota
	selectOptions
	search
	chosenDocType
	chosenDocTypeField
	results
	listFields
)

type model struct {
	state         state
	styles        *styles
	width, height int
	docType       doctypelist.Model
	field, query  textinput.Model
	store         stores.Store
}

type styles struct {
	docTypeField lipgloss.Style
	fieldField   lipgloss.Style
	queryField   lipgloss.Style
}

func DefaultStyles() *styles {
	return &styles{
		docTypeField: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#154733")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
		fieldField: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#ed095d")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
		queryField: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#a134eb")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
	}
}

func InitialModel(store stores.Store) model {
	styles := DefaultStyles()
	docTypes := store.ListDocumentTypes()
	docType := doctypelist.New(docTypes)

	field := textinput.New()
	field.Placeholder = "Input a field to search in..."

	query := textinput.New()
	query.Placeholder = "Input a word to search for..."
	return model{
		styles:  styles,
		docType: docType,
		field:   field,
		query:   query,
		store:   store,
	}
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	newModel := m

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		newModel.width = msg.Width
		newModel.height = msg.Height

	case tea.KeyMsg:
		switch s := msg.String(); s {
		// ctrl+c should exit the program from any state.
		case "ctrl+c":
			return newModel, tea.Quit

		default:
			switch newModel.state {
			case header:
				switch s {
				case "enter":
					newModel.state = selectOptions
					return newModel, nil
				default:
					return newModel, nil
				}
			case selectOptions:
				switch s {
				case "1":
					newModel.state = search
				case "2":
					newModel.state = listFields
				}
				return newModel, nil
			case search:
				switch s {
				case "ctrl+d":
					newModel.state = selectOptions
					return newModel, nil
				case "enter":
					newModel.state = chosenDocType
					newModel.field.Focus()
					return newModel, nil
				default:
					newModel.docType, cmd = newModel.docType.Update(msg)
					return newModel, cmd
				}
			case chosenDocType:
				switch s {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.docType, _ = newModel.docType.Update("")
					newModel.field, cmd = newModel.field.Update("")
					return newModel, cmd
				case "enter":
					newModel.state = chosenDocTypeField
					newModel.query.Focus()
					return newModel, cmd
				default:
					newModel.field, cmd = newModel.field.Update(msg)
					return newModel, cmd
				}
			case chosenDocTypeField:
				switch s {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.docType, _ = newModel.docType.Update("")
					newModel.field, _ = newModel.field.Update("")
					newModel.query, _ = newModel.query.Update("")
					return newModel, nil
				default:
					newModel.query, cmd = newModel.query.Update(msg)
					return newModel, cmd
				}
			case listFields:
				switch s {
				case "enter":
					newModel.state = selectOptions
					return newModel, nil
				default:
					return newModel, nil
				}
			}
		}
	}

	newModel.query, cmd = newModel.query.Update(msg)
	return newModel, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	const headerText = `
      ___           ___           ___           ___           ___           ___           ___
     /\  \         /\  \         /\__\         /\  \         /\  \         /\  \         /\__\
     \:\  \       /::\  \       /::|  |       /::\  \       /::\  \       /::\  \       /:/  /
      \:\  \     /:/\:\  \     /:|:|  |      /:/\:\  \     /:/\:\  \     /:/\ \  \     /:/__/
       \:\  \   /::\~\:\  \   /:/|:|  |__   /:/  \:\__\   /::\~\:\  \   _\:\~\ \  \   /::\__\____
 _______\:\__\ /:/\:\ \:\__\ /:/ |:| /\__\ /:/__/ \:|__| /:/\:\ \:\__\ /\ \:\ \ \__\ /:/\:::::\__\
 \::::::::/__/ \:\~\:\ \/__/ \/__|:|/:/  / \:\  \ /:/  / \:\~\:\ \/__/ \:\ \:\ \/__/ \/_|:|~~|/__/
  \:\~~\~~      \:\ \:\__\       |:/:/  /   \:\  /:/  /   \:\ \:\__\    \:\ \:\__\      |:|  |
   \:\  \        \:\ \/__/       |::/  /     \:\/:/  /     \:\ \/__/     \:\/:/  /      |:|  |
    \:\__\        \:\__\         /:/  /       \::/__/       \:\__\        \::/  /       |:|  |
     \/__/         \/__/         \/__/         ~~            \/__/         \/__/         \|__|
               ___           ___           ___           ___           ___           ___
              /\  \         /\  \         /\  \         /\  \         /\  \         /\__\
             /::\  \       /::\  \       /::\  \       /::\  \       /::\  \       /:/  /
            /:/\ \  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/__/
           _\:\~\ \  \   /::\~\:\  \   /::\~\:\  \   /::\~\:\  \   /:/  \:\  \   /::\  \ ___
          /\ \:\ \ \__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/__/ \:\__\ /:/\:\  /\__\
          \:\ \:\ \/__/ \:\~\:\ \/__/ \/__\:\/:/  / \/_|::\/:/  / \:\  \  \/__/ \/__\:\/:/  /
           \:\ \:\__\    \:\ \:\__\        \::/  /     |:|::/  /   \:\  \            \::/  /
            \:\/:/  /     \:\ \/__/        /:/  /      |:|\/__/     \:\  \           /:/  /
             \::/  /       \:\__\         /:/  /       |:|  |        \:\__\         /:/  /
              \/__/         \/__/         \/__/         \|__|         \/__/         \/__/

Welcome to Zendesk Search
Type 'ctrl+c' to exit at any time, Press 'enter' to continue.`

	const searchText = `
Select search options:
1) Search Zendesk
2) View a list of searchable fields
`

	switch m.state {
	case header:
		return headerText
	case selectOptions:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			searchText,
		)
	case search:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			m.styles.docTypeField.Render(m.docType.View()),
			"",
		)
	case chosenDocType:
		m.field.Placeholder = fmt.Sprintf(
			"Type a %s field to search...",
			m.docType.SelectedItem(),
		)
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			fmt.Sprintf("\nSearching %s documents", m.docType.SelectedItem()),
			m.styles.fieldField.Render(m.field.View()),
			"",
		)
	case chosenDocTypeField:
		m.query.Placeholder = fmt.Sprintf(
			"Type a value to search for %s in %s...",
			m.field.Value(),
			m.docType.SelectedItem(),
		)
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			fmt.Sprintf("\nSearching the %s field in %s documents", m.field.Value(), m.docType.SelectedItem()),
			m.styles.queryField.Render(m.query.View()),
			"",
		)
	case listFields:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			"Press 'enter' to go back to the main menu.",
			formatFieldsList(m.store.ListFields(), m.width),
		)
	default:
		return "Unknown state"
	}
}

func formatFieldsList(fieldsMap map[string][]string, width int) string {
	var out strings.Builder
	for docType, fields := range fieldsMap {
		_, _ = fmt.Fprintf(&out, "Search %s with:\n", docType)
		_, _ = fmt.Fprintf(&out, "%s\n", strings.Repeat("-", width))
		_, _ = fmt.Fprintf(&out, "%s\n", strings.Join(fields, "\n"))
		_, _ = fmt.Fprintln(&out, "")
	}
	return out.String()
}
