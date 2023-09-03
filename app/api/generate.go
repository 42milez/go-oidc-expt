package api

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package api -generate "types,chi-server,spec" -o api.go spec/spec.yml
