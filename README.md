# NovoDB

Document-oriented data engine with sharding, Raft-based clustering, WAL, ACID transactions, and an SQL-like command language (NQL).

## Wiki

- 🇪🇸 [Español](https://github.com/novodb/novodb/wiki/es)
- 🇬🇧 [English](https://github.com/novodb/novodb/wiki/en)
- 🇩🇪 [Deutsch](https://github.com/novodb/novodb/wiki/de)

## Repository Structure

```
novodb/
├── cmd/
│   └── novodb/              # entry point (func main)
├── internal/
│   └── novodb/               # engine, server, and CLI (single Go package)
│       └── parse/            # NQL syntax tokenizer
├── configs/
│   ├── novodb.conf.example   # configuration template (JSON)
│   └── novodb.conf           # active config; auto-created on first startup
├── deployments/
│   └── docker/
│       ├── Dockerfile
│       └── docker-compose.yml
├── docs/
│   ├── architecture.md
│   ├── configuration.md
│   ├── nql-reference.md
│   ├── known-limitations.md
│   └── api/
│       └── http-api.md
├── examples/
│   └── quickstart.md
├── scripts/
│   ├── build.sh
│   ├── test.sh
│   └── run-dev.sh
├── test/
│   ├── integration/
│   ├── fixtures/
│   └── README.md
├── .github/workflows/ci.yml
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
└── go.mod
```

## Build and Run

```bash
go build ./cmd/novodb
./novodb
```

or with the included utilities:

```bash
make build   # ./scripts/build.sh -> bin/novodb
make run     # ./scripts/run-dev.sh (development mode, ./data-dev)
make test    # go build + go vet + go test
make docker  # build Docker image
```

## Configuration

`configs/novodb.conf` (JSON), relative to the working directory, with overrides via environment variables (`NOVODB_*`). If it doesn't exist, NovoDB auto-creates it with default values. Full details in [`docs/configuration.md`](docs/configuration.md); template at [`configs/novodb.conf.example`](configs/novodb.conf.example).

## Changelog

See [`CHANGELOG.md`](CHANGELOG.md) for the complete history.

## Contributing

See [`CONTRIBUTING.md`](CONTRIBUTING.md) for guidelines.

## License

Apache License, Version 2.0. See [`LICENSE`](LICENSE) for details.

---

## Pending Verification

**This environment had no access to a Go compiler or network.** Before trusting this code, on a machine with Go 1.22+:

```bash
go build ./...
go vet ./...
go mod tidy   # and confirm go.sum is consistent
```

If anything fails, that's the first place to start.
