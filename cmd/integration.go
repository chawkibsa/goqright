/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/chawkibsa/goqright/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// integrationCmd represents the integration command
var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "Configure integrations APIs",
	Long:  `This command is setup the selected integrations (products supported by goqright) along with their correspondent API keys for automation usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("integration called")
		add, _ := cmd.Flags().GetBool("add")
		if add {
			data.CreateIntegrationsTable()
			addIntegration()
		}

		remove, _ := cmd.Flags().GetBool("remove")
		if remove {
			// remove integration
			removeIntegration()
		}
		show, _ := cmd.Flags().GetBool("show")
		if show {
			showIntegration()
		}
		list, _ := cmd.Flags().GetBool("list")
		if list {
			listIntegrations()
		}

	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	configCmd.AddCommand(integrationCmd)
	integrationCmd.Flags().BoolP("add", "a", false, "Add new integration")
	integrationCmd.Flags().BoolP("remove", "r", false, "Remove integration")
	integrationCmd.Flags().BoolP("update", "u", false, "Update integration")
	integrationCmd.Flags().BoolP("show", "s", false, "Show integration")
	integrationCmd.Flags().BoolP("list", "l", false, "List integrations")
	// Here you will define your flags and configuration settings.
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result

}

func promptGetSelect(pc promptContent) string {
	items := []string{"IBM Qradar", "MISP"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.label,
			Items: items,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func addIntegration() {

	integrationPromptContent := promptContent{
		"Integration is required",
		"Integration",
	}

	integration_type := promptGetSelect(integrationPromptContent)

	namePromptContent := promptContent{
		"Integration Name is required",
		fmt.Sprintf("Name of integration \"%s\"", integration_type),
	}

	name := promptGetInput(namePromptContent)

	apiKeyPromptContent := promptContent{
		"API Key is required",
		fmt.Sprintf("API Key of %s", name),
	}

	apiKey := promptGetInput(apiKeyPromptContent)
	data.InsertIntegration(integration_type, name, apiKey)
}

func removeIntegration() {
	// remove the integration based on its id
	idPromptContent := promptContent{
		"ID is required",
		"ID",
	}
	idStr := promptGetInput(idPromptContent)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid ID. List current integrations with \"-l\".")
		os.Exit(1)
	}

	data.RemoveIntegration(id)

}

func showIntegration() {
	// show all saved integrations
	integrations := data.GetIntegrations()

	data.PrintIntegrations(integrations)

}

func listIntegrations() {
	// list all available integrations (IBM Qradar and MISP currently)
	data.CreateSupportedIntegrationsTable()
	data.InsertSupportedIntegration("IBM QRadar")
	data.InsertSupportedIntegration("MISP")

	integrations := data.GetSupportedIntegrations()

	data.PrintSupportedIntegrations(integrations)

}
