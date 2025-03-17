/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sourceipCmd represents the sourceip command
var sourceipCmd = &cobra.Command{
	Use:   "sourceip",
	Short: "'Source IP' custom property",
	Long: `This command defines 'Source IP' as a custom property. As a Qradar administrator you will benefit of all of 
out of the box automated tasks associated to this property. It must be associated with appropriate flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sourceip called")
	},
}

func init() {
	propertyCmd.AddCommand(sourceipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
