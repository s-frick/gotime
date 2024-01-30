package cmd

import (
	"context"

	gotime "github.com/s-frick/go-time-track/pkg"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Print each recorded frame.",
	Long:  "Print each recorded frame.",
	Run: func(cmd *cobra.Command, args []string) {
		var ctx context.Context
		if logJson {
			ctx = context.WithValue(cmd.Context(), gotime.ContextKeyLogType, gotime.JsonLog)
		} else if logCsv {
			ctx = context.WithValue(cmd.Context(), gotime.ContextKeyLogType, gotime.CsvLog)
		} else {
			ctx = context.WithValue(cmd.Context(), gotime.ContextKeyLogType, gotime.PrettyLog)
		}
		gotime.Log(ctx)
	},
}
