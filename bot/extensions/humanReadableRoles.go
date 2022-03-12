package extensions

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cache map[string]*discordgo.Role = make(map[string]*discordgo.Role)
var session *discordgo.Session

func BootstrapCavDiscordCache(s *discordgo.Session) {
	guildId := viper.GetString("discord.guild-id")

	session = s

	log.Debug("bootstrapping discord cache")

	// bootstrap cache with roles
	roles, err := s.GuildRoles(guildId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"guild_id": guildId,
		}).Error("could not get roles on discord server")
		return
	}

	for _, role := range roles {
		cache[role.ID] = role
	}
}

func HumanReadableRoles(roleIds ...string) string {
	var res string
	opened := false

	for _, roleId := range roleIds {
		var roleName string
		if _, found := cache[roleId]; !found {
			log.WithFields(log.Fields{
				"role_id": roleId,
			}).Warn("could not find role in cache, refreshing cache")
			BootstrapCavDiscordCache(session)
			roleName = HumanReadableRoles(roleId)
		}

		roleName = cache[roleId].Name

		if opened {
			res += ", "
		}
		opened = true
		res += roleName
	}

	return res
}
