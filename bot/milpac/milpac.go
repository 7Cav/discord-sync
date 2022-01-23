package milpac

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var CommandName = "milpac"

func Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        CommandName,
		Description: "Fetch basic MILPACs data for a given user",
		//Options: []*discordgo.ApplicationCommandOption{
		//	{
		//		Type:        discordgo.ApplicationCommandOptionString,
		//		Name:        "username",
		//		Description: "Member Username",
		//		Required:    false,
		//	},
		//},
	}
}

func Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("Running milpac cmd")

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Some milpac",
			Embeds: []*discordgo.MessageEmbed{
				{
					Image: &discordgo.MessageEmbedImage{
						URL: "https://7cav.us/data/roster_uniforms/0/1.jpg?1638502104",
					},
				},
			},
		},
	})
	if err != nil {
		return
	}
}
