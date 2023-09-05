//go:build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/mattn/go-runewidth"
	_ "github.com/olekukonko/tablewriter"
	_ "github.com/rivo/uniseg"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
)
