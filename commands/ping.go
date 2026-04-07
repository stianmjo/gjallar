package commands

import "github.com/bwmarrin/discordgo"

// Handle processes incoming message and routes them to commands
func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check the message content and respond accordingly
	switch m.Content {
	case "!ping":

		// Send "Pong!" back to the same channel
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "!pong":

		// Send "Pong!" back to the same channel
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}