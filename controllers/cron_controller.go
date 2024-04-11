package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Vishwanathan27/gemini-project/services" // Ensure this path matches your project structure
)

func NewsLetter() {
	fmt.Printf("Fetching news stories for summarization at %v\n", time.Now())

	// Assuming you have a list of filteredStories from your previous steps
	newsletterService := services.NewNewsLetterService()
	filteredStories, err := newsletterService.FetchStories()
	if err != nil {
		log.Printf("Error fetching stories: %v\n", err)
		return
	}

	// Prepare the prompt from titles for summarization
	var titles []string
	for _, story := range filteredStories {
		// Format each title with its URL in parentheses and append to titles slice
		formattedTitle := fmt.Sprintf("%s (%s)", story.Title, story.URL)
		titles = append(titles, formattedTitle)
	}

	prompt := "Summarize these news stories just like how Jarvis would summarize all the updates to tony stark, make it sound conversational, make it a little humourous, this will be sent as an email: " + strings.Join(titles, ", ") + "."

	// Use the existing GenAIClient to generate content (i.e., summarize the prompt)
	client, err := services.NewGenAIClient()
	if err != nil {
		log.Printf("Failed to create GenAI client: %v\n", err)
		return
	}
	defer client.Close()

	summary, err := client.GenerateContent(prompt, "")
	if err != nil {
		log.Printf("Failed to summarize news stories: %v\n", err)
		return
	}

	fmt.Println("Summary of the latest news:")
	fmt.Println(summary)

	if summary != "" {
		// Send the summary content as an email
		mailSender := services.NewMailSender() // Make sure NewMailSender is properly defined in your services package.
		subject := "Your Daily News Summary"
		to := "vish1nathan27@gmail.com" // The recipient's email address

		// Call the SendEmail method to send the summary as an email
		err := mailSender.SendEmail([]string{to}, subject, summary)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
			return
		}

		fmt.Println("Email sent with the summary.")
	} else {
		fmt.Println("No summary to send.")
	}

}
