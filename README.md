# Datakit Tracing IDL

This repository stores tracing protocol IDL inputs and generated Go bindings used by Datakit.

The repository root is not a Go module. Each generated output directory is maintained as a separate Go module.

## Protocols

| Protocol | Upstream proto source | Generation ref | Generated Go module | Go version |
| --- | --- | --- | --- | --- |
| OpenTelemetry | `https://github.com/GuanceCloud/opentelemetry-proto` | `main` | `opentelemetry-gen-go` | `1.19` |
| Pinpoint | `git@github.com:CodapeWild/pinpoint-grpc-idl.git` | `v2.3.1-guance` | `pinpoint-gen-go` | `1.18` |
| SkyWalking | `git@github.com:CodapeWild/skywalking-data-collect-protocol.git` | `v9.4.0-guance` | `skywalking-gen-go` | `1.18` |

OpenTelemetry generated bindings include trace, metrics, logs, resources, common types, and profiles under `profiles/v1development` and `collector/profiles/v1development`.

## Requirements

Generation requires these tools in `PATH`:

- `protoc`
- `protoc-gen-go`
- `protoc-gen-go-grpc`
- `git`

The Makefile writes generated files to `${GOPATH}/src` using import paths, so this repository is expected to live at:

```sh
${GOPATH}/src/github.com/GuanceCloud/tracing-protos
```

## Generate Go Bindings

Generate all supported protocols:

```sh
make
```

Generate one protocol:

```sh
make gen-opentelemetry
make gen-pinpoint
make gen-skywalking
```

Cleanup targets:

```sh
make clean
make rm
```

Notes:

- `make clean` removes generated Go output directories.
- `make rm` removes upstream proto checkout directories.
- `make` runs `gen-all`, which runs `rm` before regenerating all protocols.
- OpenTelemetry proto files use `go.opentelemetry.io/proto/otlp/...` as upstream `go_package`; the Makefile maps them back into `github.com/GuanceCloud/tracing-protos/opentelemetry-gen-go/...`.

## Verification

Run compile checks from each generated module:

```sh
cd opentelemetry-gen-go
go test ./...
```

```sh
cd pinpoint-gen-go
go test ./...
```

```sh
cd skywalking-gen-go
go test ./...
```
