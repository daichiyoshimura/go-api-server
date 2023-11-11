APP_NAME=awsomeapp
ACCOUNT_DIR=internal/module/account


# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package server openapi.yaml > ./internal/server/types.gen.go
	oapi-codegen -generate "server" -package server openapi.yaml > ./internal/server/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go

# wire
WIRE_HANDLER_DIR=internal/handler
.PHONY: wire-handler
wire-handler:
	wire ${WIRE_HANDLER_DIR}/wire.go ${WIRE_HANDLER_DIR}/account.go ${WIRE_HANDLER_DIR}/iAccountUsecase.go ${WIRE_HANDLER_DIR}/handlers.go ${WIRE_HANDLER_DIR}/log.go

# wire
WIRE_USECASE_DIR=internal/module/account
.PHONY: wire-account
wire-usecase:
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

# mockgen account
ACCOUNT_DOMAIN_DIR = ${ACCOUNT_DIR}/internal/domain
.PHONY: mockgen-account
mockgen-account:
	mockgen -source=${ACCOUNT_DOMAIN_DIR}/iRepo.go -destination=${ACCOUNT_DOMAIN_DIR}/Repo_mock.go -package=domain -self_package=${APP_NAME}/${ACCOUNT_DOMAIN_DIR}