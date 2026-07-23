# API HTTP

NovoDB expone dos servidores HTTP independientes, cada uno en su
propio puerto configurable.

## Servidor de administración (`internal/novodb/http_admin.go`)

Puerto por defecto: `admin_port` (`NOVODB_ADMIN_PORT`, 1556 si no se
configura).

| Ruta | Descripción |
|---|---|
| `GET /api/v1/health` | Health check |
| `GET /api/v1/status` | Estado del motor |
| `GET /api/v1/dbs` | Listar bases de datos |
| `/api/v1/db/...` | Operaciones sobre una base de datos concreta |
| `POST /api/v1/query` | Ejecutar una consulta NQL |
| `POST /api/v1/backup` | Backup de una base de datos |
| `POST /api/v1/restore` | Restore de una base de datos |
| `POST /api/v1/compact` | Compactación |
| `/api/v1/users` | Gestión de usuarios |
| `POST /api/v1/auth/login` | Login (devuelve JWT) |
| `POST /api/v1/auth/logout` | Logout |
| `/api/v1/token` | Gestión de tokens |
| `GET /api/v1/metrics` | Métricas Prometheus (`promhttp`) |

## Servidor de consultas (`internal/novodb/http_query.go`)

Puerto por defecto: `query_port` (`NOVODB_QUERY_PORT`, 1555 si no se
configura).

| Ruta | Descripción |
|---|---|
| `POST /query` | Ejecutar una consulta NQL |
| `GET /status` | Estado del motor |
| `GET /health` | Health check |

## Autenticación

Ambos servidores soportan JWT (`auth_jwt.go`) y Basic Auth como
alternativa. Las credenciales por defecto se configuran con
`query_user` / `query_password` (o `NOVODB_QUERY_USER`) y el secreto
JWT con `jwt_secret` (o `NOVODB_JWT_SECRET`).
