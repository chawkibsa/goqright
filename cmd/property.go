/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// propertyCmd represents the property command
var propertyCmd = &cobra.Command{
	Use:   "property",
	Short: "Define one of our supported custom properties",
	Long:  `This command will allow IBM Administrators to define a custom property from our supported custom properties in order to assign to it out of the box automated tasks that improve triage/incident response process.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("property called")
	},
}

func init() {
	generateCmd.AddCommand(propertyCmd)
}
