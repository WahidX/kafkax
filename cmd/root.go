/*
Copyright © 2022 wahidx@wahidx93@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafkax",
	Short: "CLI app to interact with a kafka cluster",
	// Long: ``,

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("broker", "b", "", "Kafka broker (comma separated)")
}
