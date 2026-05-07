# TFORGE

Next-generation infrastructure execution console for Terraform, Terragrunt and OpenTofu.

TFORGE provides:
- realtime infrastructure execution tracking
- resource lifecycle visualization
- live execution overlays
- safe apply workflows
- multi-engine support
- execution analytics
- workspace management
- history persistence
- infrastructure observability

## Features

### Infrastructure Engines

- Terraform
- Terragrunt
- OpenTofu

### Execution Intelligence

- Live execution tracker
- Realtime resource status
- Execution overlays
- Progress tracking
- Cached plans
- Apply confirmation flow
- History persistence

### UI

- Interactive TUI
- Resource tree
- Detail overlays
- Analytics overlays
- Workspace switcher
- Keyboard-driven workflow

## Installation

### Requirements

- Go 1.26+
- Terraform / Terragrunt / OpenTofu

### Build

    git clone git@github.com:Quixgit/tforge.git
    cd tforge
    go build -o tforge ./cmd/tforge

## Usage

### Terraform

    ./tforge --engine terraform --dir .

### Auto detect

    ./tforge --engine auto --dir .

## Keyboard Shortcuts

- j/k → navigate
- Space → select
- Enter → details
- Tab → action menu
- Ctrl+r → refresh
- E → execution
- P → providers
- A → analytics
- W → workspaces
- Y → history
- q → quit

## Status

Current release stream:

    v0.x.x-alpha.x

TFORGE is currently under active development.

## License

MIT
