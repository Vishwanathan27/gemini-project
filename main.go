package main

import (
	"log"
	"net/http"

	"github.com/Vishwanathan27/gemini-project/controllers"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func startServer() {
	http.HandleFunc("/sendprompt", controllers.GetResponse)

	// Log the message that the server is starting
	log.Println("Server is starting on port 8080...")

	// Start the server and log any errors
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func startCronJobs() {
	// Create a new cron manager
	c := cron.New()

	// Add the news_letter function to be called at 11:00 AM.
	_, err := c.AddFunc("39 11 * * *", controllers.NewsLetter)
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	// Start the cron scheduler
	c.Start()

	// Log that the cron scheduler has started
	log.Println("Cron scheduler started.")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Start HTTP server in a separate goroutine to prevent blocking
	go startServer()

	// Start the cron jobs
	startCronJobs()

	// Block the main goroutine forever
	select {}
}
