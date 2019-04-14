# avg

Dev build status: [![Concourse CI Build Status](https://ci.k8s.bell.wtf/api/v1/teams/main/pipelines/avg/jobs/build-dev/badge)](https://ci.k8s.bell.wtf/teams/main/pipelines/avg)

An average service.

## Code Generation

gRPC library code is generated from the Protobuf definition file in `/pkg/pb/*.proto` and is output to `/pkg/pb/*.pb.go`.

Two scripts are provided for make development a bit easier:

* `hack/update_codegen.sh` - runs `protoc` code generation to build the Go library files
* `hack/verify_codegen.sh` - verifies that the generated code is up to date.

The code is verified as part of the Docker build process.

## File Structure

```
chart/
ci/
  pipeline.yml
cmd/
  client/
  server/
hack/
pkg/
  pb/
```

`chart/` - helm chart definition

`ci` - continuous integration pipeline definition

`cmd/server` - code for the server CLI

`cmd/client` - code for the client CLI

`hack` - code generation scripts

`pkg/pb` - protobuf defintion and go bindings
