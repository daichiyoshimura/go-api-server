# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package server openapi.yaml > ./internal/server/types.gen.go
	oapi-codegen -generate "server" -package server openapi.yaml > ./internal/server/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go

# generate wire_gen.go
.PHONY: wire
wire:
	wire internal/di/wire.go

# golangci-lint
.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --fix

# run server
.PHONY: run
run:
	STAGE=DEV go run cmd/server/main.go