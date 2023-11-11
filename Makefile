APP_NAME=awsomeapp
HANDLER_DIR=internal/handler
ACCOUNT_DIR=internal/module/account

# golangci-lint
.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --fix

# run server
.PHONY: run
run:
	STAGE=DEV go run cmd/server/main.go

# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package server openapi.yaml > ./internal/server/types.gen.go
	oapi-codegen -generate "server" -package server openapi.yaml > ./internal/server/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go

# wire
.PHONY: wire
wire:
	wire ${HANDLER_DIR}/wire.go ${HANDLER_DIR}/account.go ${HANDLER_DIR}/iAccountUsecase.go ${HANDLER_DIR}/accountUsecase_mock.go ${HANDLER_DIR}/handlers.go ${HANDLER_DIR}/log.go
	wire ${ACCOUNT_DIR}/wire.go ${ACCOUNT_DIR}/iRepo.go ${ACCOUNT_DIR}/Repo_mock.go ${ACCOUNT_DIR}/usecase.go	

# mockgen account
ACCOUNT_DOMAIN_DIR=${ACCOUNT_DIR}/internal/domain
.PHONY: mockgen
mockgen:
	mockgen -source=${ACCOUNT_DOMAIN_DIR}/iRepo.go -destination=${ACCOUNT_DOMAIN_DIR}/Repo_mock.go -package=domain -self_package=${APP_NAME}/${ACCOUNT_DOMAIN_DIR}
	mockgen -source=${ACCOUNT_DIR}/iRepo.go -destination=${ACCOUNT_DIR}/Repo_mock.go -package=account -self_package=${APP_NAME}/${ACCOUNT_DIR}	
	mockgen -source=${HANDLER_DIR}/iAccountUsecase.go -destination=${HANDLER_DIR}/accountUsecase_mock.go -package=handler -self_package=${APP_NAME}/${HANDLER_DIR}	