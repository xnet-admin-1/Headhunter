package ui

import (
	"errors"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestAgentSelectorListsTheAgentsAndTheDefault(t *testing.T) {
	s := NewAgentSelector([]AgentOption{
		{Name: "headhunter", Description: "Contract offensive operator", Badge: "builtin", Model: "custom/deepseek-v4-flash"},
		{Name: "scout", Description: "Recon and mapping", Badge: "builtin"},
	}, "", 200, 30)

	rendered := s.RenderOverlay()
	for _, want := range []string{"Default", "headhunter", "scout", "Contract offensive operator", "builtin"} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("selector does not show %q:\n%s", want, rendered)
		}
	}

	// The popup truncates a long description to the panel width, so the model
	// is checked on the item itself.
	items := s.popup.allItems
	if len(items) != 3 {
		t.Fatalf("items = %d, want the default entry and two agents", len(items))
	}
	if items[0].Meta != "" || items[0].Label != "Default" {
		t.Fatalf("first item = %#v, want the default entry", items[0])
	}
	if items[1].Meta != "headhunter" || !strings.Contains(items[1].Description, "custom/deepseek-v4-flash") {
		t.Fatalf("agent item = %#v", items[1])
	}
}

func TestAgentSelectorMarksTheCurrentAgent(t *testing.T) {
	s := NewAgentSelector([]AgentOption{{Name: "scout", Current: true}}, "scout", 100, 30)
	if !strings.Contains(s.RenderOverlay(), "scout") {
		t.Fatal("the current agent is missing")
	}
}

func TestAgentSelectorStartsOnTheCurrentAgent(t *testing.T) {
	s := NewAgentSelector([]AgentOption{{Name: "scout"}}, "scout", 100, 30)
	_, cmd := s.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("Enter must select an entry")
	}
	msg, ok := cmd().(AgentSelectedMsg)
	if !ok {
		t.Fatalf("selection is %#v", msg)
	}
	if msg.Name != "scout" {
		t.Fatalf("name = %q, want the current agent", msg.Name)
	}
}

func TestAgentSelectorDefaultEntryClearsTheAgent(t *testing.T) {
	s := NewAgentSelector([]AgentOption{{Name: "scout"}}, "", 100, 30)
	_, cmd := s.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("Enter must select an entry")
	}
	msg, ok := cmd().(AgentSelectedMsg)
	if !ok {
		t.Fatalf("selection is %#v", msg)
	}
	if msg.Name != "" {
		t.Fatalf("name = %q, want the default entry", msg.Name)
	}
}

func TestAgentSelectorCancelReportsACancel(t *testing.T) {
	s := NewAgentSelector([]AgentOption{{Name: "scout"}}, "", 100, 30)
	_, cmd := s.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	if cmd == nil {
		t.Fatal("Escape must cancel the selector")
	}
	if _, ok := cmd().(AgentSelectorCancelledMsg); !ok {
		t.Fatalf("cancel is %#v", cmd())
	}
}

func TestAgentCommandOpensTheSelector(t *testing.T) {
	m := &AppModel{
		listAgents: func() []AgentOption {
			return []AgentOption{{Name: "scout", Description: "Recon"}}
		},
		setPrimaryAgent: func(string) error { return nil },
	}
	m.handleAgentCommand("")
	if m.state != stateAgentSelector {
		t.Fatalf("state = %v, want the agent selector", m.state)
	}
	if m.agentSelector == nil {
		t.Fatal("the selector is missing")
	}
}

func TestAgentCommandSetsTheAgentDirectly(t *testing.T) {
	chosen := ""
	m := &AppModel{
		listAgents: func() []AgentOption { return nil },
		setPrimaryAgent: func(name string) error {
			chosen = name
			return nil
		},
	}
	m.handleAgentCommand("scout")
	if chosen != "scout" {
		t.Fatalf("chosen agent = %q", chosen)
	}
	if m.primaryAgent != "scout" {
		t.Fatalf("primary agent = %q", m.primaryAgent)
	}
}

func TestAgentCommandReportsAFailure(t *testing.T) {
	m := &AppModel{
		listAgents:      func() []AgentOption { return nil },
		setPrimaryAgent: func(string) error { return errors.New("unknown agent") },
	}
	m.handleAgentCommand("scout")
	if m.primaryAgent != "" {
		t.Fatalf("primary agent = %q, want it unchanged", m.primaryAgent)
	}
}

func TestAgentCommandWithoutSupportReportsIt(t *testing.T) {
	m := &AppModel{}
	m.handleAgentCommand("")
	if m.state == stateAgentSelector {
		t.Fatal("the selector must not open without support")
	}
}
