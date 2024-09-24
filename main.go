package main

import (
	"log"
	"os"

	"github.com/christimahu/gaikeep/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file to retrieve environment variables, if available
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Ensure the OPENAI_API_KEY environment variable is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set. Please check your .env file or set the environment variable.")
	}

	// Execute the root command of the application
	cmd.Execute() // Adjusted to not expect a return value since the function handles its own errors
}
