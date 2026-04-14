package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gjallar/commands"
)

// Run initializes and starts the Discord bot called from main.go
func Run(token string, guildID string) {

	// Create a new Discord session with the bot token
	dg, err := discordgo.New("Bot " + token)

	// Print errors if session fails
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// Handler for slash command interactions
	dg.AddHandler(interactionHandler)

	// Open a websocket connection to discord
	err = dg.Open()

	// Print errors if websocket fails
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	// Close the connection when the function returns
	defer dg.Close()

	// Register slash commands
	slashCommands := []*discordgo.ApplicationCommand{
		{Name: "ping", Description: "Replies with Pong!"},
		{Name: "pong", Description: "Replies with Ping!"},
		{Name: "joke", Description: "Tells a random safe joke"},
	}
	for _, cmd := range slashCommands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, cmd)
		if err != nil {
			log.Fatalf("Error creating slash command %s: %v", cmd.Name, err)
		}
	}

	log.Println("Gjallar is running. Press Ctrl+C to exit.")

	// Create a channel to receive OS signals
	sc := make(chan os.Signal, 1)
	// Listen for interrupt (Ctrl+C)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

// interactionHandler is called every time a slash command interaction is created
func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.Handle(s, i)
}