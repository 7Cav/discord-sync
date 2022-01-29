package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

type Bot struct {
	conn *discordgo.Session
}

func New(token string) (Bot, error) {
	conn, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Could not start discord bot: %v", err)
	}

	return Bot{
		conn: conn,
	}, nil
}

func (b Bot) Start(appId string, guildId string) {

	conn := b.conn

	conn.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	err := conn.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	commands := []*discordgo.ApplicationCommand{
		MilpacCommand(),
		HelloCommand(),
		SyncCommand(),
	}

	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		MilpacCommandName: HandleMilpac,
		HelloCommandName:  HandleHello,
		SyncCommandName:   HandleSync,
	}

	_, err = conn.ApplicationCommandBulkOverwrite(appId, guildId, commands)
	if err != nil {
		log.Panicf("Cannot create commands: %v", err)
	}

	conn.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	defer conn.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")
}

func AskToConnectDiscord(s *discordgo.Session, i *discordgo.InteractionCreate) {

	mention := fmt.Sprintf("<@%s>", i.Member.User.ID)
	kcUrl := fmt.Sprintf("%s/auth/realms/%s/account/identity", viper.GetString("keycloak.host"), viper.GetString("keycloak.realm"))
	reply := fmt.Sprintf(`Hey %s, we don't have a valid discord connection for you. Please go to %s and click 'add' to connect your discord account. Then try again!`, mention, kcUrl)

	// this was called because no KC user was found for the given discord ID
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: reply,
		},
	})

	if err != nil {
		return
	}
}
