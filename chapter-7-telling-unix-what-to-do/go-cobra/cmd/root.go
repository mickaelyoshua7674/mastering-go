/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A sample Cobra project",
	Long: `A sample Cobra project for Mastering Go, 4th edition.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
        fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
    rootCmd.PersistentFlags().StringP("directory", "d", "/tmp", "Path")
    rootCmd.PersistentFlags().Uint("depth", 2, "Depth of search")
    viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))
    viper.BindPFlag("depth", rootCmd.PersistentFlags().Lookup("depth"))

    twoCmd.Flags().StringP("username", "u", "Mike", "Username")
}


