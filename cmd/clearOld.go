/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/7cav/discord-sync/bot"
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"
)

// clearOldCmd represents the clearOld command
var clearOldCmd = &cobra.Command{
	Use:   "clearOld",
	Short: "Clear active/ret/disch/eloa roles from all users",
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

		bot.SpecialClearOldUsers(guildId)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
