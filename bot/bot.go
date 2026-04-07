package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gjallar/commands"
)

// RUn initializes and starts the Discord bot called from main.go
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

	// Create a channel to recieve OS signals
	sc := make(chan os.Signal, 1)
	// Listen for interrupt (Ctrl+C)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore messages from the bot itself to prevent loops
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Pass the message to our command handler
		commands.Handle(s, m)
	}
}