/*
Copyright © 2022 7Cav.us

*/
package cmd

import (
	"github.com/7cav/discord-sync/bot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start discord bot",
	Run: func(cmd *cobra.Command, args []string) {
		appId := viper.GetString("discord.app-id")
		guildId := viper.GetString("discord.guild-id")
		token := viper.GetString("discord.token")

		if appId == "" || guildId == "" || token == "" {
			log.Fatalf("check config, something is empty")
			return
		}

		bot, err := bot.New(token)

		if err != nil {
			log.Fatalf("Cannot start bot: %v", err)
		}

		bot.Start(appId, guildId)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
