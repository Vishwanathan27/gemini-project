package controllers

import (
	"fmt"
	"time"
)

// NewsLetter is a function meant to be called by a cron job
func NewsLetter() {
	// Your newsletter logic here
	fmt.Println("Newsletter processed at", time.Now())

	
}
