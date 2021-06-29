proto_version := 3.17.3

.PHONY: help
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: initialize
initialize: ## task description
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	mkdir -p .bin/tmp
	curl -sSfL \
		https://github.com/protocolbuffers/protobuf/releases/download/v${proto_version}/protoc-${proto_version}-linux-x86_64.zip -o .bin/tmp/protoc.zip
	cd .bin/tmp && unzip protoc.zip
	mv .bin/tmp/bin/protoc .bin/
	rm -rf .bin/tmp
