# syntax=docker/dockerfile:1

FROM golang:1.22-bookworm AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/novodb ./cmd/novodb

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /out/novodb /app/novodb
COPY configs/novodb.conf.example /app/novodb.conf

VOLUME ["/app/data"]
EXPOSE 1555 1556 9090 2335

ENTRYPOINT ["/app/novodb"]
