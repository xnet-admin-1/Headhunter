package kit

import (
	"context"
	"fmt"
	"strings"

	"github.com/xnet-admin-1/Headhunter/internal/agents"
)

// DefaultPrimaryAgent is the built-in agent that acts as the primary chat
// agent when the user has not configured a system prompt of their own. Big
// Red runs the engagement and tasks the other specialists, so it is the
// agent that answers by default. The Headhunter binary ships this doctrine,
// so the binary carries its own persona and needs no file to start.
const DefaultPrimaryAgent = "bigred"

// SetPrimaryAgent makes the named agent the primary chat agent. The system
// prompt of the agent replaces the base system prompt of the session. When
// the definition names a model, that model becomes the active model. When
// the definition sets a sampling temperature, that temperature is used.
//
// The tools of the session do not change. A tool allowlist in the definition
// applies to subagents, not to the primary agent.
//
// An unknown name returns an error that lists the known names. A definition
// without a system prompt returns an error, because a primary agent without
// a system prompt has no identity to switch to.
//
// An empty name clears the primary agent. Use ClearPrimaryAgent for that.
func (m *Kit) SetPrimaryAgent(ctx context.Context, name string) error {
	if name == "" {
		m.ClearPrimaryAgent()
		return nil
	}
	definition, ok := m.agentDefinition(name)
	if !ok {
		return fmt.Errorf("unknown agent %q (available: %s)", name, strings.Join(m.knownAgentNames(), ", "))
	}
	if strings.TrimSpace(definition.SystemPrompt) == "" {
		return fmt.Errorf("agent %q has no system prompt", name)
	}

	previousName, previousPrompt, previousTemperature := m.primaryAgentState()

	m.runtimeMu.Lock()
	m.primaryAgent = name
	m.primaryAgentPrompt = definition.SystemPrompt
	m.primaryAgentTemperature = cloneTemperature(definition.Temperature)
	m.runtimeMu.Unlock()

	// A model switch resets the base prompt, so the switch happens before
	// the agent prompt is installed.
	switched := false
	if definition.Model != "" && definition.Model != m.GetModelString() {
		if err := m.SetModel(ctx, definition.Model); err != nil {
			// The persona of the agent is installed in any case. The model
			// of the session stays the one that works, and the caller
			// learns why.
			m.installAgentPrompt(definition.SystemPrompt)
			return fmt.Errorf("switch to the model of agent %q: %w", name, err)
		}
		switched = true
	}

	m.installAgentPrompt(definition.SystemPrompt)

	// The sampling temperature lives in the provider. The provider must be
	// rebuilt for a new temperature to take effect.
	if !switched && definition.Temperature != nil && !sameTemperature(m.currentAppliedTemperature(), definition.Temperature) {
		if err := m.SetModel(ctx, m.GetModelString()); err != nil {
			m.restorePrimaryAgentState(previousName, previousPrompt, previousTemperature)
			return fmt.Errorf("apply the temperature of agent %q: %w", name, err)
		}
	}
	return nil
}

// installAgentPrompt sets the base prompt to the prompt of the agent and
// publishes the composed system prompt.
func (m *Kit) installAgentPrompt(prompt string) {
	m.runtimeMu.Lock()
	m.basePrompt = prompt
	m.runtimeMu.Unlock()
	m.applyComposedSystemPrompt()
}

// ClearPrimaryAgent removes the primary agent. The session uses the
// configured system prompt again.
func (m *Kit) ClearPrimaryAgent() {
	m.runtimeMu.Lock()
	m.primaryAgent = ""
	m.primaryAgentPrompt = ""
	m.primaryAgentTemperature = nil
	m.basePrompt = m.globalBasePrompt
	m.runtimeMu.Unlock()
	m.applyComposedSystemPrompt()
}

// PrimaryAgent returns the name of the agent that acts as the primary chat
// agent. It is empty when the session uses the configured system prompt.
func (m *Kit) PrimaryAgent() string {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return m.primaryAgent
}

// agentDefinition returns the definition of the named agent. The discovered
// agents come first, so a definition file can override a built-in. The
// built-in agents are the fallback, because agent discovery can be off
// (bare mode, the no-agents flag) while the product still needs its own
// persona.
func (m *Kit) agentDefinition(name string) (*AgentDefinition, bool) {
	if definition, ok := m.GetAgent(name); ok {
		return definition, true
	}
	for _, builtin := range agents.Builtins() {
		if builtin.Name == name {
			return cloneAgentDefinition(builtin), true
		}
	}
	return nil, false
}

// knownAgentNames returns the names of every agent that can be selected. It
// lists the discovered agents and the built-in agents.
func (m *Kit) knownAgentNames() []string {
	names := m.agentNames()
	for _, builtin := range agents.Builtins() {
		if !slicesContains(names, builtin.Name) {
			names = append(names, builtin.Name)
		}
	}
	return names
}

// primaryAgentState returns the name, the system prompt and the sampling
// temperature of the primary agent. The values are copies, so the caller
// can keep them.
func (m *Kit) primaryAgentState() (string, string, *float32) {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return m.primaryAgent, m.primaryAgentPrompt, cloneTemperature(m.primaryAgentTemperature)
}

// restorePrimaryAgentState puts back a primary agent state that was saved
// before a failed switch.
func (m *Kit) restorePrimaryAgentState(name, prompt string, temperature *float32) {
	m.runtimeMu.Lock()
	m.primaryAgent = name
	m.primaryAgentPrompt = prompt
	m.primaryAgentTemperature = temperature
	m.runtimeMu.Unlock()
}

// primaryAgentSystemPrompt returns the system prompt of the primary agent.
// It is empty when there is no primary agent.
func (m *Kit) primaryAgentSystemPrompt() string {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return m.primaryAgentPrompt
}

// primaryAgentSamplingTemperature returns the sampling temperature of the
// primary agent. It is nil when there is no primary agent or when the agent
// does not set a temperature.
func (m *Kit) primaryAgentSamplingTemperature() *float32 {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return cloneTemperature(m.primaryAgentTemperature)
}

// currentAppliedTemperature returns the sampling temperature of the
// provider that is in use.
func (m *Kit) currentAppliedTemperature() *float32 {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return cloneTemperature(m.appliedTemperature)
}

// recordAppliedTemperature stores the sampling temperature of the provider
// that was built.
func (m *Kit) recordAppliedTemperature(temperature *float32) {
	m.runtimeMu.Lock()
	m.appliedTemperature = cloneTemperature(temperature)
	m.runtimeMu.Unlock()
}

// cloneTemperature returns a copy of the temperature. It returns nil for a
// nil temperature.
func cloneTemperature(temperature *float32) *float32 {
	if temperature == nil {
		return nil
	}
	value := *temperature
	return &value
}

// sameTemperature reports whether two temperatures are the same. Two nil
// temperatures are the same.
func sameTemperature(a, b *float32) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}

// slicesContains reports whether the slice holds the value.
func slicesContains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
