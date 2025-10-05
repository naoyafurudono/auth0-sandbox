.PHONY: generate
generate:
	oapi-codegen -config oapi-codegen.yaml api/openapi.yaml

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build -o bin/server cmd/server/main.go
