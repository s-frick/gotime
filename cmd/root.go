package cmd

import (
	"context"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	at string
)

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	startCmd.Flags().StringVar(&at, "at", "", "Start the frame at a specific time. e.g. \"09:25\"")
	stopCmd.Flags().StringVar(&at, "at", "", "Stop the frame at a specific time. e.g. \"09:25\"")

	// INFO: INITIALIZE YOUR FLAGS
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}

func initConfig() {
	// if cfgFile != "" {
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	//
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".time-track")
	// }
	//
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Can't read config: ", err)
	// 	os.Exit(1)
	// }
}

var rootCmd = &cobra.Command{
	Use:   "gotime",
	Short: "GoTime is a small flexible time tracker",
	Long:  `GoTime is a small flexible time tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func ExecuteContext() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gotimeDir := fmt.Sprintf("%s/.gotime", home)
	ctx := context.WithValue(context.Background(), "gotimeDir", gotimeDir)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
