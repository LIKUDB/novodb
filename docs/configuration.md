# Configuración

NovoDB se configura por dos vías, que se combinan en este orden:

1. **Archivo `configs/novodb.conf`** (JSON), relativo al directorio
   de trabajo. Si no existe, se crea automáticamente al arrancar con
   los valores por defecto, creando también la carpeta `configs/` si
   hace falta (ver `internal/novodb/app.go` y
   `internal/novodb/defaults.go`). Una plantilla comentada está en
   [`configs/novodb.conf.example`](../configs/novodb.conf.example)
   — cópiala como `configs/novodb.conf` para personalizarla.
2. **Variables de entorno**, que sobrescriben lo que haya en el
   archivo (`internal/novodb/defaults.go`).

## Variables de entorno

| Variable | Descripción | Por defecto |
|---|---|---|
| `NOVODB_DATA` | Directorio raíz de datos | `./data` |
| `NOVODB_VM` | URL de VictoriaMetrics | — |
| `NOVODB_CACHE` | Caché L1 (MB) | `2048` |
| `NOVODB_L2_CACHE` | Entradas de índice en caché L2 | `4096` |
| `NOVODB_WORKERS` | Goroutines de trabajo | auto |
| `NOVODB_LOG_LEVEL` | `debug`\|`info`\|`warn`\|`error` | `info` |
| `NOVODB_METRICS` | Puerto de métricas | `9090` |
| `NOVODB_NODE_ID` | Identificador de nodo | auto |
| `NOVODB_RAFT_PORT` | Puerto Raft | `2335` |
| `NOVODB_RAFT_BIND` | Dirección de bind de Raft | `0.0.0.0` |
| `NOVODB_QUERY_PORT` | Puerto del servidor de consultas | `1555` |
| `NOVODB_ADMIN_PORT` | Puerto de la API de administración | `1556` |
| `NOVODB_SHARD_COUNT` | Número de shards | `16` |
| `NOVODB_REPLICA_COUNT` | Factor de replicación | `3` |
| `NOVODB_QUERY_USER` / `NOVODB_QUERY_PASSWORD` | Credenciales | — |
| `NOVODB_FAST_STARTUP` | Arranque rápido (sin replay de WAL) | `true` |
| `NOVODB_COMPRESS_LARGE` | Comprimir documentos grandes | `true` |
| `NOVODB_MAX_DOC_SIZE` | Tamaño máx. de documento (bytes) | `10485760` |
| `NOVODB_JWT_SECRET` | Secreto JWT | autogenerado |
| `NOVODB_TLS_ENABLED` | Activar TLS | `false` |
| `NOVODB_TLS_CERT_FILE` | Certificado TLS | — |

## Bloques de configuración principales

- **Core**: `data_root`, `vm_url`, `cache_mb`, `l2_cache_mb`, `workers`.
- **Red**: `raft_port`, `raft_bind`, `query_port`, `admin_port`.
- **Sharding**: `shard_count`, `replica_count`, `max_shards_per_node`,
  `auto_merge_enabled`, `auto_split_enabled`, `predictive_scaling`.
- **Seguridad**: `jwt_secret`, `token_expiry`, `max_login_attempts`,
  `lockout_duration`, `tls_enabled`.
- **FLEX-COLUMN** (motor columnar, bloque `flex_column`): caché
  columnar, detección de campos calientes, vistas materializadas,
  compresión (`zstd` por defecto).

Ver la estructura completa en
`internal/novodb/config.go` (struct `Config`) y
`internal/novodb/defaults.go` (valores por defecto + overrides por
entorno).
