package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/himalayan-institute/zoom-lib-golang"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
		userID	  = os.Getenv("USER_ID")
	)

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret

	var createMeetingOpts zoom.CreateMeetingOptions
	createMeetingOpts.HostID = userID
	meeting, err := zoom.CreateMeeting(createMeetingOpts)
	if err != nil {
		log.Printf("Got err: %+v\n", err)
	}
	log.Printf("Get Meeting Info: %+v\n", meeting)
}