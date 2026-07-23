# Changelog

## [Unreleased] — configs/novodb.conf and first subpackage (parse/)

- `novodb.conf` now loads/creates from `configs/novodb.conf` instead of the working directory root (`internal/novodb/constants.go`, `internal/novodb/defaults.go`); `SaveToFile` creates the `configs/` folder on its own if needed. Updated `help.go`, `README.md`, `docs/configuration.md`, `docs/architecture.md`, `examples/quickstart.md`, and `.gitignore` to reflect this.
- New subpackage `internal/novodb/parse`: moved the NQL syntax tokenizer (`tokenize` → `parse.Tokenize`), the only engine file without any dependency on `Engine`/`Document`/`Config`/`Session`/`Filter`/`Transaction`. `dsl_parser.go` and `cmd_view.go` now import it as `novodb/internal/novodb/parse`.
- `docs/known-limitations.md` expanded: documents this first safe split, includes a `grep` filter to find more "leaf" file candidates, and maintains the recommended path (`httpapi`, `cluster`, keep `raft_fsm.go`/`transaction.go` in the core) for anyone wanting to continue splitting the package with a Go compiler available.

## [Unreleased] — Repository Reorganization

- New repository-level folder structure: `docs/` (with `docs/api/`), `deployments/docker/`, `scripts/`, `configs/`, `examples/`, `test/` (with `test/integration/` and `test/fixtures/`), `.github/workflows/`.
- New documentation: architecture (`docs/architecture.md`), complete NQL reference (`docs/nql-reference.md`), HTTP API (`docs/api/http-api.md`), configuration (`docs/configuration.md`), known limitations (`docs/known-limitations.md`).
- Configuration template (`configs/novodb.conf.example`) generated from the actual `Config` struct.
- `Dockerfile` + `docker-compose.yml` to run NovoDB in a container.
- Build/test/run scripts (`scripts/*.sh`) and `Makefile`.
- GitHub Actions CI workflow (`go build`, `go vet`, `go test`, `go mod tidy` check).
- `internal/novodb` was intentionally left intact: it's a single Go package with over 40 files coupled to the same `Engine` via unexported fields; truly splitting it requires exporting a significant portion of that state and verifying with a compiler, which wasn't possible to confirm in this environment (see `docs/known-limitations.md`).

## Previous Pass — Compilation Fixes

- `cmd/` + `internal/novodb/` structure instead of a flat directory of 68 files at the root.
- Removed duplicate redeclaration of `views`/`viewsMu`.
- Implemented 11 NQL handlers that `dsl_parser.go` referenced but didn't exist (`handleDrop`, `handleRename`, `handleInfo`, `handleDescribe`, `handleStats`, `handleSize`, `handleRebuild`, `handleCheck`, `handleRepair`, `handleFlexCommand`, `handleTransaction`), in `cmd_admin_extra.go`.
- Removed unused imports of `github.com/hashicorp/raft` and `errors` in 11 files.
