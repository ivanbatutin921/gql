.PHONY: build

# Define variables for paths and options
PROTOC = protoc
PROTOC_GEN_GRPC = protoc-gen-go-grpc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GRAPHQL = protoc-gen-graphql
INCLUDE_PATH = ../../include/graphql
OUTPUT_PATH = ./greeter
PROTO_FILE = greeter.proto

build:
	$(PROTOC) -I. \
		-I$(INCLUDE_PATH) \
		--plugin=$(PROTOC_GEN_GRAPHQL) \
		--go_out=$(OUTPUT_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_PATH) --go-grpc_opt=paths=source_relative \
		--graphql_out=$(OUTPUT_PATH) --graphql_opt=paths=source_relative \
		$(PROTO_FILE)
