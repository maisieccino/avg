# avg

[![Concourse CI Build Status](https://ci.k8s.bell.wtf/api/v1/teams/main/pipelines/avg/badge)](https://ci.k8s.bell.wtf/teams/main/pipelines/avg)

An average service.

## Code Generation

gRPC library code is generated from the Protobuf definition file in `/pkg/pb/*.proto` and is output to `/pkg/pb/*.pb.go`.

Two scripts are provided for make development a bit easier:

* `hack/update_codegen.sh` - runs `protoc` code generation to build the Go library files
* `hack/verify_codegen.sh` - verifies that the generated code is up to date.

The code is verified as part of the Docker build process.