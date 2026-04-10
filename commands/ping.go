package commands

import "github.com/bwmarrin/discordgo"

// Handle processes incoming slash command interactions
func Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	var response string
	switch i.ApplicationCommandData().Name {
	case "ping":
		response = "Pong!"
	case "pong":
		response = "Ping!"
	default:
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
}
