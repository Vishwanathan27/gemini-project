package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Vishwanathan27/gemini-project/services"
)

// NewsLetter is called by a cron job to fetch and filter news stories
func NewsLetter() {
	fmt.Println("Starting to fetch Hacker News stories at", time.Now())

	newsletterService := services.NewNewsLetterService()
	filteredStories, err := newsletterService.FetchStories()
	if err != nil {
		log.Printf("Error fetching stories: %v", err)
		return
	}

	filteredJSON, err := json.MarshalIndent(filteredStories, "", "  ")
	if err != nil {
		log.Printf("Error marshaling filtered stories: %v", err)
		return
	}

	fmt.Println("Filtered Stories in JSON format:")
	fmt.Println(string(filteredJSON))
}
