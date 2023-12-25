package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/satrap-illustrations/zs/internal/stores"
)

type state int

const (
	header state = iota
	selectOptions
	searchZendesk
	searchZendeskDocType
	list
)

type model struct {
	state         state
	styles        *styles
	width, height int
	query         textinput.Model
	docType       textinput.Model
	store         stores.Store
}

type styles struct {
	docTypeField lipgloss.Style
	queryField   lipgloss.Style
}

func DefaultStyles() *styles {
	s := &styles{}
	s.docTypeField = lipgloss.
		NewStyle().
		BorderForeground(lipgloss.Color("#154733")).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(50)
	s.queryField = lipgloss.
		NewStyle().
		BorderForeground(lipgloss.Color("#a134eb")).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(50)
	return s
}

func InitialModel(store stores.Store) model {
	styles := DefaultStyles()
	docType := textinput.New()
	docType.Placeholder = "Input a document type to search..."

	query := textinput.New()
	query.Placeholder = "Input a word to search for..."
	return model{
		styles:  styles,
		docType: docType,
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
				if msg.String() == "enter" {
					newModel.state = selectOptions
				}
				return newModel, nil
			case selectOptions:
				switch s {
				case "1":
					newModel.state = searchZendesk
					newModel.docType.Focus()
				case "2":
					newModel.state = list
				}
				return newModel, nil
			case searchZendesk:
				switch s {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.docType, cmd = newModel.docType.Update("")
					return newModel, cmd
				case "enter":
					newModel.state = searchZendeskDocType
					newModel.query.Focus()
					return newModel, nil
				}
				newModel.docType, cmd = newModel.docType.Update(msg)
				return newModel, cmd
			case searchZendeskDocType:
				switch s {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.query, cmd = newModel.query.Update("")
					return newModel, cmd
				case "enter":
				}
				newModel.query, cmd = newModel.query.Update(msg)
				return newModel, cmd
			case list:
				if s == "enter" {
					newModel.state = selectOptions
					return newModel, cmd
				}
				return newModel, cmd
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

	headerText := `
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

	searchText := `
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
	case searchZendesk:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			searchText,
			m.styles.docTypeField.Render(m.docType.View()),
			"",
		)
	case searchZendeskDocType:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			searchText,
			m.styles.queryField.Render(m.query.View()),
			"",
		)
	case list:
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
