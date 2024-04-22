/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kalvinov.com/gocli/ops"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate env variables",
	Long:  `Validate that required env variables exist and comply with template`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ops.ValidateEnvs(viper.GetStringSlice("validate"))
		if err != nil {
			return err
		}
		return nil
	},
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
