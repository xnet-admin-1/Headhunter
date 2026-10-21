package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

// AgentOption describes a named agent that can act as the primary chat
// agent.
type AgentOption struct {
	// Name identifies the agent. It is the name that the subagent tool and
	// the /agent command use.
	Name string

	// Description summarises what the agent does. It is the description of
	// the agent definition.
	Description string

	// Badge marks where the definition came from, e.g. "builtin" or
	// "user". It is empty when the source is not worth showing.
	Badge string

	// Model is the model of the definition. It is empty when the agent
	// does not name a model and keeps the model of the session.
	Model string

	// Current is true for the agent that is the primary chat agent now.
	Current bool
}

// AgentSelectedMsg reports the agent that the user chose.
type AgentSelectedMsg struct{ Name string }

// AgentSelectorCancelledMsg reports that the user closed the selector.
type AgentSelectorCancelledMsg struct{}

// AgentSelectorComponent is the popup that lists the agents that can act as
// the primary chat agent.
type AgentSelectorComponent struct {
	popup         *PopupList
	width, height int
}

// NewAgentSelector builds the agent selector. The options are listed in the
// order they are given. The first entry is always the default, which uses
// the configured system prompt.
func NewAgentSelector(options []AgentOption, current string, width, height int) *AgentSelectorComponent {
	items := make([]PopupItem, 0, len(options)+1)
	items = append(items, PopupItem{
		Label:       "Default",
		Description: "Use the configured system prompt",
		Active:      current == "",
		Meta:        "",
	})
	for _, option := range options {
		description := option.Description
		if option.Model != "" {
			description = fmt.Sprintf("%s · model %s", description, option.Model)
		}
		items = append(items, PopupItem{
			Label:       option.Name,
			Badge:       option.Badge,
			Description: description,
			Active:      option.Name == current,
			Meta:        option.Name,
		})
	}
	popup := NewPopupList("Primary Agent", items, width, height)
	popup.Subtitle = "Select the agent that answers in this session"
	return &AgentSelectorComponent{popup: popup, width: width, height: height}
}

func (s *AgentSelectorComponent) Init() tea.Cmd { return nil }

func (s *AgentSelectorComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width, s.height = msg.Width, msg.Height
		s.popup.SetSize(msg.Width, msg.Height)
	case tea.KeyPressMsg:
		result := s.popup.HandleKey(msg.String(), msg.Text)
		if result.Selected != nil {
			name, _ := result.Selected.Meta.(string)
			return s, func() tea.Msg { return AgentSelectedMsg{Name: name} }
		}
		if result.Cancelled {
			return s, func() tea.Msg { return AgentSelectorCancelledMsg{} }
		}
	}
	return s, nil
}

func (s *AgentSelectorComponent) View() tea.View {
	v := tea.NewView(s.popup.RenderCentered(s.width, s.height))
	v.AltScreen = true
	return v
}

func (s *AgentSelectorComponent) RenderOverlay() string { return s.popup.Render() }
