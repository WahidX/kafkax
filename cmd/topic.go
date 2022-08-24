/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/xkafka"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "",

	Run: func(cmd *cobra.Command, args []string) {
		xkafka.ListTopics()
	},
}

func init() {
	listCmd.AddCommand(topicCmd)
	createCmd.AddCommand(topicCmd)
}
