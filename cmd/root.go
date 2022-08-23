/*
Copyright © 2022 wahidx@wahidx93@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafkax",
	Short: "CLI app to interact with a kafka cluster",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey there")
		val, err := cmd.Flags().GetString("topic")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Topic: ", val)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kafkax.yaml)")

	rootCmd.Flags().StringP("topic", "t", "", "Topic name")
	rootCmd.Flags().StringP("broker", "b", "http://localhost:9092", "Broker URL")

}
