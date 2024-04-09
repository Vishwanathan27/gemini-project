// main.go
package main

import (
	"log"
	"net/http"

	"github.com/Vishwanathan27/gemini-project/controllers"
)

func main() {
	http.HandleFunc("/sendprompt", controllers.GetResponse)

	// Log the message that the server is starting
	log.Println("Server is starting on port 8080...")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
