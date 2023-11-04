# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package server openapi.yaml > ./internal/server/types.gen.go
	oapi-codegen -generate "server" -package server openapi.yaml > ./internal/server/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go

# generate wire_gen.go
WIRE_USECASE_DIR=internal/module/account
.PHONY: wire
wire:
	wire internal/di/wire.go
	wire ${WIRE_USECASE_DIR}/wire.go ${WIRE_USECASE_DIR}/usecase.go ${WIRE_USECASE_DIR}/iAccountRepo.go

# golangci-lint
.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --fix

# run server
.PHONY: run
run:
	STAGE=DEV go run cmd/server/main.go

# mockgen repository FileName=${}
MOCKGEN_DIR=internal/module/account/internal/domain
.PHONY: mockgen
mockgen:
	mockgen -source=${MOCKGEN_DIR}/${FILE_NAME}.go -destination=${MOCKGEN_DIR}/mock/${FILE_NAME}_mock.go -package=mock