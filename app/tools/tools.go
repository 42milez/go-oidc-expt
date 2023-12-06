//go:build tools

package tools

import (
	_ "github.com/Masterminds/sprig"
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "go.uber.org/mock/gomock"
)
