package services

import (
	"context" // Import the context package for managing request lifecycles
	"fmt"     // Import the fmt package for formatting strings and errors
	"log"     // Import the log package for logging messages
	"os"      // Import the os package to access environment variables

	"github.com/google/generative-ai-go/genai" // Import the genai package for interacting with generative AI models
	"google.golang.org/api/option"             // Import the option package for configuring API client options
)

// The default model name to use if another is not specified
const generativeModelName = "gemini-pro"

// GenAIClient wraps the genai.Client providing a simplified interface
type GenAIClient struct {
	client *genai.Client // The underlying generative AI client
}

// NewGenAIClient creates a new GenAIClient instance
func NewGenAIClient() (*GenAIClient, error) {
	ctx := context.Background()           // Create a new context
	apiKey := os.Getenv("GEMINI_API_KEY") // Retrieve the API key from the environment variables

	// Initialize a new generative AI client with the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		// If there is an error during client creation, return it
		return nil, err
	}

	// Return a new GenAIClient instance
	return &GenAIClient{client: client}, nil
}

// GenerateContent generates content based on a prompt using a specific model
func (g *GenAIClient) GenerateContent(prompt, modelName string) (string, error) {
	ctx := context.Background() // Create a new context
	// Use a specified model name, defaulting to generativeModelName if not provided
	if modelName == "" {
		modelName = generativeModelName
	}
	model := g.client.GenerativeModel(modelName) // Get the model from the client
	// Generate content based on the prompt
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		// If there's an error generating content, wrap and return it
		return "", fmt.Errorf("error generating content: %w", err)
	}

	fmt.Printf("Gemini Response Structure: %#v\n", resp)
	return "Successfully succeeded", nil // Return the generated text
}

// Close shuts down the GenAIClient's underlying client connection
func (g *GenAIClient) Close() {
	if g.client != nil {
		// If the client exists, attempt to close it and log any errors
		if err := g.client.Close(); err != nil {
			log.Printf("Error closing GenAI client: %v", err)
		}
	}
}
