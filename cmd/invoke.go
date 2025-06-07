package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

var invokeCmdInput string

var invokeToolCmd = &cobra.Command{
	Use:   "invoke <name>",
	Short: "Invoke a tool",
	Long:  "Invokes a tool supplied by a registered MCP server",
	Args:  cobra.ExactArgs(1),
	RunE:  runInvokeTool,
}

func init() {
	invokeToolCmd.Flags().StringVar(&invokeCmdInput, "input", "{}", "valid JSON payload")
	rootCmd.AddCommand(invokeToolCmd)
}

func runInvokeTool(cmd *cobra.Command, args []string) error {
	var input map[string]any
	if err := json.Unmarshal([]byte(invokeCmdInput), &input); err != nil {
		return fmt.Errorf("invalid input: %w", err)
	}

	result, err := apiClient.InvokeTool(args[0], input)
	if err != nil {
		return fmt.Errorf("failed to invoke tool: %w", err)
	}

	if result.IsError {
		fmt.Println("The tool returned an error:")
		for k, v := range result.Meta {
			fmt.Printf("%s: %v\n", k, v)
		}
	} else {
		fmt.Println("Response from tool:")
	}

	// result Content needs to be printed regardless of whether the tool returned an error or not
	// because it may contain useful information
	fmt.Println()
	for _, c := range result.Content {
		cType, ok := c["type"]
		if !ok {
			return fmt.Errorf("content item does not have a 'type' field: %v", c)
		}
		switch cType {
		case "text":
			textContent, ok := c["text"]
			if !ok {
				return fmt.Errorf("text content item does not have a 'text' field: %v", c)
			}
			fmt.Println(textContent)
		case "image":
			// todo
		case "audio":
			// todo
		}
	}

	return nil
}
