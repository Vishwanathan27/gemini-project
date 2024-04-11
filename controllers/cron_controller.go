package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// NewsLetter fetches new stories from the Hacker News API and prints them.
func NewsLetter() {
	fmt.Println("Starting to fetch Hacker News stories at", time.Now())

	// First, get the list of new story IDs.
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty")
	if err != nil {
		log.Printf("Error fetching new stories: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var storyIDs []int
	if err := json.NewDecoder(resp.Body).Decode(&storyIDs); err != nil {
		log.Printf("Error decoding story IDs: %v\n", err)
		return
	}

	// Limit the number of stories to fetch details for, if necessary.
	const maxStories = 10 // Fetch details for 10 stories only for example purposes.
	if len(storyIDs) > maxStories {
		storyIDs = storyIDs[:maxStories]
	}

	// Fetch details for each story ID and accumulate the results.
	var stories []interface{}
	for _, id := range storyIDs {
		storyURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
		resp, err := http.Get(storyURL)
		if err != nil {
			log.Printf("Error fetching story details for ID %d: %v\n", id, err)
			continue
		}
		defer resp.Body.Close()

		var storyDetails interface{}
		if err := json.NewDecoder(resp.Body).Decode(&storyDetails); err != nil {
			log.Printf("Error decoding story details for ID %d: %v\n", id, err)
			continue
		}

		stories = append(stories, storyDetails)
	}

	// Print out all the story details.
	storiesJSON, err := json.MarshalIndent(stories, "", "  ")
	if err != nil {
		log.Printf("Error marshalling stories to JSON: %v\n", err)
		return
	}
	fmt.Printf("Fetched Stories: %s\n", storiesJSON)
}
