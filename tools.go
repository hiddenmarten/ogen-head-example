package tools

//go:generate go run scripts/genoas.go
//go:generate go tool ogen --target api/client -package client --clean api/oas/openapi.yaml
