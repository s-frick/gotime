package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GoTime",
	Long:  "Print the version number of GoTime",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoTime v0.1.0")
	},
}
