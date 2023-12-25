package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func InitialModel() model {
	styles := DefaultStyles()
	query := textinput.New()
	query.Placeholder = "Input a word to search for..."
	query.Focus()
	return model{
		styles: styles,
		query:  query,
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
		// These keys should exit the program.
		case "ctrl+c", "ctrl+d":
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
				newModel.query, cmd = newModel.query.Update(msg)
				return newModel, cmd
			case searchZendeskChosenType:
				newModel.query, cmd = newModel.query.Update(msg)
				return newModel, cmd
			case list:
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
Type 'Ctrl+c' or 'Ctrl+d' to exit at any time, Press 'Enter' to continue.`

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
			headerText,
			searchText,
			m.styles.InputField.Render(m.query.View()),
			"",
		)
	default:
		return "Unknown state"
	}
}
