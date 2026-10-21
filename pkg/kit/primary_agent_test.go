package kit

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// newPrimaryAgentTestKit builds a Kit that loads named agents from a
// temporary user agent directory. The map holds the file content of each
// agent definition. It returns the kit and the config home.
func newPrimaryAgentTestKit(t *testing.T, files map[string]string) (*Kit, string) {
	t.Helper()
	configHome := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", configHome)
	dir := filepath.Join(configHome, "kit", "agents")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatalf("write agent: %v", err)
		}
	}
	recorder := &recordingFactory{}
	k := newProviderTestKit(t, &Options{
		Model:     "local/tiny-1b",
		Providers: map[string]ProviderFactory{"local": recorder.factory("local")},
	})
	return k, configHome
}

// testAgentFile is a definition file with a system prompt.
func testAgentFile(prompt string) string {
	return "---\ndescription: Test agent\n---\n" + prompt + "\n"
}

func TestSetPrimaryAgentInstallsTheAgentPrompt(t *testing.T) {
	k, _ := newPrimaryAgentTestKit(t, map[string]string{
		"tester.md": testAgentFile("You are the tester."),
	})

	if k.PrimaryAgent() != "" {
		t.Fatalf("primary agent at start = %q", k.PrimaryAgent())
	}

	if err := k.SetPrimaryAgent(context.Background(), "tester"); err != nil {
		t.Fatalf("SetPrimaryAgent: %v", err)
	}
	if k.PrimaryAgent() != "tester" {
		t.Fatalf("primary agent = %q", k.PrimaryAgent())
	}
	if prompt := k.GetSystemPrompt(); !strings.Contains(prompt, "You are the tester.") {
		t.Fatalf("system prompt does not hold the agent prompt: %q", prompt)
	}
}

func TestSetPrimaryAgentRejectsAnUnknownName(t *testing.T) {
	k, _ := newPrimaryAgentTestKit(t, nil)

	err := k.SetPrimaryAgent(context.Background(), "missing")
	if err == nil {
		t.Fatal("an unknown agent must fail")
	}
	if !strings.Contains(err.Error(), "missing") {
		t.Fatalf("error = %v", err)
	}
	if k.PrimaryAgent() != "" {
		t.Fatalf("primary agent = %q", k.PrimaryAgent())
	}
}

func TestSetPrimaryAgentRejectsAnAgentWithoutPrompt(t *testing.T) {
	k, _ := newPrimaryAgentTestKit(t, map[string]string{
		"empty.md": "---\ndescription: No prompt\n---\n",
	})

	err := k.SetPrimaryAgent(context.Background(), "empty")
	if err == nil {
		t.Fatal("an agent without a system prompt must fail")
	}
	if !strings.Contains(err.Error(), "no system prompt") {
		t.Fatalf("error = %v", err)
	}
}

func TestClearPrimaryAgentRestoresTheConfiguredPrompt(t *testing.T) {
	k, _ := newPrimaryAgentTestKit(t, map[string]string{
		"tester.md": testAgentFile("You are the tester."),
	})
	before := k.GetSystemPrompt()

	if err := k.SetPrimaryAgent(context.Background(), "tester"); err != nil {
		t.Fatalf("SetPrimaryAgent: %v", err)
	}
	if err := k.SetPrimaryAgent(context.Background(), ""); err != nil {
		t.Fatalf("clear primary agent: %v", err)
	}
	if k.PrimaryAgent() != "" {
		t.Fatalf("primary agent = %q", k.PrimaryAgent())
	}
	if after := k.GetSystemPrompt(); after != before {
		t.Fatalf("system prompt was not restored:\nbefore %q\nafter  %q", before, after)
	}
}

// TestSetPrimaryAgentUsesTheBuiltinWithoutDiscovery checks that the built-in
// agent works as the primary agent when the agent discovery is off. The
// product needs its own persona in every mode.
func TestSetPrimaryAgentUsesTheBuiltinWithoutDiscovery(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	recorder := &recordingFactory{}
	k := newProviderTestKit(t, &Options{
		Model:     "local/tiny-1b",
		Bare:      true,
		Providers: map[string]ProviderFactory{"local": recorder.factory("local")},
	})

	if err := k.SetPrimaryAgent(context.Background(), DefaultPrimaryAgent); err != nil {
		t.Fatalf("SetPrimaryAgent: %v", err)
	}
	if k.PrimaryAgent() != DefaultPrimaryAgent {
		t.Fatalf("primary agent = %q", k.PrimaryAgent())
	}
	if prompt := k.GetSystemPrompt(); !strings.Contains(prompt, "BIG RED") {
		t.Fatalf("system prompt does not hold the built-in persona: %q", prompt)
	}
	// The built-in agents do not pin a model, so the model of the session
	// stays the one that works.
	if k.GetModelString() != "local/tiny-1b" {
		t.Fatalf("model = %q", k.GetModelString())
	}
}

// TestSetPrimaryAgentKeepsThePromptOnAModelSwitch checks that a model change
// does not take the persona of the agent away.
func TestSetPrimaryAgentKeepsThePromptOnAModelSwitch(t *testing.T) {
	k, _ := newPrimaryAgentTestKit(t, map[string]string{
		"tester.md": testAgentFile("You are the tester."),
	})

	if err := k.SetPrimaryAgent(context.Background(), "tester"); err != nil {
		t.Fatalf("SetPrimaryAgent: %v", err)
	}
	if err := k.SetModel(context.Background(), "local/other"); err != nil {
		t.Fatalf("SetModel: %v", err)
	}
	if prompt := k.GetSystemPrompt(); !strings.Contains(prompt, "You are the tester.") {
		t.Fatalf("the persona was lost on the model switch: %q", prompt)
	}
	if k.PrimaryAgent() != "tester" {
		t.Fatalf("primary agent = %q", k.PrimaryAgent())
	}
}

// temperatureFactory records the sampling temperature of every provider it
// builds.
type temperatureFactory struct {
	mu           sync.Mutex
	temperatures []*float32
}

func (f *temperatureFactory) factory(provider string) ProviderFactory {
	return func(_ context.Context, config *ProviderConfig, modelName string) (*ProviderResult, error) {
		f.mu.Lock()
		f.temperatures = append(f.temperatures, cloneTemperature(config.Temperature))
		f.mu.Unlock()
		return &ProviderResult{
			Model:  LLMLanguageModel(&echoModel{provider: provider, model: modelName}),
			Closer: closerFunc(func() error { return nil }),
		}, nil
	}
}

func (f *temperatureFactory) last() *float32 {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.temperatures) == 0 {
		return nil
	}
	return f.temperatures[len(f.temperatures)-1]
}

// TestSetPrimaryAgentAppliesTheTemperature checks that the sampling
// temperature of the agent reaches the provider.
func TestSetPrimaryAgentAppliesTheTemperature(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	dir := filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "kit", "agents")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	content := "---\ndescription: Cold agent\ntemperature: 0.2\n---\nYou are cold.\n"
	if err := os.WriteFile(filepath.Join(dir, "cold.md"), []byte(content), 0o644); err != nil {
		t.Fatalf("write agent: %v", err)
	}
	factory := &temperatureFactory{}
	k := newProviderTestKit(t, &Options{
		Model:     "local/tiny-1b",
		Providers: map[string]ProviderFactory{"local": factory.factory("local")},
	})

	if err := k.SetPrimaryAgent(context.Background(), "cold"); err != nil {
		t.Fatalf("SetPrimaryAgent: %v", err)
	}
	temperature := factory.last()
	if temperature == nil || *temperature != 0.2 {
		t.Fatalf("provider temperature = %v, want 0.2", temperature)
	}
}
