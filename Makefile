PROTO_DIR=_proto
OUT_DIR=internal/grpc/generated

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

gen-proto:
	@if [ -z "$(PROTOC_GEN_GO)" ] || [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "protoc-gen-go or protoc-gen-go-grpc not found in PATH"; \
		echo "Install with: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; \
		echo "and: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; \
		exit 1; \
	fi
	protoc -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go_opt=Maegis-service.proto=github.com/lykeion-org/aegis/internal/grpc/generated/auth \
		--go-grpc_out=$(OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_opt=Maegis-service.proto=github.com/lykeion-org/aegis/internal/grpc/generated/auth \
		$(PROTO_DIR)/*.proto