# =====================================================
# PROTOBUF CODE GENERATION
# =====================================================
#
# Purpose: Generates Go code from Protocol Buffer definitions
# Usage: Run 'make proto' after modifying any .proto files
# See: docs/workflow.md#protocol-buffer-workflow for more details
#
# This makefile handles the generation of Go code from Protocol Buffer
# definitions, creating both message types and gRPC service interfaces.
# The generated code is placed alongside the .proto files.

# proto: Compile .proto files into Go gRPC code
# This target runs the protoc compiler with Go and gRPC plugins
# to generate type definitions and client/server interfaces
.PHONY: proto
proto:
	@echo "Generating protobuf code..."
	@which protoc > /dev/null 2>&1 || (echo "Error: protoc is not installed. Please install Protocol Buffer compiler first." && exit 1)
	@which protoc-gen-go > /dev/null 2>&1 || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc > /dev/null 2>&1 || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@protoc --go_out=. --go_opt=module=github.com/ChristiMahu/gaikeep \
		--go-grpc_out=. --go-grpc_opt=module=github.com/ChristiMahu/gaikeep \
		$(PROTO_DIR)/gaikeep/*.proto
	@echo "Protobuf generation complete."
