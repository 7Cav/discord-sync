package hello

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var CommandName = "hello"

func Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        CommandName,
		Description: "hello world",
	}
}

func Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("Running hello cmd")

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello world",
		},
	})
	if err != nil {
		return
	}
}
