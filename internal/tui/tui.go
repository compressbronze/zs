package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
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

type styles struct {
	docType, field, query, fieldsList, results lipgloss.Style
}

func DefaultStyles() *styles {
	return &styles{
		docType: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#154733")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
		field: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#ed095d")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
		query: lipgloss.
			NewStyle().
			BorderForeground(lipgloss.Color("#a134eb")).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1).
			Width(80),
		results: lipgloss.
			NewStyle().
			Padding(0, 1).
			BorderForeground(lipgloss.Color("#a134eb")).
			BorderStyle(lipgloss.RoundedBorder()),
		fieldsList: lipgloss.
			NewStyle().
			Padding(0, 1).
			BorderForeground(lipgloss.Color("#ed095d")).
			BorderStyle(lipgloss.RoundedBorder()),
	}
}

type model struct {
	state         state
	styles        *styles
	width, height int
	docType       doctypelist.Model
	field, query  textinput.Model
	resultsErr    error
	store         stores.Store
	veiwport      viewport.Model
	quitting      bool
}

func InitialModel(store stores.Store) model {
	styles := DefaultStyles()
	docTypes := store.ListDocumentTypes()
	docType := doctypelist.New(docTypes)
	field := textinput.New()
	query := textinput.New()
	query.ShowSuggestions = true

	return model{
		styles:   styles,
		docType:  docType,
		field:    field,
		query:    query,
		store:    store,
		veiwport: viewport.New(0, 0),
	}
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Clear() (model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := make([]tea.Cmd, 0, 4)
	m.docType, cmd = m.docType.Update("")
	cmds = append(cmds, cmd)
	m.query, cmd = m.query.Update("")
	cmds = append(cmds, cmd)
	m.field, cmd = m.field.Update("")
	cmds = append(cmds, cmd)
	m.veiwport.SetContent("")

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

		m.veiwport.Width = m.width - 4
		switch m.state {
		case listFields, header, selectOptions, search, chosenDocType, chosenDocTypeField:
			m.veiwport.Height = m.height - 4
		case results:
			m.veiwport.Height = m.height - 5
		}

		return m, nil
	case tea.KeyMsg:
		switch s := msg.String(); s {
		// ctrl+c should exit the program from any state.
		case "ctrl+c":
			m.quitting = true
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
					m, cmd = m.Clear()
					if cmd != nil {
						return m, cmd
					}
					m.veiwport.Width = m.width - 4
					m.veiwport.Height = m.height - 4
					m.veiwport.SetContent(formatFieldsList(m.store.ListFields(), m.veiwport.Width))
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
					resultDocs, err := m.store.Search(
						m.docType.SelectedItem(),
						m.field.Value(),
						m.query.Value(),
					)
					if err != nil {
						m.state = results
						m.resultsErr = err
						return m, nil
					}

					m, cmd = m.Clear()
					if cmd != nil {
						return m, cmd
					}
					m.veiwport.Width = m.width - 4
					m.veiwport.Height = m.height - 5
					formattedResults, err := formatResults(resultDocs, m.veiwport.Width)
					if err != nil {
						m.state = results
						m.resultsErr = err
						return m, nil
					}
					m.veiwport.SetContent(formattedResults)
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
					m.veiwport, cmd = m.veiwport.Update(msg)
					return m, cmd
				}
			case listFields:
				switch s {
				case "enter":
					m.state = selectOptions
					return m.Clear()
				default:
					m.veiwport, cmd = m.veiwport.Update(msg)
					return m, cmd
				}
			}
		}
	}

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

	s := func() string {
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
				m.styles.docType.Render(m.docType.View()),
			)
		case chosenDocType:
			return lipgloss.JoinVertical(
				lipgloss.Left,
				headerText,
				fmt.Sprintf("\nSearching %s documents", m.docType.SelectedItem()),
				m.styles.field.Render(m.field.View()),
			)
		case chosenDocTypeField:
			return lipgloss.JoinVertical(
				lipgloss.Left,
				headerText,
				fmt.Sprintf("\nSearching the %s field in %s documents", m.field.Value(), m.docType.SelectedItem()),
				m.styles.query.Render(m.query.View()),
			)
		case results:
			if m.resultsErr != nil {
				return lipgloss.JoinVertical(
					lipgloss.Left,
					headerText,
					"\nError searching for documents:",
					m.resultsErr.Error(),
					"Press 'enter' to go back to the main menu.",
				)
			}
			return lipgloss.JoinVertical(
				lipgloss.Left,
				"Found the following documents:",
				"Press 'enter' to go back to the main menu.",
				m.styles.results.Render(m.veiwport.View()),
			)
		case listFields:
			return lipgloss.JoinVertical(
				lipgloss.Left,
				"Press 'enter' to go back to the main menu.",
				m.styles.fieldsList.Render(m.veiwport.View()),
			)
		default:
			return "Unknown state"
		}
	}()

	// So that the prompt does no overwrite the last line.
	// See https://github.com/charmbracelet/bubbletea/issues/304
	if m.quitting {
		return s + "\n"
	}
	return s
}

func formatResults(results []models.Model, width int) (string, error) {
	var out strings.Builder
	for _, result := range results {
		_, _ = fmt.Fprintf(&out, "%s\n", result.DocumentType())
		_, _ = fmt.Fprintf(&out, "%s\n", strings.Repeat("-", width))

		buf, err := models.StringOf(result)
		if err != nil {
			return "", fmt.Errorf("failed to string value: %w", err)
		}
		_, _ = fmt.Fprintf(&out, "%s\n", buf)
		_, _ = fmt.Fprintln(&out, "")
	}
	return out.String(), nil
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
