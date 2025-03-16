/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// integrationCmd represents the integration command
var integrationCmd = &cobra.Command{
	Use:   "integration",
	Short: "Configure integrations APIs",
	Long:  `This command is setup the selected integrations (products supported by goqright) along with their correspondent API keys for automation usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("integration called")
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	configCmd.AddCommand(integrationCmd)

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

	fmt.Printf("Input: %s\n", result)
	return result

}

func promptGetSelect(pc promptContent) string {
	items := []string{"IBM Qradar", "MISP"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.label,
			Items: items,
			//AddLabel: "Add new",
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

	fmt.Printf("Input: %s\n", result)
	return result
}

func addIntegration() {

	integrationPromptContent := promptContent{
		"Integration is required",
		"Integration",
	}

	integration := promptGetSelect(integrationPromptContent)

	apiKeyPromptContent := promptContent{
		"API Key is required",
		fmt.Sprintf("API Key of %s", integration),
	}

	apiKey := promptGetInput(apiKeyPromptContent)
	fmt.Sprintf(apiKey)
}
