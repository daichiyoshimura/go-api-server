# oapi-codegen
.PHONY: openapi
openapi:
	oapi-codegen -generate "types" -package server openapi.yaml > ./internal/server/types.gen.go
	oapi-codegen -generate "server" -package server openapi.yaml > ./internal/server/server.gen.go
	oapi-codegen -generate "types" -package api openapi.yaml > ./api/types.gen.go
	oapi-codegen -generate "client" -package api openapi.yaml > ./api/client.gen.go