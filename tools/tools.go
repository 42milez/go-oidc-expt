//go:build tools

package tools

import (
	_ "github.com/Masterminds/sprig"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "go.uber.org/mock/gomock"
	_ "golang.org/x/tools/cmd/goimports"
)
