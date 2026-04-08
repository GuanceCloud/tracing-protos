# AGENT.md

## Project Overview

This repository stores protobuf IDL inputs and generated Go bindings used by Datakit tracing integrations.

It currently covers three tracer/protocol families:

- OpenTelemetry: upstream proto submodule `opentelemetry-proto`, generated module `opentelemetry-gen-go`
- Pinpoint: upstream proto submodule `pinpoint-grpc-idl`, generated module `pinpoint-gen-go`
- SkyWalking: upstream proto submodule `skywalking-data-collect-protocol`, generated module `skywalking-gen-go`

The root of the repository is not a Go module. Each generated output directory is its own Go module.

## Repository Layout

- `Makefile`: generation entry points and upstream ref configuration.
- `.gitmodules`: upstream proto repositories. OpenTelemetry uses `https://github.com/GuanceCloud/opentelemetry-proto`; Pinpoint and SkyWalking currently use `git@github.com:CodapeWild/...`.
- `opentelemetry-proto`: OpenTelemetry proto submodule.
- `pinpoint-grpc-idl`: Pinpoint proto submodule.
- `skywalking-data-collect-protocol`: SkyWalking proto submodule.
- `opentelemetry-gen-go`: generated Go module for OpenTelemetry proto files.
- `pinpoint-gen-go`: generated Go module for Pinpoint proto files.
- `skywalking-gen-go`: generated Go module for SkyWalking proto files.

## Generated Go Modules

Generated module Go versions:

- `opentelemetry-gen-go`: `go 1.19`
- `pinpoint-gen-go`: `go 1.18`
- `skywalking-gen-go`: `go 1.18`

All three generated modules currently use:

- `google.golang.org/grpc v1.56.2`
- `google.golang.org/protobuf v1.31.0`

Module paths:

- `github.com/GuanceCloud/tracing-protos/opentelemetry-gen-go`
- `github.com/GuanceCloud/tracing-protos/pinpoint-gen-go`
- `github.com/GuanceCloud/tracing-protos/skywalking-gen-go`

The committed `.pb.go` files are generated code. Prefer changing upstream proto inputs and regenerating instead of hand-editing generated files, unless the task explicitly asks for a narrow generated-code patch.

## Generation Workflow

Generation is driven by `Makefile` and requires these tools in `PATH`:

- `protoc`
- `protoc-gen-go`
- `protoc-gen-go-grpc`
- `git`

The Makefile writes generated files to `${GOPATH}/src` using `--go_opt=paths=import`, so this repository is expected to live at:

```sh
${GOPATH}/src/github.com/GuanceCloud/tracing-protos
```

Current upstream ref configuration in `Makefile`:

- OpenTelemetry: `main`
- Pinpoint: `v2.3.1-guance`
- SkyWalking: `v9.4.0-guance`

Common commands:

```sh
make
make gen-opentelemetry
make gen-pinpoint
make gen-skywalking
make clean
make rm
```

Important behavior:

- `make` runs `gen-all`, which first runs `rm` and deletes the upstream proto directories, then regenerates all three outputs.
- `make clean` deletes the generated Go output directories: `opentelemetry-gen-go`, `pinpoint-gen-go`, and `skywalking-gen-go`.
- `make rm` deletes the upstream proto directories: `opentelemetry-proto`, `pinpoint-grpc-idl`, and `skywalking-data-collect-protocol`.
- Individual generation targets expect the corresponding upstream proto directory to be a usable git checkout. If submodule directories exist but are empty, run `git submodule update --init --recursive` first, or use the Makefile flow that recreates/clones them.
- The Makefile checks out the configured ref before generation, then switches OpenTelemetry back to `main` and Pinpoint/SkyWalking back to `master`.
- OpenTelemetry upstream proto files use `go.opentelemetry.io/proto/otlp/...` as `go_package`; keep the explicit `M...=github.com/GuanceCloud/tracing-protos/opentelemetry-gen-go/...` mappings in `Makefile` so generation lands in this repository.
- OpenTelemetry generated bindings include profiles packages under `opentelemetry-gen-go/profiles/v1development` and `opentelemetry-gen-go/collector/profiles/v1development`.

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

The current repository state compiles successfully with these commands. Packages currently have no test files.

## Agent Notes

- Check `git status --short` before editing. Do not overwrite unrelated user changes.
- Keep edits scoped. This repo is mostly generated protobuf bindings, so broad refactors are usually inappropriate.
- When updating generated code, record the source tag or submodule commit used.
- Be careful with `make`, `make clean`, and `make rm`; they delete directories as part of the generation flow.
- If only documentation changes are requested, avoid regenerating protobuf outputs.
- If submodules appear as empty directories, treat them as uninitialized before assuming proto sources are missing.
