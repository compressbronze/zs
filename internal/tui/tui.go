package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/satrap-illustrations/zs/internal/models"
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
	results       []models.Model
	resultsErr    error
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
	query := textinput.New()
	query.ShowSuggestions = true
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

func (m model) Clear() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := make([]tea.Cmd, 0, 3)
	m.docType, cmd = m.docType.Update("")
	cmds = append(cmds, cmd)
	m.query, cmd = m.query.Update("")
	cmds = append(cmds, cmd)
	m.field, cmd = m.field.Update("")
	cmds = append(cmds, cmd)

	m.results = nil
	m.resultsErr = nil

	return m, tea.Batch(cmds...)
}

//nolint:revive
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch s := msg.String(); s {
		// ctrl+c should exit the program from any state.
		case "ctrl+c":
			return m, tea.Quit

		default:
			switch m.state {
			case header:
				switch s {
				case "enter":
					m.state = selectOptions
					return m, nil
				default:
					return m, nil
				}
			case selectOptions:
				switch s {
				case "1":
					m.state = search
				case "2":
					m.state = listFields
				}
				return m, nil
			case search:
				switch s {
				case "ctrl+d":
					m.state = selectOptions
					return m.Clear()
				case "enter":
					m.state = chosenDocType
					m.field.Placeholder = fmt.Sprintf(
						"Type a field to search in %s...",
						m.docType.SelectedItem(),
					)
					m.field.SetSuggestions(m.store.ListFields()[m.docType.SelectedItem()])
					m.field.Focus()
					return m, nil
				default:
					m.docType, cmd = m.docType.Update(msg)
					return m, cmd
				}
			case chosenDocType:
				switch s {
				case "ctrl+d":
					m.state = selectOptions
					return m.Clear()
				case "enter":
					m.state = chosenDocTypeField
					m.query.Placeholder = fmt.Sprintf(
						"Type a value of %s in %s to search for...",
						m.field.Value(),
						m.docType.SelectedItem(),
					)
					m.query.Focus()
					return m, cmd
				default:
					m.field, cmd = m.field.Update(msg)
					return m, cmd
				}
			case chosenDocTypeField:
				switch s {
				case "ctrl+d":
					m.state = selectOptions
					return m.Clear()
				case "enter":
					m.results, m.resultsErr = m.store.Search(
						m.docType.SelectedItem(),
						m.field.Value(),
						m.query.Value(),
					)
					if m.resultsErr != nil {
						m.state = results
						return m, nil
					}
					m.state = results
					return m, nil
				default:
					m.query, cmd = m.query.Update(msg)
					return m, cmd
				}
			case results:
				switch s {
				case "enter":
					m.state = selectOptions
					return m.Clear()
				default:
					return m, nil
				}
			case listFields:
				switch s {
				case "enter":
					m.state = selectOptions
					return m, nil
				default:
					return m, nil
				}
			}
		}
	}

	m.query, cmd = m.query.Update(msg)
	return m, cmd
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
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			fmt.Sprintf("\nSearching %s documents", m.docType.SelectedItem()),
			m.styles.fieldField.Render(m.field.View()),
			"",
		)
	case chosenDocTypeField:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			fmt.Sprintf("\nSearching the %s field in %s documents", m.field.Value(), m.docType.SelectedItem()),
			m.styles.queryField.Render(m.query.View()),
			"",
		)
	case results:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			"\nFound the following documents:",
			formatResults(m.results),
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

func formatResults(results []models.Model) string {
	var out strings.Builder
	for _, result := range results {
		buf, err := models.StringOf(result)
		if err != nil {
			return fmt.Sprintf("failed to string value: %s", err)
		}
		_, _ = fmt.Fprintf(&out, "%s\n", buf)
	}
	return out.String()
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
