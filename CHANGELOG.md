# Changelog

## [Unreleased] — configs/novodb.conf y primer subpaquete (parse/)

- `novodb.conf` ahora se carga/crea en `configs/novodb.conf` en vez
  de la raíz del directorio de trabajo (`internal/novodb/constants.go`,
  `internal/novodb/defaults.go`); `SaveToFile` crea la carpeta
  `configs/` sola si hace falta. Actualizados `help.go`, `README.md`,
  `docs/configuration.md`, `docs/architecture.md`,
  `examples/quickstart.md` y `.gitignore` para reflejarlo.
- Nuevo subpaquete `internal/novodb/parse`: se movió el tokenizer de
  la sintaxis NQL (`tokenize` → `parse.Tokenize`), el único archivo
  del motor sin ninguna dependencia de `Engine`/`Document`/`Config`/
  `Session`/`Filter`/`Transaction`. `dsl_parser.go` y `cmd_view.go`
  ahora lo importan como `novodb/internal/novodb/parse`.
- `docs/known-limitations.md` ampliado: documenta este primer split
  seguro, incluye un filtro `grep` para encontrar más archivos "hoja"
  candidatos, y mantiene la ruta recomendada (`httpapi`, `cluster`,
  dejar `raft_fsm.go`/`transaction.go` en el núcleo) para quien
  quiera seguir dividiendo el paquete con un compilador Go a mano.

## [Unreleased] — Reorganización del repositorio

- Nueva estructura de carpetas a nivel de repo: `docs/` (con
  `docs/api/`), `deployments/docker/`, `scripts/`, `configs/`,
  `examples/`, `test/` (con `test/integration/` y `test/fixtures/`),
  `.github/workflows/`.
- Documentación nueva: arquitectura (`docs/architecture.md`),
  referencia NQL completa (`docs/nql-reference.md`), API HTTP
  (`docs/api/http-api.md`), configuración (`docs/configuration.md`),
  limitaciones conocidas (`docs/known-limitations.md`).
- Plantilla de configuración (`configs/novodb.conf.example`) generada
  a partir del struct `Config` real.
- `Dockerfile` + `docker-compose.yml` para levantar NovoDB en
  contenedor.
- Scripts de build/test/run (`scripts/*.sh`) y `Makefile`.
- Workflow de CI en GitHub Actions (`go build`, `go vet`, `go test`,
  `go mod tidy` check).
- `internal/novodb` se dejó intacto a propósito: es un único paquete
  Go con más de 40 archivos acoplados al mismo `Engine` mediante
  campos no exportados; dividirlo de verdad requiere exportar buena
  parte de ese estado y verificarlo con un compilador, algo que no
  fue posible confirmar en este entorno (ver
  `docs/known-limitations.md`).

## Pase anterior — Corrección de compilación

- Estructura `cmd/` + `internal/novodb/` en vez de un directorio
  plano de 68 archivos en la raíz.
- Eliminada la redeclaración duplicada de `views`/`viewsMu`.
- Implementados 11 manejadores NQL que `dsl_parser.go` referenciaba
  pero no existían (`handleDrop`, `handleRename`, `handleInfo`,
  `handleDescribe`, `handleStats`, `handleSize`, `handleRebuild`,
  `handleCheck`, `handleRepair`, `handleFlexCommand`,
  `handleTransaction`), en `cmd_admin_extra.go`.
- Eliminados imports no usados de `github.com/hashicorp/raft` y
  `errors` en 11 archivos.
