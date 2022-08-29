/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/xkafka"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish message to a topic",

	Run: func(cmd *cobra.Command, args []string) {
		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			fmt.Println("Invalid topic name")
			return
		}

		message, err := cmd.Flags().GetString("message")
		if err != nil {
			fmt.Println("Invalid message")
			return
		}

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println("Invalid key")
			return
		}

		xkafka.Publish(topic, message, key)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.PersistentFlags().StringP("topic", "t", "", "")
	publishCmd.PersistentFlags().StringP("message", "m", "", "")
	publishCmd.PersistentFlags().StringP("key", "k", "", "")
}
