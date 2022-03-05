package extensions

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var emptyRoles []string = nil

func RemoveRoles(session *discordgo.Session, member *discordgo.Member, rolesToRemove []string) {
	UpdateCavUser(session, &CavUserUpdate{
		DiscordUser: member,
		Nickname:    &member.Nick,
		RemoveRoles: rolesToRemove,
	})
}

func AddRoles(session *discordgo.Session, member *discordgo.Member, rolesToAdd []string) {
	UpdateCavUser(session, &CavUserUpdate{
		DiscordUser: member,
		Nickname:    &member.Nick,
		AddRoles:    rolesToAdd,
	})
}

type CavUserUpdate struct {
	DiscordUser *discordgo.Member
	Nickname    *string
	AddRoles    []string
	RemoveRoles []string
}

func UpdateCavUser(session *discordgo.Session, update *CavUserUpdate) {
	guildId := viper.GetString("discord.guild-id")

	newRoles, delta := intersect(update.DiscordUser.Roles, update.AddRoles, update.RemoveRoles)

	if delta == 0 {
		log.WithFields(log.Fields{
			"nickname": update.DiscordUser.Nick,
		}).Info("no changes to roles, skipping update")
		return
	}

	err := session.GuildMemberEdit(guildId, update.DiscordUser.User.ID, newRoles)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":  update.DiscordUser.User.ID,
			"username": update.DiscordUser.User.Username,
			"error":    err,
		}).Errorf("Error updating roles on user")
	}
}

// Figure out what roles we actually need to add to the user,taking into account the pre-existing roles they have
// already. I assume that the existing roles will be the longest, so that will be the primary loop. I turn the roles
// we want to add, and the roles we want to remove, into hashmaps. Then loop over the existing roles and check for both
// of them
func intersect(original, toAdd, toRemove []string) ([]string, int) {
	max := len(original) + len(toAdd) + len(toRemove)
	res := make([]string, 0, max)
	delta := 0

	removeMap := make(map[string]struct{}, len(toRemove))
	for _, el := range toRemove {
		removeMap[el] = struct{}{}
	}
	addMap := make(map[string]struct{}, len(toAdd))
	for _, el := range toAdd {
		addMap[el] = struct{}{}
	}

	for i, el := range original {

		if _, found := removeMap[el]; found {
			if _, found := addMap[el]; !found {
				delta++
				continue
			}
		}

		delta++
		res = append(res, el)

		if _, found := addMap[el]; !found && i < len(toAdd) {
			delta++
			res = append(res, toAdd[i])
		}
	}

	return res, delta
}
