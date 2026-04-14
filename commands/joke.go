package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// jokeResponse maps the JSON fields returned by JokeAPI
// The `json:"..."` tags tell Go which JSON field maps to which struct field
type jokeResponse struct {
	Type     string `json:"type"`     // "single" or "twopart"
	Joke     string `json:"joke"`     // only populated when Type is "single"
	Setup    string `json:"setup"`    // only populated when Type is "twopart"
	Delivery string `json:"delivery"` // the punchline, only when Type is "twopart"
	Error    bool   `json:"error"`    // true if the API returned an error
}

// handleJoke fetches a random safe joke from JokeAPI and responds to the interaction
func handleJoke(s *discordgo.Session, i *discordgo.InteractionCreate) {

	// Make a GET request to JokeAPI
	// safe-mode filters out nsfw, religious, political, racist, sexist and explicit jokes
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any?safe-mode&lang=en")
	if err != nil {
		// If the HTTP request itself failed (e.g. no network), tell the user
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to fetch a joke.",
			},
		})
		return
	}

	// Ensure the response body is closed after we are done reading it
	defer resp.Body.Close()

	// Decode the JSON response body into our jokeResponse struct
	var joke jokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil || joke.Error {
		// If JSON parsing failed, or the API itself returned an error flag, tell the user
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Could not parse the joke.",
			},
		})
		return
	}

	// Build the message content depending on the joke type
	var content string
	if joke.Type == "single" {
		// Single jokes are just one line
		content = joke.Joke
	} else {
		// Two-part jokes have a setup and a punchline
		// The punchline is wrapped in || so Discord hides it as a spoiler
		// Users click it to reveal the punchline
		content = fmt.Sprintf("%s\n||%s||", joke.Setup, joke.Delivery)
	}

	// Send the joke as a response to the slash command interaction
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}