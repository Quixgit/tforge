# TFORGE

<p align="center">
  <strong>Next-generation infrastructure execution console</strong>
</p>

<p align="center">
  Terraform • Terragrunt • OpenTofu
</p>

<p align="center">
  Realtime execution • Infrastructure observability • Safe apply workflows
</p>

---

TFORGE is a modern terminal UI and execution layer for Infrastructure as Code.

It transforms Terraform, Terragrunt and OpenTofu workflows into a realtime operational console with:
- live execution tracking
- execution overlays
- resource lifecycle visibility
- execution analytics
- safe apply flows
- workspace management
- execution history
- infrastructure observability

---

# Vision

Traditional Terraform UX is:
- noisy
- opaque
- difficult to observe
- hard to reason about during execution

TFORGE aims to become:

> the realtime operational cockpit for Infrastructure as Code

---

# Features

## Infrastructure Engines

- Terraform
- Terragrunt
- OpenTofu

## Execution Intelligence

- Live execution tracker
- Realtime resource lifecycle tracking
- Execution overlays
- Progress visualization
- Cached execution plans
- Apply confirmation flows
- Resource execution analytics
- Execution history persistence

## UI

- Interactive terminal UI
- Resource tree view
- Resource detail overlays
- Execution overlays
- Workspace switcher
- Analytics dashboard
- Keyboard-driven workflow

## Safety

- Cached plan apply
- Safe execution layer
- Structured command execution
- No shell interpolation
- History isolation
- Local cache permissions

---

# Screenshots

Coming soon.

---

# Installation

## Requirements

- Go 1.26+
- Terraform / Terragrunt / OpenTofu

---

## Build from source

```bash
git clone git@github.com:Quixgit/tforge.git
cd tforge

go build -o tforge ./cmd/tforge
