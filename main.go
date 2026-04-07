package main

import (
	"log"
	"os"

	"gjallar/bot"	// Importing bot package
)

func main() {

	// Retrieve and store the DISCORD_TOKEN from set .env
	token := os.Getenv("DISCORD_TOKEN")
	
	// If token empty then send fatal log message and exit
	if token == "" {
		log.Fatal("DISCORD_TOKEN not set")
	}

	// Run the bot with the set token
	bot.Run(token)
}