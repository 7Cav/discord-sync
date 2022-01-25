package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var HelloCommandName = "hello"

func HelloCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        HelloCommandName,
		Description: "hello world",
	}
}

func HandleHello(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
