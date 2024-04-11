package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// NewsLetterService is responsible for fetching and filtering stories
type NewsLetterService struct {
	InterestedTopics []string
}

// NewNewsLetterService creates a new instance of NewsLetterService with predefined topics
func NewNewsLetterService() *NewsLetterService {
	return &NewsLetterService{
		InterestedTopics: []string{"AI", "node js", "gemini", "openai", "startup", "gadget", "apple", "ipad", "iphone", "mac"},
	}
}

// FetchStories fetches and filters Hacker News stories based on interested topics
func (n *NewsLetterService) FetchStories() ([]Story, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty")
	if err != nil {
		return nil, fmt.Errorf("error fetching new stories: %v", err)
	}
	defer resp.Body.Close()

	var storyIDs []int
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if err := json.Unmarshal(body, &storyIDs); err != nil {
		return nil, fmt.Errorf("error unmarshaling story IDs: %v", err)
	}

	var filteredStories []Story
	for _, id := range storyIDs {
		storyURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", id)
		resp, err := http.Get(storyURL)
		if err != nil {
			log.Printf("Error fetching story details for ID %d: %v", id, err)
			continue
		}

		var storyDetails map[string]interface{}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error reading response body for ID %d: %v", id, err)
			continue
		}

		if err := json.Unmarshal(body, &storyDetails); err != nil {
			log.Printf("Error unmarshaling story details for ID %d: %v", id, err)
			continue
		}

		title, titleOk := storyDetails["title"].(string)
		url, urlOk := storyDetails["url"].(string)
		if titleOk && urlOk {
			titleWords := strings.Fields(strings.ToLower(title)) // Split title into words
			for _, topic := range n.InterestedTopics {
				topicLower := strings.ToLower(topic)
				matchFound := false

				for _, word := range titleWords {
					if word == topicLower {
						matchFound = true
						break // Break the word loop if a match is found
					}
				}

				if matchFound {
					filteredStories = append(filteredStories, Story{Title: title, URL: url})
					break // Break the topic loop if a match is found
				}
			}

		}
	}

	return filteredStories, nil
}
