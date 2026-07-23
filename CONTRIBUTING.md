# Contributing to NovoDB

## Before You Start

This repository was prepared/reorganized in an environment without a Go toolchain or network access, so **nothing here was actually compiled**. Before your first change:

```bash
go build ./...
go vet ./...
```

and report any errors — that's the first thing to fix.

## Structure

See [`docs/architecture.md`](docs/architecture.md) for the complete code map. Quick summary:

- `cmd/novodb/` — entry point, should not grow.
- `internal/novodb/` — the entire engine (single Go package, see [`docs/known-limitations.md`](docs/known-limitations.md) for why). Group your new file with the appropriate prefix (`ops_`, `cmd_`, `http_`, etc.) instead of creating a new folder.
- `docs/` — documentation.
- `deployments/` — Docker/compose.
- `scripts/` — build/test/run.
- `test/` — integration and fixtures.

## Style

- `gofmt -l .` should return nothing before a PR.
- Follow the existing file prefix convention in `internal/novodb/`.
- New NQL commands go in `cmd_*.go` + their entry in `dsl_parser.go` + documentation in `docs/nql-reference.md` and in the `HELP` text (`help.go`).

## Pull Requests

1. Descriptive branch (`fix/...`, `feat/...`).
2. `go build ./... && go vet ./... && go test ./...` passes.
3. Update affected documentation (`docs/`, `README.md`, `CHANGELOG.md`).
