.PHONY: tools
tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint@latest
	npm i -g autorest

.PHONY: spec.update
spec.update:
	curl -o swagger.json https://circleci.com/api/v2/openapi.json
	prettier --write swagger.json
	git apply removedeprecation.diff
	autorest --v3 --go --namespace="ApiClient" --input-file="swagger.json" config-file.yaml --use=@autorest/go@4.0.0-preview.40 --title="circleci" --module-version="0.0.1"
	go run cmd/modpackage/main.go
	goimports -w . || true

.PHONY: lint
lint:
	golangci-lint run -E gofmt

.PHONY: fmt
fmt:
	go fmt ./...
	goimports -w .
