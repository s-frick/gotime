package cmd

import (
	gotime "github.com/s-frick/go-time-track/pkg"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new time frame",
	Long:  "Start a new time frame",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := make([]gotime.Tag, len(args))
		for i, e := range args {
			t[i] = gotime.Tag(e)
		}
		gotime.Start(cmd.Context(), t, gotime.Options{At: at})
	},
}
