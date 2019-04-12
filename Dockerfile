FROM golang:1.11-alpine AS base
ENV GO111MODULE=on
RUN mkdir /build
WORKDIR /build
RUN apk add --no-cache ca-certificates git protobuf bash
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get -u github.com/golang/protobuf/protoc-gen-go
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .

FROM base AS tester
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./...

FROM base as builder
RUN ./hack/verify_codegen.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/avg-server

FROM alpine
ENV SITE_DIR /usr/share/www
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/avg-server /bin/avg-server
ENTRYPOINT [ "/bin/avg-server" ]