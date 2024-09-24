package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/christimahu/gaikeep/internal/gaikeep/api"
	"github.com/christimahu/gaikeep/internal/gaikeep/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gaikeep",
	Short: "GaiKeep is a companion to LLM tools",
	Long: `GaiKeep aggregates requests, listens for cues, and stores important responses
in a local file or database.`,
}

// TestGPTCmd represents a subcommand to test interactions with ChatGPT
var testGPTCmd = &cobra.Command{
	Use:   "testgpt",
	Short: "Test ChatGPT interaction",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("OPENAI_API_KEY must be set in the environment variables")
		}

		response, err := api.FetchResponse(apiKey, "Hello, how can I help you today?")
		if err != nil {
			log.Fatalf("Failed to get response from OpenAI: %v", err)
		}

		err = storage.SaveResponse(response)
		if err != nil {
			log.Fatalf("Failed to save response: %v", err)
		}

		fmt.Println("Response from ChatGPT saved successfully.")
	},
}

func Execute() {
	rootCmd.AddCommand(testGPTCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
