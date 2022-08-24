/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list topics or brokers",

	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("list called")
	// },
}

func init() {
	rootCmd.AddCommand(listCmd)

}
