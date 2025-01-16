/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// threeCmd represents the three command
var threeCmd = &cobra.Command{
	Use:   "three",
    Aliases: []string{"cmd3"},
	Short: "Command three",
	Long: `Command three demo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("three called")
	},
}

func init() {
	rootCmd.AddCommand(threeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// threeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// threeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
