default: help

.PHONY: help
help: pad = 24 # padding for two columns
help:	## Show this help
	@echo "Commands:"
	@fgrep -h "##" $(MAKEFILE_LIST) \
		| fgrep -v fgrep \
		| sed -e 's/^/  /' -e 's/:/ /' -e 's/	//g' \
		| sort -k 1 \
		| grep -v '^  #' \
		| awk -F "#" '{printf ("%s% *s%s\n", $$1, $(pad)-length($$1), "", $$3)}'
	@echo

.PHONY: run
run:  ## Run application
	go run cmd/client/main.go

.PHONY: grpc-gen-client
grpc-gen-client:  ## Generate gRPC code from proto files
	mkdir -p pkg/client/grpc

	protoc --proto_path=deps/grpc \
		--go_out=pkg/client/grpc --go_opt=paths=source_relative \
		--go-grpc_out=pkg/client/grpc --go-grpc_opt=paths=source_relative \
		deps/grpc/user/*.proto
