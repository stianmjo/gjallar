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
func Run(token string) {

	// Create a new Discord session with the bot token
	dg, err := discordgo.New("Bot " + token)

	// Print errors if session fails
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// Handler to see sent messages in a channel the bot can see
	dg.AddHandler(messageHandler)

	// Open a websocket connection to discord
	err = dg.Open()

	// Print errors if websocket fails
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	// Close the connection when the function returns
	defer dg.Close()

	log.Println("Gjallar is running. Press Ctrl+C to exit.")

	// Create a channel to receive OS signals
	sc := make(chan os.Signal, 1)
	// Listen for interrupt (Ctrl+C)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

// messageHandler is called every time a message is created
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages from the bot itself to prevent loops
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Pass the message to our command handler
	commands.Handle(s, m)
}