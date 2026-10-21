package kit

import (
	"context"
	"fmt"
	"strings"
)

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
	definition, ok := m.GetAgent(name)
	if !ok {
		return fmt.Errorf("unknown agent %q (available: %s)", name, strings.Join(m.agentNames(), ", "))
	}
	if strings.TrimSpace(definition.SystemPrompt) == "" {
		return fmt.Errorf("agent %q has no system prompt", name)
	}

	previousName, previousTemperature := m.primaryAgentState()

	m.runtimeMu.Lock()
	m.primaryAgent = name
	m.primaryAgentTemperature = cloneTemperature(definition.Temperature)
	m.runtimeMu.Unlock()

	// A model switch resets the base prompt to the configured one, so the
	// switch happens before the agent prompt is installed.
	if definition.Model != "" && definition.Model != m.GetModelString() {
		if err := m.SetModel(ctx, definition.Model); err != nil {
			m.restorePrimaryAgentState(previousName, previousTemperature)
			return fmt.Errorf("switch to the model of agent %q: %w", name, err)
		}
	}

	m.runtimeMu.Lock()
	m.basePrompt = definition.SystemPrompt
	m.runtimeMu.Unlock()
	m.applyComposedSystemPrompt()
	return nil
}

// ClearPrimaryAgent removes the primary agent. The session uses the
// configured system prompt again.
func (m *Kit) ClearPrimaryAgent() {
	m.runtimeMu.Lock()
	m.primaryAgent = ""
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

// primaryAgentState returns the name and the sampling temperature of the
// primary agent. The temperature is a copy, so the caller can keep it.
func (m *Kit) primaryAgentState() (string, *float32) {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return m.primaryAgent, cloneTemperature(m.primaryAgentTemperature)
}

// restorePrimaryAgentState puts back a primary agent state that was saved
// before a failed switch.
func (m *Kit) restorePrimaryAgentState(name string, temperature *float32) {
	m.runtimeMu.Lock()
	m.primaryAgent = name
	m.primaryAgentTemperature = temperature
	m.runtimeMu.Unlock()
}

// primaryAgentSamplingTemperature returns the sampling temperature of the
// primary agent. It is nil when there is no primary agent or when the agent
// does not set a temperature.
func (m *Kit) primaryAgentSamplingTemperature() *float32 {
	m.runtimeMu.RLock()
	defer m.runtimeMu.RUnlock()
	return cloneTemperature(m.primaryAgentTemperature)
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
