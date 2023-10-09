# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package account openapi.yaml > ./internal/account/types.gen.go
	oapi-codegen -generate "server" -package account openapi.yaml > ./internal/account/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go