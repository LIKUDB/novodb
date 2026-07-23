# Known Limitations

## `internal/novodb` Remains a Single Go Package

This is the most significant limitation and the reason why this reorganization pass focused on **repository** structure (`docs/`, `deployments/`, `scripts/`, `configs/`, `examples/`, `test/`, CI) rather than splitting the Go package itself.

A concrete investigation was conducted into what would be required to split it into packages like `internal/engine`, `internal/storage`, `internal/cluster`, `internal/httpapi`, etc. Results of the static analysis (grepping for access to unexported `.field` / `.method` across files):

- Files like `raft_fsm.go` and `transaction.go` directly access over ten **unexported** fields/methods of the `Engine` from different subsystems (`engine.pool`, `engine.dirMgr`, `engine.cacheKey`, `engine.lockManager`, `engine.l1Cache`, `engine.shardMgr`, `engine.externalStore`, `engine.intelEngine`, `engine.flexEngine`, `engine.buildSecondaryIndex`, ...). Moving them to another package would force exporting a significant portion of the engine's internal state.
- Several subsystems (`cluster.go`, `dist_query.go`, `flexcolumn.go`, `http_admin.go`, `http_query.go`, `shard_manager.go`, `transaction.go`) store a back-reference to the `*Engine` itself (`engine *Engine`). If these types were moved to separate packages while `Engine` keeps them as fields, it would create import cycles. Properly avoiding this requires inverting the dependency with interfaces (defining in the new package only the methods it actually uses, and having `Engine` implicitly satisfy them). This is a real and reasonable change, but mechanical and extensive, with high risk of breaking a 15,000+ line build if done without a compiler to confirm every usage point.

This environment has no network access or Go toolchain installed, so there was no way to compile and confirm a change of that magnitude. The decision was made not to apply it blindly.

### What Was Already Separated: `internal/novodb/parse`

`tokenizer.go` (the `tokenize` function, now `parse.Tokenize`) had no references to `Engine`, `Document`, `Config`, `Session`, `Filter`, or `Transaction` — only standard library `strings`. Being a true "leaf" file (zero coupling, not just low coupling), it was moved without needing to export anything from the engine or risking import cycles. The two callers (`dsl_parser.go`, `cmd_view.go`) now import `novodb/internal/novodb/parse`.

### If You Want to Go Further

This is an approachable task with a Go compiler available (local or CI), subsystem by subsystem. A quick first filter to find more "leaf" candidates like `parse`:

```bash
cd internal/novodb
for f in *.go; do
  grep -qE "\bEngine\b|\bDocument\b|\bConfig\b|\bSession\b|\bFilter\b|\bTransaction\b" "$f" || echo "$f"
done
```

This flags ~18 files (`auth_jwt.go`, `cache.go`, `buffer_pool.go`, `compression.go`, `ratelimit.go`, `audit.go`, `logging.go`, `metrics.go`, `locks.go`, `worker_pool.go`, `storage_external.go`, `raft_logstore.go`, `misc_utils.go`, `nested_fields.go`, among others) that don't reference the core engine types by name — they're good candidates for their own packages, but **before moving them**, you need to review the coupling *between them* (e.g., if several share `log()` from `logging.go` or a cache type), which this filter doesn't detect. After that, in order of least to most coupling with the `Engine`:

1. `internal/httpapi` — move `http_admin.go` and `http_query.go`; only requires exporting ~10 `Engine` fields (`nodeID`, `startupTime`, `opCount`, `l1Cache`, `metrics`, `tokenManager`, `cluster`, `flexEngine`, `shardMgr`, `config`) via getters.
2. `internal/cluster` — move `cluster.go`, `dist_query.go`, `flexcolumn.go`, `shard_manager.go`; moderate coupling.
3. Keep `raft_fsm.go` and `transaction.go` in the core package (`internal/engine`) — these most deeply touch internal state and have the least benefit/risk ratio when separated.

Each step should conclude with `go build ./... && go vet ./...` before moving to the next.

## `go.sum` Not Verified

`go.mod` remains the same as in the original project. Run `go mod tidy` on your machine after the first successful build to confirm that `go.sum` is consistent.

## No Automated Test Suite

The original project didn't include tests. `test/` was added with a template and a guide (`test/README.md`) to get started, but no tests were invented that simulate coverage that doesn't exist.
