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

	// Retrieve the GUILD_ID for slash command registration (instant propagation)
	guildID := os.Getenv("GUILD_ID")

	// If guild ID is empty then send fatal log message and exit
	if guildID == "" {
		log.Fatal("GUILD_ID not set")
	}

	// Run the bot with the set token and guild ID
	bot.Run(token, guildID)
}