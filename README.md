# TFORGE

Build & Test • GitHub Release • Go Report Card

Interactive infrastructure workspace for Terraform, Terragrunt and OpenTofu.

TFORGE is a next-generation TUI for infrastructure discovery, reusable module inspection, stack orchestration, realtime execution tracking, and infrastructure operations.

## Demo

Coming soon.

## Features

### Infrastructure Discovery

TFORGE automatically detects Terraform modules, Terragrunt stacks, OpenTofu projects, reusable modules, and runnable environments.

It supports monorepos, env/live repositories, nested infrastructure layouts, and multi-cloud repositories.

### Project Explorer

Interactive infrastructure explorer with auto-discovery, module/stack classification, keyboard navigation, search/filter, multi-select, and batch workflows.

Example:

```text
module terraform AWS/vpc
module terraform AWS/rds
stack  terragrunt live/prod
stack  terragrunt live/staging
```

### Module Inspector

Inspect reusable Terraform modules without execution.

Displays variables, outputs, resources, and providers.

Variable metadata includes type, default value, required/optional status, and descriptions.

### Execution Engine

Supports Terraform, Terragrunt, and OpenTofu.

Features include realtime logs, execution overlays, retry flows, apply confirmation, smart recovery, and execution history.

### Smart Recovery

TFORGE automatically detects missing providers, terraform init requirements, stale plans, and execution failures.

Example:

```text
Required plugins are not installed
Press I to run terraform init
```

## Install

### Go Install

```bash
go install github.com/Quixgit/tforge/cmd/tforge@latest
```

### Build From Source

```bash
git clone git@github.com:Quixgit/tforge.git
cd tforge

go build -o tforge ./cmd/tforge
```

## Usage

### Scan infrastructure repository

```bash
tforge --engine auto --dir .
```

### Scan specific directory

```bash
tforge --engine auto --dir /path/to/repository
```

### Terraform mode

```bash
tforge --engine terraform --dir ./infra
```

### Terragrunt mode

```bash
tforge --engine terragrunt --dir ./live
```

### OpenTofu mode

```bash
tforge --engine tofu --dir ./infra
```

## Repository Layout Support

TFORGE supports repositories like:

```text
modules/
  vpc/
  rds/

live/
  prod/
  staging/
```

Or:

```text
AWS/
Azure/
GCP/
```

Or large mixed infrastructure monorepos.

## Keyboard Shortcuts

### Explorer

| Key | Action |
|---|---|
| ↑ / ↓ | Navigate |
| Space | Select target |
| Enter | Open target |
| / | Search/filter |
| P | Plan selected |
| A | Apply selected |
| Esc | Close |

### Module Inspector

| Key | Action |
|---|---|
| 1 | Variables |
| 2 | Outputs |
| 3 | Resources |
| 4 | Providers |
| Esc | Back |

### Recovery

| Key | Action |
|---|---|
| I | Run terraform init |
| Ctrl+r | Retry |
| q | Quit |

## Current Status

Release stage: `alpha`

Implemented:
- project explorer
- reusable module detection
- module inspector
- live filtering
- execution overlays
- smart recovery
- risk engine foundation
- sensitive value masking
- execution history
- Terraform / Terragrunt / OpenTofu engine abstraction

Planned:
- dependency graph
- drift detection
- state inspection
- analytics dashboard
- workspace profiles
- stack graph visualization
- multi-environment orchestration

## Vision

TFORGE is building toward a full infrastructure operations platform:
- Terraform IDE
- infrastructure runtime console
- deployment orchestrator
- infrastructure observability workspace

Built for modern Infrastructure-as-Code workflows.
