/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// slackCmd represents the slack command
var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slack called")
	},
}

func init() {
	rootCmd.AddCommand(slackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	slackCmd.PersistentFlags().String("title", "", "A title for the message")
	slackCmd.PersistentFlags().String("title-url", "", "An url for the title")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// slackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
