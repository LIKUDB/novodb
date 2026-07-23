# HTTP API

NovoDB exposes two independent HTTP servers, each on its own configurable port.

## Admin Server (`internal/novodb/http_admin.go`)

Default port: `admin_port` (`NOVODB_ADMIN_PORT`, 1556 if not configured).

| Route | Description |
|---|---|
| `GET /api/v1/health` | Health check |
| `GET /api/v1/status` | Engine status |
| `GET /api/v1/dbs` | List databases |
| `/api/v1/db/...` | Operations on a specific database |
| `POST /api/v1/query` | Execute an NQL query |
| `POST /api/v1/backup` | Backup a database |
| `POST /api/v1/restore` | Restore a database |
| `POST /api/v1/compact` | Compaction |
| `/api/v1/users` | User management |
| `POST /api/v1/auth/login` | Login (returns JWT) |
| `POST /api/v1/auth/logout` | Logout |
| `/api/v1/token` | Token management |
| `GET /api/v1/metrics` | Prometheus metrics (`promhttp`) |

## Query Server (`internal/novodb/http_query.go`)

Default port: `query_port` (`NOVODB_QUERY_PORT`, 1555 if not configured).

| Route | Description |
|---|---|
| `POST /query` | Execute an NQL query |
| `GET /status` | Engine status |
| `GET /health` | Health check |

## Authentication

Both servers support JWT (`auth_jwt.go`) and Basic Auth as alternatives. Default credentials are configured with `query_user` / `query_password` (or `NOVODB_QUERY_USER` / `NOVODB_QUERY_PASSWORD`) and the JWT secret with `jwt_secret` (or `NOVODB_JWT_SECRET`).
