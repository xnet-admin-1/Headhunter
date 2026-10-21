<p align="center">
  <a href="https://github.com/xnet-admin-1/Headhunter/actions/workflows/ci.yml"><img src="https://github.com/xnet-admin-1/Headhunter/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/xnet-admin-1/Headhunter/releases/latest"><img src="https://img.shields.io/github/v/release/xnet-admin-1/Headhunter?style=flat&color=blue" alt="Release"></a>
  <a href="https://www.npmjs.com/package/@xnet-admin-1/headhunter"><img src="https://img.shields.io/npm/v/@xnet-admin-1/headhunter?style=flat&color=cb3837" alt="npm"></a>
  <a href="https://pkg.go.dev/github.com/xnet-admin-1/Headhunter"><img src="https://pkg.go.dev/badge/github.com/xnet-admin-1/Headhunter.svg" alt="Go Reference"></a>
  <a href="https://github.com/xnet-admin-1/Headhunter/blob/master/LICENSE"><img src="https://img.shields.io/github/license/xnet-admin-1/Headhunter?style=flat" alt="License"></a>
</p>

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║                            [  TARGET ACQUIRED  ]                             ║
║                                                                              ║
║     █   █  █████   ███   ████   █   █  █   █  █   █  █████  █████  ████      ║
║     █   █  █      █   █  █   █  █   █  █   █  ██  █    █    █      █   █     ║
║     █   █  █      █   █  █   █  █   █  █   █  █ █ █    █    █      █   █     ║
║     █████  ████   █████  █   █  █████  █   █  █  ██    █    ████   ████      ║
║      █   █  █      █   █  █   █  █   █  █   █  █   █    █    █      █ █      ║
║     █   █  █      █   █  █   █  █   █  █   █  █   █    █    █      █  █      ║
║     █   █  █████  █   █  ████   █   █   ███   █   █    █    █████  █   █     ║
║                                                                              ║
║                    [  RECON · EXPLOIT · PIVOT · PROVE  ]                     ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

# HEADHUNTER

> Lock on. Line up the shot. Take the domain.

**HEADHUNTER is a red-team offensive agent platform.** It is the command brain of the kill chain — a multi-agent AI strike package that hunts from first recon to the last shell you hold. It is built to be aimed at an adversary's network, coupled to a Kali tool server, and its named agents are the specialist operators who walk the chain end to end.

## The strike package

Headhunter is not one agent: it is a task force. Under a lead operator running overwatch, each named agent carries doctrine, a tool allowlist, and a mission:

| Agent | Role |
|-------|------|
| **Big Red** | Command — overwatch, op board, coordination, tasking |
| **Headhunter** | Point operator — hunts the kill chain: recon, exploit, pivot, prove |
| **Scout** | Eyes and ears — reachability, ports, banners, exposed seams |
| **Hands** | Fabricator — builds honeypots, sensors, proxies, custom tooling, and proofs it |
| **Doc** | Combat medic — detects infections, patches holes, cleans compromised hosts |
| **Agent X** | Spymaster — the human side: social engineering, cover, quiet work |

All of it rides one infrastructure spine: an agent client that speaks MCP and carries arbitrary tools, so the operators above reach a live offensive toolset instead of a sandbox.

## Coupled with zebbern-kali-mcp

Headhunter is designed to be fired through a Kali MCP server. It ships to find [zebbern-kali-mcp](https://github.com/zebbern/zebbern-kali-mcp) — the capability-aware MCP client for the Docker-hosted Kali Linux tools API — as its primary tool surface.

- **Repository:** <https://github.com/zebbern/zebbern-kali-mcp>
- Configure it under `.kit.yml`:
  ```yaml
  mcpServers:
    kali:
      type: local
      command: [/path/to/kali-mcp.sh]   # token-injected zebbern-kali-mcp launcher
  ```
- That unlocks the full `kali.*` surface — nmap, masscan, msfconsole, payloads, SSH/Chisel/Ligolo pivots, reverse-shell callbacks, AD tooling — as tools Headhunter's agents call directly.

With the Kali server on the wire, Headhunter goes from a coding agent to a live red-team operator: discover the attack surface, take the foothold, pivot through the network, and hold the shell.

## Under the hood

Beyond the operator frame, Headhunter is an extensible, multi-provider agent CLI (see [Features](#features)): named-agent presets with per-agent tool allowlists, an interpreted-Go extension system, MCP client integration, a tree-based session model, detachable/remote daemon sessions, and a public Go SDK — everything a strike package needs to be aimed, instrumented, and re-built in the field.

## Features

- **Named-Agent Strike Package**: a pre-assembled red-team roster (Big Red, Headhunter, Scout, Hands, Doc, Agent X) delivered as markdown presets with per-agent tool allowlists — taskable, replaceable, and hidden from the catalog on demand
- **Multi-Provider LLM Support**: Anthropic, OpenAI, Google Gemini, Ollama, Azure OpenAI, AWS Bedrock, OpenRouter, and more
- **Built-in Core Tools**: shell (configurable shell, bash by default, with interactive sudo password prompt), read, write, edit, grep, find, ls, subagent - no MCP overhead
- **Named Agents**: Reusable subagent presets defined in markdown with per-agent tool allowlists, advertised to the LLM for delegation. A built-in status bar above parent activity shows active runs and recent results
- **Smart @ Attachments**: Binary files auto-detected via MIME type, MCP resources via `@mcp:server:uri`
- **MCP Integration**: Connect external MCP servers for expanded capabilities
- **Extension System**: Write custom tools, commands, widgets, and UI modifications in Go
- **Theming**: 22 built-in color themes (KITT, Catppuccin, Dracula, Nord, etc.) with runtime switching, persistence, and custom theme files
- **Model Persistence**: Model and thinking level selections are automatically saved and restored across sessions
- **Prompt Templates**: Create reusable prompt templates with shell-style argument substitution
- **Interactive TUI**: Rich terminal interface powered by Bubble Tea with streaming, syntax highlighting, and custom rendering
- **Session Management**: Tree-based conversation history with branching support
- **Non-Interactive Mode**: Script-friendly positional args with JSON output
- **GitHub Integration**: Scaffold a GitHub Actions workflow with `headhunter github install` to run Headhunter as a collaborator/reviewer on `/kit` comments
- **ACP Server**: Run Headhunter as an [Agent Client Protocol](https://agentclientprotocol.com) agent over stdio
- **Sessions**: with the daemon running, a plain `headhunter` is already a detachable session in the directory you ran it from (tmux-style: `Ctrl+]` `d` detach, `s` switch, `c` new, `n`/`p` cycle); `headhunter session attach` picks which session to return to, across this machine *and* every paired host, and `headhunter session ls` lists them
- **Remote Sessions**: `headhunter daemon` on one machine; pair a client once with `headhunter daemon pair` + `headhunter remote --pair <code>` (accept/reject on the host), then connect any time with `headhunter remote --host <name>` — sessions can be detached (`Ctrl+] d`) and reattached from the in-client session list, or shared by several clients at once (tmux-style mirrored view) — end-to-end encrypted iroh transport, revocable public-key credentials, clipboard image paste from the client machine, systemd service support
- **Go SDK**: Embed Headhunter in your own applications with full agent lifecycle events (30+ event types) and behavior-modifying hooks

## Installation

### Using the install script (macOS / Linux)

```bash
curl -fsSL https://raw.githubusercontent.com/xnet-admin-1/Headhunter/master/install.sh | bash
```

The script downloads the latest release binary for your platform, verifies its
SHA-256 checksum, and installs it to `~/.local/bin` (or another writable
directory on your system). Options:

```bash
# Install a specific version into a custom directory
curl -fsSL https://raw.githubusercontent.com/xnet-admin-1/Headhunter/master/install.sh | bash -s -- --version v0.111.0 --bin-dir /usr/local/bin
```

| Option | Description |
|--------|-------------|
| `--bin-dir <dir>` | Install directory (env: `KIT_BIN_DIR`) |
| `--version <tag>` | Release tag to install, default is latest (env: `KIT_VERSION`) |
| `--no-provider-check` | Skip the LLM provider API key reminder |

### Using npm / bun / pnpm

```bash
npm install -g @xnet-admin-1/headhunter
# or
bun install -g @xnet-admin-1/headhunter
# or
pnpm install -g @xnet-admin-1/headhunter
```

### Using Nix

```bash
nix profile install github:xnet-admin-1/Headhunter --no-write-lock-file
```

The flake ships no lock file: its only inputs are `nixpkgs` and the pinned
release, and an unlocked `nixpkgs` is what keeps the build working as kit asks
for newer Go. Nix cannot create a lock file inside the read-only source it
fetched, so an install without `--no-write-lock-file` fails with `cannot
write modified lock file`.

This installs the newest tagged release, built from source for linux and macOS
on amd64 and arm64. To install a specific release instead, name its tag:

```bash
nix profile install github:xnet-admin-1/Headhunter/v0.111.0 --no-write-lock-file
```

Run `nix develop` in a clone to hack on kit with a Nix-provided toolchain: it
puts `go` (the version `go.mod` asks for), `gopls`, and `golangci-lint` on
`$PATH`; build the binary yourself with `go build ./cmd/headhunter`. See
[Development](/development) for details.

### Using Go

```bash
go install github.com/xnet-admin-1/Headhunter/cmd/headhunter@latest
```

### Building from source

```bash
git clone https://github.com/xnet-admin-1/Headhunter.git
cd Headhunter
go build -o headhunter ./cmd/headhunter
```

## Quick Start

### Basic Usage

```bash
# Start interactive session
headhunter

# Run a one-off prompt
headhunter "List files in src/"

# Attach files as context
headhunter @main.go @test.go "Review these files"

# Continue the most recent session
headhunter --continue

# Model and thinking level selections are automatically persisted
# across sessions and restored on next launch

# Use specific model
headhunter --model anthropic/claude-sonnet-latest
```

### Non-Interactive Mode

```bash
# Get JSON output for scripting (stdout carries only the JSON envelope,
# so it pipes straight into jq; banners and warnings go to stderr)
headhunter "Explain main.go" --json | jq -r .response

# Quiet mode (final response only)
headhunter "Run tests" --quiet

# Ephemeral mode (no session file)
headhunter "Quick question" --no-session
```

### ACP Server Mode

Headhunter can run as an [ACP (Agent Client Protocol)](https://agentclientprotocol.com) agent server, enabling ACP-compatible clients (such as [OpenCode](https://github.com/sst/opencode)) to drive Headhunter as a remote coding agent over stdio.

```bash
# Start Headhunter as an ACP server (communicates via JSON-RPC 2.0 on stdin/stdout)
headhunter acp

# With debug logging to stderr
headhunter acp --debug
```

The ACP server exposes Headhunter's full capabilities — LLM execution, tool calls (shell, read, write, edit, grep, etc.), and session persistence — over the standard ACP protocol (protocol version 1). Sessions are persisted to Headhunter's normal JSONL session files.

Supported ACP features:

- **Sessions**: `session/new`, `session/load` (replays the history), `session/resume`, `session/list`, `session/close`, `session/cancel`
- **Working directory**: each session uses the `cwd` the client sends for project discovery (AGENTS.md, skills, `.kit.yml`) and as the base directory of the file and shell tools
- **MCP servers**: servers the client passes in `mcpServers` (stdio, HTTP, SSE) are connected in addition to the ones in Headhunter's config
- **Tool approval**: before edits, commands, subagents and MCP/extension tools, Headhunter asks the client with `session/request_permission` (allow/reject, once or always). Set the default with `headhunter acp --approval ask|auto_edit|auto`; clients change it per session with the `approval` config option
- **Editor integration**: when the client supports them, file reads and writes go through `fs/read_text_file` / `fs/write_text_file` (so the agent sees unsaved buffers), and shell commands run in client terminals (`terminal/*`) that are shown in the tool call
- **Config options**: `approval`, `model` and, for reasoning models, `thinking_level`, set with `session/set_config_option`
- **Slash commands**: prompt templates and skills are sent as `available_commands_update`; `/name args` expands them
- **Plans**: the model publishes its plan with the `update_plan` tool, shown as an ACP plan
- **Prompt content**: text, images, embedded resources, and resource links
- **Streaming**: agent messages, thoughts, and tool calls (with kind, file locations, and diffs for edits)
- **Auth errors**: missing provider credentials give the `auth_required` error with steps to fix them

To check an ACP build end to end, run `python3 scripts/acp_smoke_test.py` (set `MODEL=provider/model` to pick a model).

## Configuration

Headhunter looks for configuration in the following locations (in order of priority):

1. CLI flags
2. Environment variables (with `KIT_` prefix)
3. `./.kit.yml` / `./.kit.yaml` / `./.kit.json` (project-local)
4. `~/.kit.yml` / `~/.kit.yaml` / `~/.kit.json` (global)

### Basic Configuration

Create `~/.kit.yml`:

```yaml
model: anthropic/claude-sonnet-latest
max-tokens: 4096
temperature: 0.7
stream: true
thinking-level: off       # model-specific effort; off omits the field
no-core-tools: false      # set to true to disable all built-in core tools
exclude-core-tools:       # List of core tools to exclude, mutually exclusive to `include-core-tools`
#include-core-tools:
# - "shell"               # List of core tools to exclude, mutually exclusive to `exclude-core-tools`
shell: "bash"             # shell the shell tool runs commands through, plus
                          # its own leading arguments, e.g. "busybox ash". Set
                          # this on images that do not ship bash
shell-timeout: 120        # default per-call timeout in seconds
shell-max-timeout: 600    # ceiling a single call may request
codemode:                 # one "codemode" tool runs JavaScript that calls the
  enabled: true           # other tools; intermediate results stay out of context
  mcp-exposure: direct    # direct | codemode (scripts only) | deferred | model-only

# Images read by the read tool. PNG, JPEG, GIF and WebP files are sent to the
# model as viewable images instead of as text. An image above either limit is
# scaled down before it is sent, and re-encoded as JPEG when lossless encoding
# cannot meet the byte budget.
image-max-edge: 1568      # largest width or height in pixels
image-max-bytes: 3500000  # largest encoded payload in bytes
image-no-resize: false    # set to true to reject an oversized image instead

# Skills — all keys are optional
no-skills: false          # set to true to disable all skill loading
skill:                    # explicit skill files/dirs (disables auto-discovery)
  - /path/to/skill.md
skills-dir: ""            # scan this directory directly for skills (overrides auto-discovery)
skill-disable:            # hide skills from the model catalog by name (still usable via /<name>)
  - some-skill

# Named agents
no-agents: false          # set to true to disable named agent discovery (built-ins and definition files)
```

All of the above keys can also be set programmatically via the SDK
(`kit.Options.MaxTokens`, `Options.Temperature`, `Options.ThinkingLevel`, etc.)
without touching config files — see [SDK options](#with-options).

### Model-specific reasoning effort

`/thinking` lists the current model's catalog choices. Headhunter preserves their exact
spelling and order, including `xhigh`, `max`, and future named efforts. `off`
omits the effort field; it does not guarantee that the model stops reasoning.
`none` is a literal provider effort. Explicit CLI and SDK requests for an
unsupported effort on a graded model return an error before a request is sent.
The TUI reports any suggested substitution.

When metadata is missing, the default ladder is `off`, `none`, `minimal`, `low`,
`medium`, `high`, `xhigh`. Unknown-capability OpenAI-compatible endpoints also
accept explicit effort strings; the remote provider decides whether they are
valid. Token-budget adapters retain their finite controls. If an adapter cannot
encode a selected effort, Headhunter returns an error instead of silently dropping it.

Custom endpoints can declare an authoritative list in YAML or JSON:

```yaml
customModels:
  jundot-qwen:
    name: Qwen3.8-27B
    baseUrl: http://localhost:8000/v1
    apiModelName: Qwen3.8-27B-oQ6e-fp16-mtp
    reasoning: true
    reasoningLevels: [low, medium, xhigh]
    limit:
      context: 262144
      output: 32768
```

A non-empty list requires `reasoning: true`. Entries must be non-blank strings;
`off` is reserved. Duplicates are removed in order. Invalid models are excluded
with a configuration warning. An omitted or empty list leaves capabilities
unknown. `params.thinkingLevel` and `modelSettings` set defaults, not capability
lists. No model-name inference or effort translation is applied.

### Environment Variables

```bash
export ANTHROPIC_API_KEY="sk-..."
export OPENAI_API_KEY="sk-..."
export KIT_MODEL="openai/gpt-4o"
```

### MCP Server Configuration

Add external MCP servers to `.kit.yml`:

```yaml
mcpServers:
  filesystem:
    type: local
    command: ["npx", "-y", "@modelcontextprotocol/server-filesystem", "/path/to/allowed"]
    environment:
      LOG_LEVEL: "info"
    allowedTools: ["read_file", "write_file"]
  
  search:
    type: remote
    url: "https://mcp.example.com/search"

  pubmed:
    type: remote
    url: "https://pubmed.mcp.example.com"
    noOAuth: true  # skip OAuth for public servers that don't require auth

  builds:
    type: remote
    url: "https://builds.mcp.example.com"
    tasksMode: always  # async task execution — see MCP Tasks below
```

## CLI Reference

### Global Flags

```bash
# Model and provider
--model, -m              Model to use (provider/model format)
--provider-api-key       API key for the provider
--provider-url           Base URL for provider API
--provider-wire          Wire protocol for auto-routed providers (openai, openai-compat, anthropic, google)
--tls-skip-verify        Skip TLS certificate verification

# Session management
--session, -s            Open specific JSONL session file
--continue, -c           Resume most recent session for current directory
--resume, -r             Interactive session picker
--no-session             Ephemeral mode, no persistence

# Behavior (non-interactive: pass prompt as positional arg)
--quiet                  Suppress all output (non-interactive only)
--json                   Output response as JSON (non-interactive only)
--no-exit                Enter interactive mode after prompt completes
--max-steps              Maximum agent steps (0 for unlimited)
--stream                 Enable streaming output (default: true)
--compact                Enable compact output mode
--auto-compact           Auto-compact conversation near context limit (reactive compact-and-retry on provider context-overflow errors is always on)

# Extensions and tools
--extension, -e          Load additional extension file(s) (repeatable)
--no-extensions          Disable all extensions
--no-core-tools          Disable all built-in core tools (shell, read, write, edit, grep, find, ls, subagent)
--include-core-tools
--exclude-core-tools     Mutually exclusive lists of core tool names to include or not to include in agent
--codemode               Enable code mode (default on; --codemode=false disables): the model can run JavaScript that calls the other tools
--shell                  Shell the shell tool runs commands through, e.g. "/bin/dash" or "busybox ash" (default "bash")

--prompt-template        Load a specific prompt template by name
--no-prompt-templates    Disable prompt template loading

# Skills
--skill                  Load skill file or directory (repeatable)
--skills-dir             Scan this directory directly for skills (overrides auto-discovery)
--skill-disable          Hide a skill from the model catalog by name (repeatable); still usable via /<name>
--no-skills              Disable skill loading (auto-discovery and explicit)
--no-agents              Disable named agent discovery (built-ins and definition files)

# Generation parameters
--max-tokens             Maximum tokens in response (default: 8192, auto-raised up to 32768 for models with larger known output limits)
--temperature            Randomness 0.0-1.0 (default: 0.7)
--top-p                  Nucleus sampling 0.0-1.0 (default: 0.95)
--top-k                  Limit top K tokens (default: 40)
--stop-sequences         Custom stop sequences (comma-separated)
--frequency-penalty      Penalize frequent tokens 0.0-2.0 (default: 0.0)
--presence-penalty       Penalize present tokens 0.0-2.0 (default: 0.0)
--thinking-level         Model-specific reasoning effort (default: off; /thinking lists choices)

# System
--config                 Config file path (default: ~/.kit.yml)
--system-prompt          System prompt text or file path
--debug                  Enable debug logging
```

### Commands

```bash
# Authentication (OAuth for anthropic/openai/copilot, API key for everything else)
headhunter auth login [provider]          # Start OAuth flow (e.g., anthropic) or prompt for an API key
headhunter auth login groq --api-key ...  # Store an API key without prompting
headhunter auth login [provider] --set-default  # Set provider's default model as system default
headhunter auth logout [provider]         # Remove credentials for provider
headhunter auth status                    # Check authentication status
# Stored keys ($XDG_CONFIG_HOME/.kit/credentials.json) win over env vars; env vars stay the fallback.
# No key yet? Headhunter still starts — use /connect inside the TUI to add one.

# GitHub Copilot login (experimental; requires active Copilot subscription)
headhunter auth login copilot
headhunter --model copilot/gpt-5.5 "Hello"

# Model database
headhunter models [provider]        # List available models (optionally filter by provider)
headhunter models --all             # Show all providers (not just LLM-compatible)
headhunter update-models [source]   # Update model database (from models.dev, URL, file, or 'embedded')

# Extension management
headhunter extensions list          # List discovered extensions
headhunter extensions validate      # Validate extension files
headhunter extensions init          # Generate example extension template
headhunter install <git-url>        # Install extensions from git repositories
headhunter install -l <git-url>     # Install to project-local .kit/git/ directory
headhunter install -u <git-url>     # Update an already-installed package
headhunter install --uninstall <pkg> # Remove an installed package

# Skills
headhunter skill list               # List discovered skills with scope and spec warnings
headhunter skill validate <path>    # Validate a skill against the agentskills.io spec
headhunter skill install            # Install the Headhunter skills (kit-extensions, kit-sdk) via skills.sh

# GitHub integration
headhunter github install           # Scaffold .github/workflows/kit.yml (run Headhunter on '/kit' comments)
headhunter github install --model anthropic/claude-sonnet-4-5-20250929
headhunter github install --force   # Overwrite an existing workflow file
headhunter github install --no-secret # Skip the offer to set the provider secret via the gh CLI

# ACP server
headhunter acp                      # Start as ACP agent (stdio JSON-RPC)
headhunter acp --debug              # With debug logging to stderr

# Detachable sessions (survive closing the terminal; Ctrl+] is the prefix)
headhunter daemon service install   # Start the daemon once; after this, plain `headhunter` is detachable
cd ~/project && headhunter          # A detachable session, right here — Ctrl+] d to leave it working
headhunter session attach           # Pick a live session (local or paired host), or start one
headhunter session attach 3         # Attach straight to session 3
headhunter session attach --new     # Skip the picker, start a new session
headhunter session ls               # List live sessions
headhunter session ls --all         # Include sessions on every paired host
headhunter session rename 3 work    # Name session 3
headhunter session kill 3           # Confirm, then stop session 3
headhunter --no-daemon              # Run in this terminal, daemon or no daemon

# Remote sessions (host)
headhunter daemon                   # Host sessions for local and paired clients
headhunter daemon pair              # One-time pairing: show code, confirm here
headhunter daemon status            # Endpoint, paired clients, active sessions
headhunter daemon service install   # Install + start the systemd user service

# Remote sessions (client)
headhunter remote --pair A1B2C3D4   # Pair with a host and save it under a name
headhunter session attach --host homelab # Attach to the paired host
headhunter session attach --all     # Pick across paired hosts without starting a local daemon
headhunter session rename --host homelab 3 work
headhunter session kill --host homelab --yes 3
```

While attached, use `Ctrl+] ,` to edit the session name, or `Ctrl+] k` to
kill the session with confirmation. `/name <name>` updates the live-session
picker name. `/kill` asks you to enter `/kill` again to confirm; `/kill cancel`
cancels the request. Killing a session retains its saved conversation history.
The old `headhunter attach`, `headhunter ls`, and `headhunter remote --host` attachment commands
remain as deprecated aliases.

## Themes

Headhunter ships with 22 built-in color themes that control all UI elements. Switch at runtime:

```
/theme dracula
/theme catppuccin
/theme tokyonight
```

Theme selections are automatically saved and restored on next launch (stored in `~/.config/kit/preferences.yml`). This persistence also applies to **model** and **thinking level** selections — all are saved together and restored on startup.

### Custom themes

Drop a `.yml` file in `~/.config/kit/themes/` (user) or `.kit/themes/` (project):

```yaml
# ~/.config/kit/themes/my-theme.yml
primary:
  light: "#8839ef"
  dark: "#cba6f7"
success:
  light: "#40a02b"
  dark: "#a6e3a1"
```

Built-in themes: `kitt`, `catppuccin`, `dracula`, `tokyonight`, `nord`, `gruvbox`, `monokai`, `solarized`, `github`, `one-dark`, `rose-pine`, `ayu`, `material`, `everforest`, `kanagawa`, `amoled`, `synthwave`, `vesper`, `flexoki`, `matrix`, `vercel`, `zenburn`

## Extension System

Extensions are Go source files that run via Yaegi interpreter. They can add custom tools, slash commands, widgets, keyboard shortcuts, themes, and intercept lifecycle events.

### Minimal Extension

```go
//go:build ignore

package main

import "kit/ext"

func Init(api ext.API) {
    api.OnSessionStart(func(_ ext.SessionStartEvent, ctx ext.Context) {
        ctx.SetFooter(ext.HeaderFooterConfig{
            Content: ext.WidgetContent{Text: "Custom Footer"},
        })
    })
}
```

**Usage:**

```bash
headhunter -e examples/extensions/minimal.go
```

### Extension Capabilities

**Lifecycle Events**: OnSessionStart, OnSessionShutdown, OnBeforeAgentStart, OnAgentStart, OnAgentEnd, OnLLMUsage, OnToolCall, OnToolCallInputStart, OnToolCallInputDelta, OnToolCallInputEnd, OnToolExecutionStart, OnToolOutput, OnToolExecutionEnd, OnToolResult, OnInput, OnMessageStart, OnMessageUpdate, OnMessageEnd, OnModelChange, OnThinkingLevelChange, OnTerminalResize, OnTurnStateChange, OnMCPToolsReady, OnContextPrepare, OnBeforeFork, OnBeforeSessionSwitch, OnBeforeCompact, OnCustomEvent, OnSubagentStart, OnSubagentChunk, OnSubagentEnd

`OnAgentEnd` carries per-turn aggregates (`ToolCallCount`, `ToolNames`, `LLMCallCount`, `InputTokensDelta`, `OutputTokensDelta`, `CostDelta`, `DurationMs`) so observers don't need to maintain parallel bookkeeping. `OnLLMUsage` fires after each LLM provider call with token + cost deltas attributed to that specific call/model — use it for accurate budget enforcement *between* calls instead of waiting for the turn to finish.

`OnTurnStateChange` is a superset of `OnAgentStart`/`OnAgentEnd`: it also covers work that never reaches the agent loop (shell commands) and fires on every path back to idle, including cancellation and error — use it for spinners and turn timers, and `OnAgentStart`/`OnAgentEnd` when you specifically care about agent turns. `OnTerminalResize` and `OnTurnStateChange` fire in the interactive TUI only.

**Custom Components**:
- **Tools**: Add new tools the LLM can invoke
- **Commands**: Register slash commands (e.g., `/mycommand`)
- **Options**: Register configurable extension options
- **Session State**: Last-write-wins key-value store via `ctx.SetState` / `GetState` / `DeleteState` / `ListState`, persisted to a per-session sidecar file outside the conversation tree
- **Widgets**: Persistent status displays above/below input
- **Headers/Footers**: Persistent content above/below the conversation (rendered at full width without truncation — size with `ctx.GetTerminalSize()`)
- **Status Bar**: Custom status bar entries
- **Shortcuts**: Global keyboard shortcuts
- **Overlays**: Modal dialogs with markdown content
- **Tool Renderers**: Customize how tool calls display
- **Message Renderers**: Custom rendering for assistant messages
- **Editor Interceptors**: Handle key events and wrap rendering
- **Interactive Prompts**: Select, confirm, input, and multi-select dialogs
- **Subagents**: Spawn in-process child Headhunter instances
- **LLM Completion**: Direct model calls via `Complete()`
- **Themes**: Register, switch, and read color themes via `RegisterTheme`, `SetTheme`, `ListThemes`, `GetTheme` (read the active palette to paint widgets in the user's theme)
- **Custom Events**: Inter-extension communication via `EmitCustomEvent`

**Bridged SDK APIs** (NEW): Extensions can now access internal SDK capabilities:
- **Tree Navigation**: Navigate conversation history (`GetTreeNode`, `GetCurrentBranch`, `NavigateTo`), summarize branches (`SummarizeBranch`), and implement fresh context loops (`CollapseBranch`)
- **Skill Loading**: Dynamically load and inject skills at runtime (`LoadSkill`, `DiscoverSkills`, `InjectSkillAsContext`)
- **Template Parsing**: Parse and render templates with `{{variables}}` (`ParseTemplate`, `RenderTemplate`), parse CLI-style arguments (`ParseArguments`, `SimpleParseArguments`), and evaluate model conditionals (`EvaluateModelConditional`, `RenderWithModelConditionals`)
- **Model Resolution**: Resolve model fallback chains (`ResolveModelChain`), query model capabilities (`GetModelCapabilities`, `CheckModelAvailable`), and extract provider/model ID (`GetCurrentProvider`, `GetCurrentModelID`)
- **Model Pricing**: Per-million-token costs on `ModelCapabilities.Pricing` and `ModelInfoEntry.Pricing` — check `Pricing.Known` before rendering a cost, since local and OpenAI-compatible models report no pricing
- **Usage & Environment**: Session tokens and cost via `GetSessionUsage` (including `IsOAuth`, since subscriptions report `$0` cost), terminal dimensions via `GetTerminalSize`, reasoning effort via `GetThinkingLevel`, and the composed system prompt with its configured source via `GetSystemPrompt`

### Extension Examples

See the `examples/extensions/` directory:

- [`minimal.go`](examples/extensions/minimal.go) - Clean UI with custom footer
- [`auto-commit.go`](examples/extensions/auto-commit.go) - Auto-commit on shutdown
- [`bookmark.go`](examples/extensions/bookmark.go) - Bookmark conversations
- [`branded-output.go`](examples/extensions/branded-output.go) - Branded output rendering
- [`bridge-demo.go`](examples/extensions/bridge_demo.go) - Bridged SDK API demo (tree navigation, skills, templates, model resolution)
- [`compact-notify.go`](examples/extensions/compact-notify.go) - Notification on compaction
- [`confirm-destructive.go`](examples/extensions/confirm-destructive.go) - Confirm destructive operations
- [`context-inject.go`](examples/extensions/context-inject.go) - Inject context into conversations
- [`conversation-manager.go`](examples/extensions/conversation-manager.go) - **NEW** Tree navigation, branch summarization, and fresh context loops
- [`custom-editor-demo.go`](examples/extensions/custom-editor-demo.go) - Vim-like modal editor
- [`dev-reload.go`](examples/extensions/dev-reload.go) - Development live-reload
- [`header-footer-demo.go`](examples/extensions/header-footer-demo.go) - Custom headers and footers
- [`inline-bash.go`](examples/extensions/inline-bash.go) - Inline bash execution
- [`interactive-shell.go`](examples/extensions/interactive-shell.go) - Interactive shell integration
- [`kit-kit.go`](examples/extensions/kit-kit.go) - Headhunter-in-Headhunter (sub-agent spawning)
- [`lsp-diagnostics.go`](examples/extensions/lsp-diagnostics.go) - LSP diagnostic integration
- [`notify.go`](examples/extensions/notify.go) - Desktop notifications
- [`overlay-demo.go`](examples/extensions/overlay-demo.go) - Modal dialogs
- [`permission-gate.go`](examples/extensions/permission-gate.go) - Permission gating for tools
- [`pirate.go`](examples/extensions/pirate.go) - Pirate-themed personality
- [`plan-mode.go`](examples/extensions/plan-mode.go) - Read-only planning mode
- [`project-rules.go`](examples/extensions/project-rules.go) - Project-specific rules
- [`prompt-demo.go`](examples/extensions/prompt-demo.go) - Interactive prompts (select/confirm/input)
- [`prompt-templates.go`](examples/extensions/prompt-templates.go) - **NEW** Frontmatter-driven templates with model switching and skill injection
- [`protected-paths.go`](examples/extensions/protected-paths.go) - Path protection for sensitive files
- [`status-footer.go`](examples/extensions/status-footer.go) - **NEW** Configurable status bar footer (model, context, cache, cost, clock, turn timing)
- [`subagent-monitor.go`](examples/extensions/subagent-monitor.go) - **NEW** Live subagent dashboard: per-frame `Render` strip, `RefreshHz` animation, theme-aware colors, and a real `ShowOverlay` modal for the full transcript (`ctrl+alt+s`)
- [`subagent-widget.go`](examples/extensions/subagent-widget.go) - Multi-agent orchestration with status widget
- [`subagent-test.go`](examples/extensions/subagent-test.go) - Subagent testing utilities
- [`summarize.go`](examples/extensions/summarize.go) - Conversation summarization
- [`tool-logger.go`](examples/extensions/tool-logger.go) - Log all tool calls
- [`neon-theme.go`](examples/extensions/neon-theme.go) - Custom theme registration and switching
- [`tool-renderer-demo.go`](examples/extensions/tool-renderer-demo.go) - Custom tool call rendering
- [`usage-budget.go`](examples/extensions/usage-budget.go) - Per-call usage callback (`OnLLMUsage`), session state, and enriched `OnAgentEnd` per-turn report
- [`widget-status.go`](examples/extensions/widget-status.go) - Persistent status widgets

Also see [`.kit/extensions/go-edit-lint.go`](.kit/extensions/go-edit-lint.go) (in this repo) for a project-local extension example that runs gopls and golangci-lint on Go file edits.

### Loading Extensions

**Auto-discovery** (loads automatically):
- `/usr/share/kit/extensions/*.go` (system-wide single files)
- `/usr/share/kit/extensions/*/main.go` (system-wide subdirectory extensions)
- `~/.config/kit/extensions/*.go` (user single files)
- `~/.config/kit/extensions/*/main.go` (user subdirectory extensions)
- `.kit/extensions/*.go` (project-local single files)
- `.kit/extensions/*/main.go` (project-local subdirectory extensions)
- `~/.local/share/kit/git/` (global git-installed packages)
- `.kit/git/` (project-local git-installed packages)

The system-wide directory serves packaged installs (rpm, deb, ...). Set
`KIT_SYSTEM_EXTENSIONS_DIR` to use other directories (separated by the platform
list separator: `:` on Unix, `;` on Windows), or set it empty to disable
system-wide discovery. Packagers can change the
compiled-in default with
`-ldflags "-X github.com/xnet-admin-1/Headhunter/internal/extensions.SystemExtensionsDir=/opt/kit/extensions"`.

**Explicit loading**:
```bash
headhunter -e path/to/extension.go
headhunter -e ext1.go -e ext2.go  # Multiple extensions
```

**Disable auto-load**:
```bash
headhunter --no-extensions
```

### Testing Extensions

Headhunter provides a testing package to help you write unit tests for your extensions.
`pkg/extensions/test` holds the harness and the assertions; `pkg/extensions`
holds the event, result, and config types the harness speaks in. Both are
public, so these tests build from your own repository:

```go
package main

import (
    "testing"

    "github.com/xnet-admin-1/Headhunter/pkg/extensions"
    "github.com/xnet-admin-1/Headhunter/pkg/extensions/test"
)

func TestMyExtension(t *testing.T) {
    harness := test.New(t)
    harness.LoadFile("my-ext.go")
    
    // Emit events and verify behavior
    _, err := harness.Emit(extensions.SessionStartEvent{SessionID: "test"})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    // Verify the extension printed something
    test.AssertPrinted(t, harness, "session started")
}
```

**Available assertions:**
- `AssertBlocked()`, `AssertNotBlocked()` — Verify tool blocking
- `AssertWidgetSet()`, `AssertWidgetText()` — Verify widget content
- `AssertPrinted()`, `AssertPrintedContains()` — Verify output
- `AssertToolRegistered()`, `AssertCommandRegistered()` — Verify registration

See [`examples/extensions/tool-logger_test.go`](examples/extensions/tool-logger_test.go) for a complete example with 14 test cases covering tool calls, input handling, and session lifecycle.

### Prompt Templates

Create reusable prompt templates with shell-style argument substitution. Templates are loaded from `~/.kit/prompts/*.md` and `.kit/prompts/*.md`.

**Example template** (`~/.kit/prompts/review.md`):
```markdown
---
description: Review code for issues
---
Review the following code for bugs and security issues.
Focus on $1 specifically.
```

**Usage:**
```
/review error handling
```

**Argument placeholders:**
- `$1`, `$2`, etc. — Individual arguments
- `$@` or `$ARGUMENTS` — All arguments (zero or more)
- `$+` — All arguments (one or more required; error if none given)
- `${@:2}` — Arguments from position 2 onwards
- `${@:1:3}` — 3 arguments starting at position 1

Placeholders inside fenced code blocks (```) and inline code spans are ignored.

Disable templates with `--no-prompt-templates` or load a specific template with `--prompt-template <name>`.

### Named Agents

Define reusable subagent presets as markdown files. Named agents are advertised to the LLM in the `subagent` tool description, so the main agent can delegate tasks to the right specialist by name — with a preset system prompt, model, tool allowlist, temperature, and timeout.

**Discovery locations** (highest to lowest precedence):

1. `<project>/.agents/agents/*.md` — project-local (cross-client convention)
2. `<project>/.kit/agents/*.md` — project-local (Headhunter-specific)
3. `~/.config/kit/agents/*.md` — user-level (`$XDG_CONFIG_HOME` aware)
4. Built-in agents: `general` (full tool access) and `explore` (read-only: `read`, `grep`, `find`, `ls`)

The filename (minus `.md`) is the agent name. Higher-precedence definitions override lower ones, so a project can replace — or disable — a built-in or user-level agent.

**Example** (`.agents/agents/code-reviewer.md`):

```markdown
---
description: Reviews code for quality and best practices   # required
model: anthropic/claude-sonnet-4                           # optional model override
tools: [read, grep, find, ls]                              # optional tool allowlist
temperature: 0.1                                           # optional
timeout: 300                                               # optional, seconds
hidden: false                                              # optional: resolvable but not advertised
disabled: false                                            # optional: remove this agent (and anything it shadows)
---
You are in code review mode. Focus on correctness, security, and
maintainability. Report findings with file paths and line references.
```

The LLM invokes it through the `subagent` tool's `agent` parameter; explicit `model` / `system_prompt` / `timeout_seconds` arguments override the agent's presets. An agent without a `tools:` list gets the default subagent tool set (everything except `subagent`, preventing recursion); with a `tools:` allowlist the subagent is restricted to exactly those tools.

From the Go SDK:

```go
agents := k.GetAgents() // discovered definitions (built-ins included)

result, err := k.Subagent(ctx, kit.SubagentConfig{
    Prompt: "Map out the session persistence flow",
    Agent:  "explore", // preset prompt + read-only tools
})
```

Disable discovery entirely with `--no-agents`, the `no-agents` config key (`.kit.yml`), `KIT_NO_AGENTS=true`, or `Options.NoAgents` in the SDK.

### Inspecting Subagent Runs

Inspect retained child output and tool activity in the TUI with `/subagents`. Press Ctrl+Alt+A to open a subagent run picker, then Enter to inspect the selected run. The subagent status bar shows this shortcut when space permits. Use Left/Right to switch sibling runs, Up/Down or Page Up/Page Down to scroll, Esc to return to the parent, and Ctrl+K to stop an active run. The conversation view is read-only and refreshes once per second.

The status bar shows active runs with blinking lights and finished results for 10 seconds. Older retained runs remain available in the picker. The view uses in-memory history, not saved session files.

**Command migration:** `/subagent-sessions` and `/agents` have been removed; use `/subagents` instead. Extensions can no longer override `/subagents`. The `subagent-monitor.go` example now opens its overlay with `/subagent-monitor` (or Ctrl+Alt+S).

From the SDK, use `k.SubagentRuns()` for active and retained finished runs, or `k.GetSubagentRun(runID)` for one snapshot. `k.RunningSubagents()` continues to list only active runs. History retains up to 100 finished runs plus all active runs and up to 500 event records per run.

See [Inspecting Subagent Runs](#inspecting-subagent-runs) for controls, history limits, and SDK examples.

## GitHub Integration

Headhunter can run as an automated collaborator/reviewer inside GitHub Actions. The
`headhunter github install` command scaffolds a workflow that triggers when someone
comments `/kit ...` on an issue or pull request review, runs the agent
non-interactively in the runner, and lets it respond.

```bash
headhunter github install
```

This writes `.github/workflows/kit.yml`. By default the command prompts for the
model (pre-filled with a sensible default); pass `--model` to skip the prompt.
If the [`gh` CLI](https://cli.github.com/) is detected on your `PATH` and the
provider API key is present in your environment, you'll be offered the option to
store it as a repository secret automatically.

The generated workflow:

- Triggers only on `issue_comment` and `pull_request_review_comment` (`types: [created]`).
- Runs only when the comment begins with the `/kit` command token.
- Restricts triggers to repository owners, members, and collaborators (via `author_association`).
- Uses least-privilege `permissions` and `persist-credentials: false`.
- Authenticates git/PR operations with the built-in `secrets.GITHUB_TOKEN` and
  the provider via a repository secret (e.g. `ANTHROPIC_API_KEY`).

After committing the workflow and setting the provider secret, comment
`/kit <your request>` on any issue or pull request to trigger Headhunter.

The generated workflow uses the bundled [`xnet-admin-1/Headhunter`](action.yml) composite
action, which installs the Headhunter binary and runs `headhunter github run`. That command
reads the triggering event, enforces permissions, reacts with an emoji, runs the
agent against the issue thread or pull request, posts the response as a comment,
and — if the agent changed files — pushes a `kit-agent[bot]` branch and opens a
pull request.

| Flag | Description |
| --- | --- |
| `--model` | Provider/model to write into the workflow |
| `--force` | Overwrite an existing workflow file |
| `--no-secret` | Skip the offer to set the provider secret via the `gh` CLI |

## Session Management

Headhunter uses a tree-based session model that supports branching and forking conversations.

### Session Locations

- Default: `~/.kit/sessions/<cwd-path>/<timestamp>_<id>.jsonl`
- Path separators in the working directory are replaced with `--` (e.g., `/home/user/project` becomes `home--user--project`)
- Each line is a session entry (messages, tool calls, extension data)
- Supports branching from any message to explore alternate paths

### Session Commands

```bash
# Resume most recent session for current directory
headhunter --continue
headhunter -c

# Interactive session picker
headhunter --resume
headhunter -r

# Open specific session file
headhunter --session path/to/session.jsonl
headhunter -s path/to/session.jsonl

# Ephemeral mode (no file persistence)
headhunter --no-session
```

### Interactive Session Commands

During an interactive session, use these slash commands:

| Command | Description |
|---------|-------------|
| `/name [name]` | Set or display the session's display name |
| `/session` | Show session info (path, ID, message count) |
| `/resume` | Open the session picker to switch sessions |
| `/export [path]` | Export session as JSONL (auto-generates path if omitted) |
| `/import <path>` | Import and switch to a session from a JSONL file |
| `/share` | Upload session to GitHub Gist and get a shareable viewer URL |
| `/tree` | Navigate the session tree |
| `/fork` | Fork to new session from an earlier message |
| `/new` | Start a fresh session |

### Keyboard Shortcuts

| Shortcut | Description |
|----------|-------------|
| `Ctrl+V` | Paste an image from the clipboard — shows an inline low-res thumbnail preview (tmux/zellij-safe); over `headhunter remote` it pastes from the **local** machine's clipboard |
| `Ctrl+U` | Clear all pending image attachments |
| `Ctrl+X e` | Open `$VISUAL`/`$EDITOR` to compose or edit your prompt |
| `Ctrl+X s` | Steer — inject a system-level instruction mid-turn |
| `Ctrl+]` | Session prefix in any daemon-hosted session (`headhunter` with a daemon running, `headhunter attach`, `headhunter remote`): `d` detach · `s` switch · `c` new · `n`/`p` cycle · `w` across hosts |
| `ESC ESC` | Cancel the current operation (tool call or streaming) |
| `↑` / `↓` | Navigate prompt history |
| Mouse drag | Select text in the conversation. Headhunter copies it on release and keeps it highlighted until the next key or click, or until the text changes. On Linux and the BSDs it also sets the primary selection for middle-click paste in other windows |

## Go SDK

Embed Headhunter in your Go applications:

```go
package main

import (
    "context"
    "log"
    
    kit "github.com/xnet-admin-1/Headhunter/pkg/kit"
)

func main() {
    ctx := context.Background()
    
    // Create Headhunter instance with default configuration
    host, err := kit.New(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer func() { _ = host.Close() }()
    
    // Send a prompt
    response, err := host.Prompt(ctx, "What is 2+2?")
    if err != nil {
        log.Fatal(err)
    }
    
    println(response)
}
```


### HTTP request identity

Headhunter-controlled HTTP requests use `User-Agent: KIT` by default. Set a per-instance
identity with either API:

```go
host, err := kit.New(ctx, &kit.Options{UserAgent: "MyApp/1.0"})
// Or use functional options:
host, err = kit.NewAgent(ctx, kit.WithUserAgent("MyApp/1.0"))
```

Empty means `KIT`, not an opt-out. Valid values are preserved without trimming,
including spaces. HTTP control characters (including CR, LF, NUL and DEL) are
rejected before construction or login does network work. This setting has no
CLI flag, environment variable, or config-file key.

Each instance keeps its own identity. It passes to model changes, temporary
completions and subagents, including isolated/bare children and provider changes.
It covers built-in model requests (including streams), Ollama setup, remote MCP
and their OAuth requests, and provider token refresh. The identity replaces
conflicting user-agent headers from model defaults, catalog headers, per-call
headers and MCP configuration, case-insensitively. Other headers stay unchanged.
Custom provider factories receive `ProviderConfig.UserAgent` and must enforce it
on their own requests; Headhunter does not intercept arbitrary factory traffic.

Two inference paths keep their fixed compatibility identities, even with an
SDK override. Their token endpoints still use the configured Headhunter identity.
These exceptions follow existing source requirements; they were not live-tested.

| Inference path | Fixed user agent |
| --- | --- |
| Codex | `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36` |
| Copilot | `GitHubCopilotChat/0.31.0` |

Standalone `kit.LoginProvider` has a separate `kit.LoginOptions.UserAgent` field
(default `KIT`), independent of any Headhunter instance or supplied credential manager.
Standalone CLI OAuth and model-catalog downloads use `KIT`.

This policy does not change global HTTP defaults. It does not cover subprocesses
(shell, git/gh/curl, installers, npm or MCP server processes), extension/application
code, browser navigation, or remote services acting for Headhunter. Third-party cloud
credential discovery (metadata services or default credential chains) inherits
the identity only when it uses the HTTP client Headhunter supplies; model endpoint
requests use the configured identity. A custom transport can rewrite headers
after Headhunter's boundary and is responsible for its own policy.

### With Options

```go
host, err := kit.New(ctx, &kit.Options{
    Model:        "ollama/llama3",
    SystemPrompt: "You are a helpful bot",
    ConfigFile:   "/path/to/config.yml",
    MaxSteps:     10,
    Streaming:    ptr(true), // *bool: nil = unset (default true), &false = off
    Quiet:        true,

    // Generation parameters (override env/config/per-model defaults)
    MaxTokens:        16384,             // 0 = auto-resolve (env → config → per-model → 8192 floor)
    ThinkingLevel:    "medium",          // model-specific; see SupportedThinkingLevels
    Temperature:      ptr(float32(0.2)), // pointer so 0.0 != unset; nil = provider default
    TopP:             nil,                // nil = leave provider/per-model default
    TopK:             nil,
    FrequencyPenalty: nil,
    PresencePenalty:  nil,

    // Provider configuration (override env/config without reaching into viper)
    ProviderAPIKey: "sk-...",                      // "" = use config / provider env var
    ProviderURL:    "https://proxy.internal/v1",   // "" = provider default
    TLSSkipVerify:  false,                         // only takes effect when true

    // Session options
    SessionPath:  "./session.jsonl",  // Open specific session
    Continue:     true,                // Resume most recent session
    NoSession:    true,                // Ephemeral mode

    // Tool options
    Tools:            []kit.Tool{...},     // Replace default tool set entirely
    ExtraTools:       []kit.Tool{...},     // Add tools alongside defaults
    DisableCoreTools: true,                // Disable all built-in core tools; also controllable via
                                           // --no-core-tools flag, KIT_NO_CORE_TOOLS env var,
                                           // or no-core-tools: true in .kit.yml
    CoreToolList      []string,            // List of core tool names to  register. If empty (default), include all.
                                           // DisableCoreTools has precedence.

    // Configuration
    SkipConfig:   true,                   // Skip .kit.yml files (viper defaults + env vars still apply)

    // Compaction — proactive check before turns near the context limit.
    // Reactive compaction (compact + replay once on provider context-overflow
    // errors) is always on, independent of this setting.
    AutoCompact:  true,

    Debug:        true,                // Debug logging
})
```

**Generation & provider fields** (added in v0.55+) let SDK consumers configure
Headhunter entirely in-code without `viper.Set()` workarounds or shipping a `.kit.yml`.
Precedence is `Options` > `KIT_*` env vars > `.kit.yml` > per-model defaults
(`modelSettings` / `customModels`) > provider-level defaults. Sampling params
are pointer types so explicit `0.0` is distinguishable from "leave alone"; a
non-zero `MaxTokens` suppresses automatic right-sizing the same way `--max-tokens`
does on the CLI.

### Functional options (`NewAgent`)

For simple programmatic setups, `kit.NewAgent` offers an ergonomic
functional-options front door over `kit.New`. Streaming is **enabled by
default**; pass `kit.WithStreaming(false)` to opt out.

```go
host, err := kit.NewAgent(ctx,
    kit.WithModel("anthropic/claude-sonnet-4-5-20250929"),
    kit.WithSystemPrompt("You are a helpful assistant."),
    kit.WithMaxTokens(8192),
    kit.WithThinkingLevel("medium"),
    kit.Ephemeral(), // in-memory session, no persistence
)
```

Available options: `WithModel`, `WithSystemPrompt`, `WithStreaming`,
`WithMaxTokens`, `WithThinkingLevel`, `WithTools`, `WithExtraTools`,
`WithProviderAPIKey`, `WithProviderURL`, `WithProvider`, `WithConfigFile`,
`WithDebug`, `WithDebugLogger`, and `Ephemeral`. For advanced configuration not covered by
the helpers (custom MCP config, in-process MCP servers, session backends, MCP
task tuning) construct an `Options` value explicitly and call `kit.New`.

An embedded Headhunter discovers `.kit.yml`, `AGENTS.md`, skills, extensions and
named agents on the host, as the CLI does. It also writes session files and
enables the core file and shell tools. `kit.NewIsolatedAgent` turns all of
this off. The opt-in options (`WithConfig`, `WithContextFiles`, `WithSkills`,
`WithExtensions`, `WithAgents`, `WithSessions`, `WithCoreTools`) turn each
feature back on:

```go
host, err := kit.NewIsolatedAgent(ctx,
    kit.WithModel("anthropic/claude-sonnet-4-5-20250929"),
    kit.WithCoreTools("read", "grep"), // only these core tools
)
```

`kit.Isolated()` is the same preset as an `Option` for `kit.NewAgent`. See
[Isolated agents](pkg/kit/README.md#isolated-agents-isolated-newisolatedagent).

### Per-instance config isolation

Each `kit.New` / `kit.NewAgent` call owns an **isolated configuration store**,
so constructing multiple Headhunter instances in the same process is safe: setting the
model, thinking level, or generation parameters on one never affects another,
and runtime mutators (`SetModel`, `SetThinkingLevel`) only touch the owning
instance. This makes subagent spawning and multi-Headhunter embedding race-free with
no external synchronization required.

### MCP OAuth (remote MCP servers)

When a remote MCP server returns 401, Headhunter runs the full OAuth flow (dynamic
client registration → PKCE → token exchange → persistence) but delegates the
user-facing step — showing the authorization URL and receiving the callback —
to an `MCPAuthHandler` that you pass explicitly via `Options.MCPAuthHandler`.
If nil, OAuth is disabled and the authorization-required error surfaces to the
caller; the SDK never auto-opens a browser or binds a localhost port.

```go
// CLI/TUI apps: opens the system browser + prints status to stderr.
authHandler, _ := kit.NewCLIMCPAuthHandler()
defer authHandler.Close()

host, _ := kit.New(ctx, &kit.Options{
    MCPAuthHandler: authHandler,
})

// Custom UX: reuse the SDK's port + callback server, supply your own
// presentation via OnAuthURL (TUI modal, QR code, web redirect, etc.).
//   h, _ := kit.NewDefaultMCPAuthHandler()
//   h.OnAuthURL = func(server, authURL string) { myUI.Show(server, authURL) }
//
// Full control (web apps, daemons): implement kit.MCPAuthHandler yourself —
// no localhost binding, no side effects.
```

Tokens are persisted to `$XDG_CONFIG_HOME/.kit/mcp_tokens.json` by default; swap
in a custom `MCPTokenStoreFactory` for encrypted, DB-backed, or in-memory
storage. See the [SDK options docs](/sdk/options#mcp-oauth-authorization) for
the full matrix.

### MCP Tasks (long-running tools)

Headhunter advertises [MCP task support](https://modelcontextprotocol.io/specification/2025-11-25/basic/utilities/tasks)
during `initialize`, so cooperating MCP servers can respond to `tools/call`
with a `taskId` instead of blocking the connection. Headhunter then polls
`tasks/get` / `tasks/result` until the task reaches a terminal state, and
best-effort `tasks/cancel`s on context cancellation.

Defaults are safe — a server that doesn't advertise task capability runs
synchronously, exactly as before. Opt in per server via `tasksMode` in
`.kit.yml` (`auto` | `never` | `always`) or programmatically through the SDK:

```go
host, _ := kit.New(ctx, &kit.Options{
    MCPTaskMode: map[string]kit.MCPTaskMode{
        "build-server": kit.MCPTaskModeAlways,
    },
    MCPTaskTimeout:  15 * time.Minute,
    MCPTaskProgress: func(p kit.MCPTaskProgress) {
        log.Printf("%s: %s", p.TaskID, p.Status)
    },
})

tasks, _ := host.ListMCPTasks(ctx, "build-server")
_, _    = host.CancelMCPTask(ctx, "build-server", tasks[0].TaskID)
```

See the [configuration docs](/configuration#mcp-tasks-long-running-tools) and
[SDK options → MCP Tasks](/sdk/options#mcp-tasks) for the full surface.

### Custom Tools

Create custom tools with automatic schema generation — no external dependencies needed:

```go
type SearchInput struct {
    Query string `json:"query" description:"Search query"`
}

searchTool := kit.NewTool("search", "Search the codebase",
    func(ctx context.Context, input SearchInput) (kit.ToolOutput, error) {
        return kit.TextResult("Found: ..."), nil
    },
)

host, _ := kit.New(ctx, &kit.Options{
    ExtraTools: []kit.Tool{searchTool}, // adds alongside built-in tools
})
```

Use `kit.NewParallelTool` for tools safe to run concurrently. Binary data (images, audio, etc.) in `ToolOutput.Data` is automatically forwarded to the LLM when `MediaType` is set. See the [SDK docs](/sdk/overview) for full details on struct tags, `ToolOutput` fields, and `ToolCallIDFromContext`.

#### Return Helpers

| Helper | Description |
| --- | --- |
| `kit.TextResult(content)` | Successful text result |
| `kit.ErrorResult(content)` | Error result (LLM sees it as a tool error) |
| `kit.ImageResult(content, data, mediaType)` | Image result with binary data (e.g. `"image/png"`) |
| `kit.MediaResult(content, data, mediaType)` | Non-image media result (e.g. `"audio/mpeg"`) |

#### ToolOutput Fields

```go
kit.ToolOutput{
    Content:   "result text",     // text returned to the LLM
    IsError:   false,             // true = LLM sees this as an error
    Data:      pngBytes,          // optional binary data (images, audio)
    MediaType: "image/png",       // MIME type for binary Data
    Metadata:  map[string]any{},  // opaque metadata for hooks/UI (not sent to LLM)
}
```

### With Callbacks

```go
unsub := host.OnToolCall(func(e kit.ToolCallEvent) {
    println("Calling tool:", e.ToolName)
})
defer unsub()

unsub2 := host.OnToolResult(func(e kit.ToolResultEvent) {
    if e.IsError {
        println("Tool failed:", e.ToolName)
    }
})
defer unsub2()

unsub3 := host.OnMessageUpdate(func(e kit.MessageUpdateEvent) {
    print(e.Chunk)
})
defer unsub3()

response, err := host.Prompt(
    ctx,
    "List files in current directory",
)
```

### Session Management

```go
// Multi-turn conversations retain context automatically
host.Prompt(ctx, "My name is Alice")
response, _ := host.Prompt(ctx, "What's my name?")

// Sessions are persisted automatically to JSONL files.
// Access session info:
path := host.GetSessionPath()
id := host.GetSessionID()

// Clear conversation history
host.ClearSession()
```

Session persistence is configured via `Options`:

```go
host, _ := kit.New(ctx, &kit.Options{
    SessionPath: "./my-session.jsonl",  // Open specific session
    Continue:    true,                   // Resume most recent session
    NoSession:   true,                   // Ephemeral mode
})
```

### Runtime Skills & Context Files

For multi-tenant hosts (chatbots, per-user agents, web services), the SDK
lets you swap skills and `AGENTS.md`-style context files **after** Headhunter
construction. Every mutation recomposes the system prompt and applies it to
the agent so the next turn picks up the new instructions — no restart needed.

```go
// Programmatic skill (no file on disk required).
host.AddSkill(&kit.Skill{
    Name:        "polite-french",
    Description: "Respond in French and always greet the user.",
    Content:     "Always reply in French. Open every response with 'Bonjour'.",
})

// Or load one from disk.
host.LoadAndAddSkill("/var/skills/refund-policy.md")

// Per-user AGENTS.md content pulled from a database.
host.AddContextFileContent(
    fmt.Sprintf("session://%s/AGENTS.md", userID),
    rulesFromDB,
)

// Tear down session-specific state on logout.
host.RemoveSkill("polite-french")
host.RemoveContextFile(fmt.Sprintf("session://%s/AGENTS.md", userID))

// Hide a skill from the model catalog without unloading it (still usable
// via /<name>); EnableSkill reverses it.
host.DisableSkill("refund-policy")
host.EnableSkill("refund-policy")

// Or replace the whole set atomically.
host.SetSkills(activeSkillsForUser)
host.SetContextFiles(activeContextForUser)
```

Skills dedupe by `Name`, context files dedupe by `Path` (which can be any
opaque identifier — it doesn't have to be a real filesystem path). All
mutators and readers (`GetSkills`, `GetContextFiles`) are safe to call
concurrently from multiple goroutines. See the [SDK overview docs](/sdk/overview#runtime-skills-and-context-files)
for the full reference.

### Runtime Native Tools

Native Go tools can also be added and removed on a live host, mirroring the
runtime MCP-server and skill APIs. This is useful for progressive disclosure
(dynamically loading a domain toolset only when the model asks for it) and for
multi-tenant hosts that need to swap tool catalogs without rebuilding the host.

```go
// Add tools that persist for the session.
host.AddTools(crmTools...)

// Drop a domain when it is no longer needed.
if err := host.RemoveTools("crm_search_contacts", "crm_create_deal"); err != nil {
    log.Printf("remove tools: %v", err)
}

// Replace the entire extra-tool set wholesale.
host.SetExtraTools(activeTools...)

// Read back the currently registered extra tools.
extra := host.GetExtraTools()
```

`AddTools` replaces existing extra tools by name (last-write-wins) and appends
new ones. `RemoveTools` is atomic: if any supplied name is not currently
registered, it returns an error and leaves the tool set unchanged. Core tools
and MCP tools are unaffected. Mutations take effect on the next LLM step.

## Advanced Usage

### Subagent Pattern

Spawn Headhunter as a subprocess for multi-agent orchestration:

```bash
headhunter "Analyze codebase" \
    --json \
    --no-session \
    --no-extensions \
    --quiet \
    --model anthropic/claude-haiku-3-5-20241022
```

Parse the JSON output:

```json
{
  "response": "Final assistant response text",
  "model": "anthropic/claude-haiku-3-5-20241022",
  "stop_reason": "end_turn",
  "session_id": "a1b2c3d4e5f6",
  "usage": {
    "input_tokens": 1024,
    "output_tokens": 512,
    "total_tokens": 1536,
    "cache_read_tokens": 0,
    "cache_creation_tokens": 0
  },
  "messages": [
    {
      "role": "assistant",
      "parts": [
        {"type": "text", "data": "..."},
        {"type": "tool_call", "data": {"name": "...", "args": "..."}},
        {"type": "tool_result", "data": {"name": "...", "result": "..."}}
      ]
    }
  ]
}
```

### Testing with tmux

Test the TUI non-interactively:

```bash
# Start Headhunter in detached tmux session
tmux new-session -d -s kittest -x 120 -y 40 \
  "kit -e ext.go --no-session 2>kit.log"

# Wait for startup
sleep 3

# Capture screen
tmux capture-pane -t kittest -p

# Send input
tmux send-keys -t kittest '/command' Enter

# Cleanup
tmux kill-session -t kittest
```

## Development

### Build and Test

```bash
# Build
go build -o output/headhunter ./cmd/headhunter

# Run tests
go test -race ./...

# Run specific test
go test -race ./cmd -run TestScriptExecution

# Lint
go vet ./...

# Format
go fmt ./...
```

### Project Structure

```
cmd/headhunter/             - CLI entry point (main.go)
cmd/                 - CLI command implementations (root, auth, models, etc.)
pkg/kit/             - Go SDK for embedding Headhunter
internal/app/        - Application orchestrator (agent loop, message store, queue)
internal/agent/      - Agent execution and tool dispatch
internal/auth/       - OAuth authentication and credential storage
internal/acpserver/  - ACP (Agent Client Protocol) server
internal/clipboard/  - Cross-platform clipboard operations
internal/compaction/ - Conversation compaction and summarization
internal/config/     - Configuration management
internal/core/       - Built-in tools (shell, read, write, edit, grep, find, ls)
internal/extensions/ - Yaegi extension system
internal/kitsetup/   - Initial setup wizard
internal/message/    - Message content types and structured content blocks
internal/models/     - Provider and model management
internal/session/    - Session persistence (tree-based JSONL)
internal/skills/     - Skill loading and system prompt composition
internal/tools/      - MCP tool integration
internal/ui/         - Bubble Tea TUI components
examples/extensions/ - Example extension files
npm/                 - NPM package wrapper for distribution
```

## Supported Providers

- **Anthropic** - Claude models (native, prompt caching, OAuth)
- **OpenAI** - GPT models
- **Copilot** - GitHub Copilot models (`copilot`, requires active Copilot subscription)
- **Google** - Gemini models
- **Ollama** - Local models
- **Azure OpenAI** - Azure-hosted OpenAI
- **AWS Bedrock** - Bedrock models
- **Google Vertex** - Claude on Vertex AI
- **OpenRouter** - Multi-provider router
- **Vercel AI** - Vercel AI SDK models
- **Custom** - Any OpenAI-compatible endpoint via `--provider-url`
- **Auto-routed** - Any provider from models.dev database
- **App-supplied (Go SDK)** - Your own inference backend under any provider name, via `kit.RegisterProvider` or `Options.Providers` (see [pkg/kit/README.md](pkg/kit/README.md#bundled-inference-backends-custom-providers))

### Custom Provider

Use `custom/custom` when pointing Headhunter at any OpenAI-compatible endpoint with `--provider-url`:

```bash
headhunter --provider-url "http://localhost:8080/v1" "Hello"
```

This automatically defaults to `custom/custom` without needing to specify a model. The custom provider routes through the `openaicompat` provider and supports:

- Zero cost tracking (input/output = 0)
- 262K context window, 65K output limit
- Reasoning and temperature support
- Optional `CUSTOM_API_KEY` environment variable or `--provider-api-key` flag

### Auto-routed Providers

Any provider in the [models.dev](https://models.dev) database can be used as
`provider/model` without a dedicated native integration. Headhunter auto-routes the
request through the matching **wire protocol** based on the provider's npm package
(or per-model override), using its `api` URL as the base:

| npm package | Wire protocol |
|-------------|---------------|
| `@ai-sdk/openai` | OpenAI (Responses API) |
| `@ai-sdk/openai-compatible` | OpenAI (chat completions) |
| `@ai-sdk/anthropic` | Anthropic |
| `@ai-sdk/google` | Google Gemini |

Providers with an `api` URL but an unrecognized npm package fall back to the
OpenAI-compatible wire. Because routing follows the wire protocol, aggregator/proxy
providers work across all of their models — including Claude, GPT, *and* Gemini
routes:

```bash
headhunter --model opencode/claude-haiku-4-5 "Hello"     # → Anthropic wire
headhunter --model opencode/gpt-5 "Hello"                # → OpenAI wire
headhunter --model opencode/gemini-3.5-flash "Hello"     # → Google wire
headhunter --model opencode/grok-4.6 "Hello"             # → OpenAI Responses wire
```

When an aggregator mixes wires inside one catalog without publishing a hint for
every family, Headhunter pins the family itself. opencode's `gpt-*` and `grok-*` models
answer only on the OpenAI Responses API — their `/chat/completions` route
returns HTTP 500/503 — so Headhunter routes them to the Responses wire automatically.

### Provider Overrides

Declare or patch providers in your config file's `providers` section — fix a
wire protocol the database gets wrong, or define an internal gateway the
database doesn't know about:

```yaml
# ~/.kit.yml
providers:
  minimax:                    # patch a known provider
    wire: anthropic
  corp-llm:                   # declare a new one
    wire: anthropic
    baseUrl: https://llm.internal.corp/api
    apiKeyEnv: [CORP_LLM_KEY]
    headers:
      X-Team: platform
```

Accepted `wire` values: `openai` (Responses API), `openai-compat` (chat
completions), `anthropic`, `google`. Unset fields inherit database values.
For one-off use, `--provider-wire` with `--provider-url` routes any
`provider/model` through the declared wire without a config entry.

### Model String Format

```bash
provider/model            # Standard format
anthropic/claude-sonnet-latest
openai/gpt-4o
ollama/llama3
google/gemini-2.0-flash-exp
```

### Model Aliases

```bash
# Anthropic Claude
claude-opus-latest        → claude-opus-4-6
claude-sonnet-latest      → claude-sonnet-4-6
claude-haiku-latest       → claude-haiku-4-5
claude-4-opus-latest      → claude-opus-4-6
claude-4-sonnet-latest    → claude-sonnet-4-6
claude-4-haiku-latest     → claude-haiku-4-5
claude-3-7-sonnet-latest  → claude-3-7-sonnet-20250219

# OpenAI GPT
o1-latest                 → o1
o3-latest                 → o3
o4-latest                 → o4-mini
gpt-5-latest              → gpt-5.4
gpt-5-chat-latest         → gpt-5.4
gpt-4-latest              → gpt-4o
gpt-4                     → gpt-4o
gpt-3.5-latest            → gpt-3.5-turbo
gpt-3.5                   → gpt-3.5-turbo

# Google Gemini
gemini-pro-latest         → gemini-2.5-pro
gemini-flash              → gemini-2.5-flash
gemini-pro                → gemini-2.5-pro
gemini-2-flash            → gemini-2.0-flash
gemini-2-pro              → gemini-2.5-pro
```

An alias applies only when the target model is available for the provider you
name, so `anthropic/claude-sonnet-latest` resolves while a provider that does
not carry that model reports a normal unknown-model error.

## Contributing

Contributions are welcome! Please see the [contribution guide](contribute/contribute.md) for guidelines.

## License

[MIT](LICENSE)

## Community

- [GitHub Issues](https://github.com/xnet-admin-1/Headhunter/issues)
- [Documentation](https://github.com/xnet-admin-1/Headhunter/wiki)
