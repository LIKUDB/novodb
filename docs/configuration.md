# Configuration

NovoDB is configured through two channels, which are combined in this order:

1. **`configs/novodb.conf` file** (JSON), relative to the working directory. If it doesn't exist, it's automatically created on startup with default values, also creating the `configs/` folder if needed (see `internal/novodb/app.go` and `internal/novodb/defaults.go`). A commented template is available at [`configs/novodb.conf.example`](../configs/novodb.conf.example) — copy it as `configs/novodb.conf` to customize it.
2. **Environment variables**, which override whatever is in the file (`internal/novodb/defaults.go`).

## Environment Variables

| Variable | Description | Default |
|---|---|---|
| `NOVODB_DATA` | Root data directory | `./data` |
| `NOVODB_VM` | VictoriaMetrics URL | — |
| `NOVODB_CACHE` | L1 cache (MB) | `2048` |
| `NOVODB_L2_CACHE` | L2 cache index entries | `4096` |
| `NOVODB_WORKERS` | Worker goroutines | auto |
| `NOVODB_LOG_LEVEL` | `debug`\|`info`\|`warn`\|`error` | `info` |
| `NOVODB_METRICS` | Metrics port | `9090` |
| `NOVODB_NODE_ID` | Node identifier | auto |
| `NOVODB_RAFT_PORT` | Raft port | `2335` |
| `NOVODB_RAFT_BIND` | Raft bind address | `0.0.0.0` |
| `NOVODB_QUERY_PORT` | Query server port | `1555` |
| `NOVODB_ADMIN_PORT` | Admin API port | `1556` |
| `NOVODB_SHARD_COUNT` | Number of shards | `16` |
| `NOVODB_REPLICA_COUNT` | Replication factor | `3` |
| `NOVODB_QUERY_USER` / `NOVODB_QUERY_PASSWORD` | Credentials | — |
| `NOVODB_FAST_STARTUP` | Fast startup (skip WAL replay) | `true` |
| `NOVODB_COMPRESS_LARGE` | Compress large documents | `true` |
| `NOVODB_MAX_DOC_SIZE` | Max document size (bytes) | `10485760` |
| `NOVODB_JWT_SECRET` | JWT secret | auto-generated |
| `NOVODB_TLS_ENABLED` | Enable TLS | `false` |
| `NOVODB_TLS_CERT_FILE` | TLS certificate file | — |

## Main Configuration Blocks

- **Core**: `data_root`, `vm_url`, `cache_mb`, `l2_cache_mb`, `workers`.
- **Network**: `raft_port`, `raft_bind`, `query_port`, `admin_port`.
- **Sharding**: `shard_count`, `replica_count`, `max_shards_per_node`, `auto_merge_enabled`, `auto_split_enabled`, `predictive_scaling`.
- **Security**: `jwt_secret`, `token_expiry`, `max_login_attempts`, `lockout_duration`, `tls_enabled`.
- **FLEX-COLUMN** (columnar engine, `flex_column` block): columnar cache, hot field detection, materialized views, compression (`zstd` by default).

See the complete structure in `internal/novodb/config.go` (struct `Config`) and `internal/novodb/defaults.go` (default values + environment overrides).
