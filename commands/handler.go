package commands

import "github.com/bwmarrin/discordgo"

// Handle routes incoming slash commands to their respective handlers
func Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "ping":
		handlePing(s, i)
	case "pong":
		handlePong(s, i)
	case "joke":
		handleJoke(s, i)
	}
}