package bot

import (
	"github.com/7cav/discord-sync/bot/hello"
	"github.com/7cav/discord-sync/bot/milpac"
	"github.com/bwmarrin/discordgo"
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
		milpac.Command(),
		hello.Command(),
	}

	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		milpac.CommandName: milpac.Handle,
		hello.CommandName:  hello.Handle,
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
