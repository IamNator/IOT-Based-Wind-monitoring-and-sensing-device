#!/bin/sh
go get github.com/99designs/gqlgen/cmd@v0.14.0
echo "running go generate"
go generate ./...
echo "running go fmt"
go fmt ./...
# binary will be $(go env GOPATH)/bin/golangci-lin
#echo "installing golandci-lint"
#curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
#golangci-lint --version
echo "running golangci-lint"
golangci-lint run ./... -v
echo "running tests with coverage flag"
# shellcheck disable=SC2006
go test `go list ./... | grep -v example` -coverprofile=coverage.txt -covermode=atomic
echo "Genius is in simplicity!"
echo "done!"