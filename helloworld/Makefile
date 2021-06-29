.PHONY: help
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: grpc_generate
grpc_generate: ## task description
	protoc --go_out=helloworld \
		--go_opt=paths=source_relative \
		--go-grpc_out=helloworld \
		--go-grpc_opt=paths=source_relative \
		-I=../proto \
		../proto/helloworld.proto
