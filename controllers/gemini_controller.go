package controllers

import (
	"encoding/json" // Import the json package for encoding and decoding JSON
	"io/ioutil"
	"log"      // Import the log package for logging messages
	"net/http" // Import the http package to work with HTTP requests and responses

	"github.com/Vishwanathan27/gemini-project/services" // Import the services package where GenAIClient is defined
)

// GetResponse handles an HTTP request to generate content based on a prompt
func GetResponse(w http.ResponseWriter, r *http.Request) {
	// Read the entire request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Log and return an error if the body cannot be read
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	// Attempt to unmarshal the JSON body into a map
	var payload map[string]string
	err = json.Unmarshal(body, &payload)
	if err != nil {
		// Log and return an error if the JSON cannot be parsed
		log.Printf("Error parsing JSON payload: %v", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Extract the 'prompt' field from the payload
	prompt, ok := payload["prompt"]
	if !ok {
		// Return an error if the 'prompt' field is missing
		log.Printf("Missing 'prompt' field in payload")
		http.Error(w, "Missing 'prompt' field", http.StatusBadRequest)
		return
	}

	// Initialize a new GenAIClient
	client, err := services.NewGenAIClient()
	if err != nil {
		// Log and return an error if the client cannot be created
		log.Printf("Failed to create GenAI client: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close() // Ensure the client is closed after processing

	// Generate content based on the provided prompt
	content, err := client.GenerateContent(prompt, "")
	if err != nil {
		// Log and return an error if content generation fails
		log.Printf("Failed to generate content: %v", err)
		http.Error(w, "Failed to generate content", http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the generated content as the response
	// Marshal the content into JSON
	jsonContent, err := json.Marshal(map[string]string{"response": content})
	if err != nil {
		// If JSON marshaling fails, log the error and send a server error response
		log.Printf("Failed to marshal content to JSON: %v", err)
		http.Error(w, "Failed to marshal content", http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the marshaled JSON content to the response writer
	w.Write(jsonContent)
}
