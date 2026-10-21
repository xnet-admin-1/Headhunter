package kit

import (
	"context"
	"os"
	"path/filepath"
	"strings"
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
