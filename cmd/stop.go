package cmd

import (
	gotime "github.com/s-frick/go-time-track/pkg"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current time frame",
	Long:  "Stop the current time frame",
	Run: func(cmd *cobra.Command, args []string) {
		gotime.Stop(cmd.Context(), gotime.Options{At: at})
	},
}
