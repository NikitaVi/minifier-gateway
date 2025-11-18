include .env

LOCAL_BIN:=$(CURDIR)/bin
PROTOC = PATH=$(LOCAL_BIN):$$PATH protoc
PROTO_REPO := $(shell go list -m -f '{{.Dir}}' github.com/NikitaVi/minifier-protos)

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	$(PROTOC) -I $(PROTO_REPO)/proto \
	  $(PROTO_REPO)/proto/auth_v1/auth_v1.proto \
	  --go_out=./pkg --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative
