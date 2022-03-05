package extensions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var emptyRoles []string = nil

func RemoveRoles(session *discordgo.Session, member *discordgo.Member, rolesToRemove []string) {
	guildId := viper.GetString("discord.guild-id")

	newRoles := intersect(member.Roles, emptyRoles, rolesToRemove)

	err := session.GuildMemberEdit(guildId, member.User.ID, newRoles)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":  member.User.ID,
			"username": member.User.Username,
			"error":    err,
		}).Errorf("Error removing role from user")
	}
}

func AddRoles(session *discordgo.Session, member *discordgo.Member, rolesToAdd []string) {
	guildId := viper.GetString("discord.guild-id")

	newRoles := intersect(member.Roles, rolesToAdd, emptyRoles)

	err := session.GuildMemberEdit(guildId, member.User.ID, newRoles)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":  member.User.ID,
			"username": member.User.Username,
			"error":    err,
		}).Errorf("Error adding role to user")
	}
}

type CavUserUpdate struct {
	discordUser *discordgo.Member
	nickname    *string
	addRoles    []string
	removeRoles []string
}

func UpdateCavUser(session *discordgo.Session, update *CavUserUpdate) {
	guildId := viper.GetString("discord.guild-id")

	newRoles := intersect(update.discordUser.Roles, update.addRoles, update.removeRoles)

	err := session.GuildMemberEdit(guildId, update.discordUser.User.ID, newRoles)
	if err != nil {
		log.WithFields(log.Fields{
			"user_id":  update.discordUser.User.ID,
			"username": update.discordUser.User.Username,
			"error":    err,
		}).Errorf("Error updating roles on user")
	}
}

// Figure out what roles we actually need to add to the user,taking into account the pre-existing roles they have
// already. I assume that the existing roles will be the longest, so that will be the primary loop. I turn the roles
// we want to add, and the roles we want to remove, into hashmaps. Then loop over the existing roles and check for both
// of them
func intersect(original, toAdd, toRemove []string) []string {
	max := len(original) + len(toAdd) + len(toRemove)
	res := make([]string, 0, max)

	removeMap := make(map[string]struct{}, len(toRemove))
	for _, el := range toRemove {
		removeMap[el] = struct{}{}
	}
	addMap := make(map[string]struct{}, len(toAdd))
	for _, el := range toAdd {
		addMap[el] = struct{}{}
	}

	for i, el := range original {
		fmt.Printf("\nNew Size of res array %d", len(res))
		fmt.Printf("\nChecking item %s", el)

		if _, ok := removeMap[el]; ok {
			fmt.Printf("\nNeed to remove %s", el)
			continue
		}

		fmt.Printf("\nExisting item, keeping %s", el)
		res = append(res, el)

		if _, ok := addMap[el]; !ok && i < len(toAdd) {
			fmt.Printf("\nNeed to add %s", toAdd[i])
			res = append(res, toAdd[i])
		}
	}

	return res
}

// func (session, user, rolesToAdd, rolesToRemove) -> err
// 1. get current roles for user
// 2. intersect existing roles from those to add/remove
// 3. batch update user
