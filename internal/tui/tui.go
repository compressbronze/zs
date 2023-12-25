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
	searchZendeskChosenType
	list
)

type model struct {
	state         state
	styles        *Styles
	width, height int
	query         textinput.Model
	store         stores.Store
}

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := &Styles{}
	s.BorderColor = lipgloss.Color("#a134eb")
	s.InputField = lipgloss.
		NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).
		Width(50)
	return s
}

func InitialModel(store stores.Store) model {
	styles := DefaultStyles()
	query := textinput.New()
	query.Placeholder = "Input a word to search for..."
	query.Focus()
	return model{
		styles: styles,
		query:  query,
		store:  store,
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
		switch msg.String() {
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
				switch msg.String() {
				case "1":
					newModel.state = searchZendesk
				case "2":
					newModel.state = list
				}
				return newModel, nil
			case searchZendesk:
				switch msg.String() {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.query, cmd = newModel.query.Update("")
					return newModel, cmd
				case "enter":
				}
				newModel.query, cmd = newModel.query.Update(msg)
				return newModel, cmd
			case searchZendeskChosenType:
				switch msg.String() {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.query, cmd = newModel.query.Update("")
					return newModel, cmd
				case "enter":
				}
				newModel.query, cmd = newModel.query.Update(msg)
				return newModel, cmd
			case list:
				switch msg.String() {
				case "ctrl+d":
					newModel.state = selectOptions
					newModel.query, cmd = newModel.query.Update("")
					return newModel, cmd
				case "enter":
					newModel.state = selectOptions
					return newModel, cmd
				}
				newModel.query, cmd = newModel.query.Update(msg)
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

	headerText = fmt.Sprintf("%d\n%s", m.state, headerText)

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
			m.styles.InputField.Render(m.query.View()),
			"",
		)
	case searchZendeskChosenType:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			headerText,
			searchText,
			m.styles.InputField.Render(m.query.View()),
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
