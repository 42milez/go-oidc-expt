//go:build tools

package tools

import (
	_ "github.com/Masterminds/sprig"
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/golang/mock/gomock"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
