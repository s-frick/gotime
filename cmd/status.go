package cmd

import (
	gotime "github.com/s-frick/go-time-track/pkg"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the current status of GoTime",
	Long:  "Print the current status of GoTime",
	Run: func(cmd *cobra.Command, args []string) {
		gotime.Status(cmd.Context())
	},
}
