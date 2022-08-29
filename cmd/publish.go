/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/utils"
	"github.com/wahidx/kafkax/xkafka"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish message to a topic",

	Run: func(cmd *cobra.Command, args []string) {
		topic, err := cmd.Flags().GetString("topic")
		if err != nil || len(topic) == 0 {
			fmt.Println("Invalid topic")
			return
		}

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println("Invalid key")
			return
		}

		headers, err := cmd.Flags().GetStringArray("header")
		if err != nil {
			fmt.Println("Invalid header")
			return
		}
		header := utils.ReadHeaders(headers)

		fmt.Println("Press Ctrl+C to exit")

		reader := bufio.NewReader(os.Stdin)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println("Failed to read\n", err)
				return
			}

			xkafka.Publish(topic, string(line), key, header)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.PersistentFlags().StringP("topic", "t", "", "")
	publishCmd.PersistentFlags().StringP("key", "k", "", "")
	publishCmd.PersistentFlags().StringArrayP("header", "H", []string{}, "syntax: key1=val1")
}
