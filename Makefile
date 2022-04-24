.PHONY: tools
tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint@latest

.PHONY: spec.update
spec.update:
	curl -o swagger.json https://circleci.com/api/v2/openapi.json
	prettier --write swagger.json
	git apply patch.diff
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/swagger.json -g go -o /local/client -p packageName=client,isGoSubmodule=true
	goimports -w . || true
#	oapi-codegen -generate "client" -package client swagger.json > client/client.go || true

.PHONY: lint
lint:
	golangci-lint run -E gofmt

.PHONY: fmt
fmt:
	go fmt ./...
	goimports -w .
