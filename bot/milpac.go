package bot

import (
	"github.com/7cav/discord-sync/cavAPI"
	"github.com/7cav/discord-sync/keycloak"
	"github.com/bwmarrin/discordgo"
	"log"
)

var MilpacCommandName = "milpac"

func MilpacCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        MilpacCommandName,
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

func HandleMilpac(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("Running milpac cmd")

	log.Printf("attempt get kc user for %s\n", i.Member.User.ID)

	kcUser, err := keycloak.KCUserViaDiscordID(i.Member.User.ID)

	if err != nil {
		AskToConnectDiscord(s, i)
		return
	}

	cavUser := cavAPI.GetUserViaKCID(kcUser.ID)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Some milpac",
			Embeds: []*discordgo.MessageEmbed{
				{
					Image: &discordgo.MessageEmbedImage{
						URL: cavUser.UniformUrl,
					},
				},
			},
		},
	})

	if err != nil {
		return
	}
}
