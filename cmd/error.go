/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kalvinov.com/gocli/ops"
)

// errorCmd represents the error command
var errorCmd = &cobra.Command{
	Use:   "error",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetString("title")
		titleUrl, _ := cmd.Flags().GetString("title-url")
		err := ops.SendSlackMessage(viper.GetString("slackUrl"), "#ff0000", title, titleUrl, strings.Join(args, " "))
		if err != nil {
			return err
		}
		return nil
	},
	SilenceErrors: true,
}

func init() {
	slackCmd.AddCommand(errorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
