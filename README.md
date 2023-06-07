# Datakit Tracing IDL

proto files for Datakit tracing generation automation.

generate Golang files for tracers used in Datakit.

tracers included in Datakit

- [open-telemetry proto idl](https://github.com/CodapeWild/opentelemetry-proto)
- [pinpoint proto idl](https://github.com/CodapeWild/pinpoint-grpc-idl)
- [skywalking proto idl](https://github.com/CodapeWild/skywalking-data-collect-protocol)

## How TO Use this repo

- generate all idl files just run `make` under this repo dir
- generate open-telemetry idl run:
  ```shell
    make clean
    make gen-opentelemetry
  ```
- generate pinpoint idl run:
  ```shell
    make clean
    make gen-pinpoint
  ```
- generate skywalking idl run:
  ```shell
    make clean
    make gen-skywalking
  ```