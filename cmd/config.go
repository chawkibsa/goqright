/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup goqright configurations",
	Long: `Configures goqright backend integrations and other related services

It must be used along with other subcommands (e.g.,integration ) and flags (e.g., --add, --remove, --show, --list).`,
	//Run: func(cmd *cobra.Command, args []string) {
	//fmt.Println("config called")

	//},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
