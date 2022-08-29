/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahidx/kafkax/utils"
	"github.com/wahidx/kafkax/xkafka"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume messages of a topic",

	Run: func(cmd *cobra.Command, args []string) {
		broker, err := cmd.Flags().GetString("broker")
		if err != nil {
			fmt.Println("Invalid broker")
			return
		} else if len(broker) != 0 {
			utils.SetBroker(broker)
		}

		topic, err := cmd.Flags().GetString("topic")
		if err != nil || len(topic) == 0 {
			fmt.Println("Invalid topic")
			return
		}

		groupID, err := cmd.Flags().GetString("groupID")
		if err != nil {
			fmt.Println("Invalid groupID")
			return
		}

		isJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println("Invalid JSON option")
			return
		}

		xkafka.Consume(topic, groupID, isJSON)
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	consumeCmd.PersistentFlags().StringP("topic", "t", "", "")
	consumeCmd.PersistentFlags().StringP("groupID", "g", "", "")
	consumeCmd.PersistentFlags().BoolP("json", "j", false, "print messages in JSON")
}
