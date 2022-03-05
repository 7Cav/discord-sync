package bot

import (
	"fmt"
	"github.com/7cav/api/proto"
	"github.com/7cav/discord-sync/bot/extensions"
	"github.com/7cav/discord-sync/cavAPI"
	"github.com/7cav/discord-sync/cavDiscord"
	"github.com/7cav/discord-sync/keycloak"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var SyncCommandName = "sync"

func SyncCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        SyncCommandName,
		Description: "Sync account to 7Cav.us",
	}
}

type cavDcPair struct {
	cavProfile *proto.Profile
	dcMember   *discordgo.Member
}

func HandleSync(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Infof("Sync command")

	log.Infof("Attempt get kc user for %s\n", i.Member.User.ID)

	kcUser, err := keycloak.KCUserViaDiscordID(i.Member.User.ID)

	if err != nil {
		AskToConnectDiscord(s, i)
		return
	}

	cavUser := cavAPI.GetUserViaKCID(kcUser.ID)

	var currentUser = &cavDcPair{
		cavProfile: cavUser,
		dcMember:   i.Member,
	}
	var rolesToAdd []string
	var rolesToRemove []string
	var newNick string

	// get new Nickname item (doesn't matter if same - in future Guild Edit will drop the payload anyway)
	newNick = cavDiscord.GenerateCavNickName(currentUser.cavProfile)

	// get all current rank roles (to remove)
	rankRoles := currentRankRoles(currentUser)
	rolesToRemove = append(rolesToRemove, rankRoles...)

	// get correct rank role (to add)
	rankRole := correctRankRole(currentUser)
	rolesToAdd = append(rolesToAdd, rankRole)

	// get all current 'roster' roles (to remove)
	rosterRoles := currentRosterRoles(currentUser)
	rolesToRemove = append(rolesToRemove, rosterRoles...)

	// get correct roster role (to add)
	rosterRole := correctRosterRole(currentUser)
	rolesToAdd = append(rolesToAdd, rosterRole)

	log.WithFields(log.Fields{
		"new_roles":      extensions.HumanReadableRoles(rolesToAdd...),
		"removing_roles": extensions.HumanReadableRoles(rolesToRemove...),
		"new_nick":       newNick,
	}).Info("Updating user")

	// update user
	extensions.UpdateCavUser(s, &extensions.CavUserUpdate{
		DiscordUser: currentUser.dcMember,
		AddRoles:    rolesToAdd,
		RemoveRoles: rolesToRemove,
	})

	// (manually update user until PR to update discord-go is merged)
	// https://github.com/bwmarrin/discordgo/pull/1122
	if currentUser.dcMember.Nick != newNick {
		guildId := viper.GetString("discord.guild-id")
		if err := s.GuildMemberNickname(guildId, currentUser.dcMember.User.ID, newNick); err != nil {
			log.WithFields(log.Fields{
				"old_nick": currentUser.dcMember.Nick,
				"new_nick": newNick,
				"err":      err,
			}).Error("could not update nickname")
		}
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Account synchronised",
		},
	})

	if err != nil {
		return
	}
}

/////////////////////////////////////////////////////////////////
///////////////////// Sync logic steps below ////////////////////
/////////////////////////////////////////////////////////////////

func currentRankRoles(pair *cavDcPair) []string {
	var res []string

	for _, role := range pair.dcMember.Roles {
		if _, found := cavDiscord.RoleRankMapping[cavDiscord.DiscordRankRoleId(role)]; found {
			res = append(res, role)
		}
	}

	return res
}

func correctRankRole(pair *cavDcPair) string {
	if _, found := cavDiscord.RankRoleMapping[proto.RankType(pair.cavProfile.Rank.RankId)]; !found {
		log.WithFields(log.Fields{
			"rank":             pair.cavProfile.GetRank().GetRankFull(),
			"username":         pair.cavProfile.GetUser().GetUsername(),
			"discord_nickname": pair.dcMember.Nick,
		}).Error("could not find a matching discord role for profile rank")
	}

	return string(cavDiscord.RankRoleMapping[proto.RankType(pair.cavProfile.Rank.RankId)])
}

func currentRosterRoles(pair *cavDcPair) []string {
	var res []string

	for _, role := range pair.dcMember.Roles {
		if _, found := cavDiscord.RoleRosterMapping[cavDiscord.DiscordRosterRole(role)]; found {
			res = append(res, role)
		}
	}

	return res
}

func correctRosterRole(pair *cavDcPair) string {
	if _, found := cavDiscord.RosterRoleMapping[pair.cavProfile.Roster]; !found {
		log.WithFields(log.Fields{
			"roster":           pair.cavProfile.Roster.String(),
			"username":         pair.cavProfile.User.GetUsername(),
			"discord_nickname": pair.dcMember.Nick,
		}).Error("could not find a matching discord role for profile roster")
	}

	return string(cavDiscord.RosterRoleMapping[pair.cavProfile.Roster])
}

func syncRosterOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {

	var correctRosterRole cavDiscord.DiscordRosterRole

	if rosterRole, found := cavDiscord.RosterRoleMapping[cavUser.Roster]; !found {
		return fmt.Errorf("no matching discord role for roster %s", cavUser.Roster)
	} else if cavUser.Primary.PositionTitle == cavDiscord.RETIRED_POSITION_TITLE {
		correctRosterRole = cavDiscord.Discord7CavRet
	} else {
		correctRosterRole = rosterRole
	}

	guildId := viper.GetString("discord.guild-id")

	var currentRoster proto.RosterType
	var currentRosterRole string
	for _, role := range user.Roles {
		if value, found := cavDiscord.RoleRosterMapping[cavDiscord.DiscordRosterRole(role)]; found || cavDiscord.SpecialRETRoleCheck(role) {
			currentRosterRole = role

			// special case for retired members
			if !found {
				currentRoster = proto.RosterType_ROSTER_TYPE_PAST_MEMBERS
				break
			}

			currentRoster = value
			break
		}
	}

	if currentRoster == cavUser.Roster {
		log.Println("User roster already sync'd - skipping roster sync")
		return nil
	}

	if currentRosterRole != "" {
		err := session.GuildMemberRoleRemove(guildId, user.User.ID, currentRosterRole)
		if err != nil {
			log.Printf("error removing role, user: %s, roster role: %s, on guild: %s,  %v", user.User.ID, currentRosterRole, guildId, err)
			return err
		}
	}

	err := session.GuildMemberRoleAdd(guildId, user.User.ID, string(correctRosterRole))
	if err != nil {
		log.Printf("error adding role, user: %s, roster role: %s, on guild: %s,  %v", user.User.ID, string(correctRosterRole), guildId, err)
		return err
	}

	return nil
}

func syncRankOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {

	skipRoleChange := false
	skipNickChange := false
	skipRankGroupRoleChange := false

	// Sync correct rank
	rankRoleId := cavDiscord.RankRoleMapping[proto.RankType(cavUser.Rank.RankId)]

	if rankRoleId == "" {
		return fmt.Errorf("no matching discord role for rank %s", cavUser.Rank.RankShort)
	}
	guildId := viper.GetString("discord.guild-id")

	var currentRank proto.RankType
	var currentRankRole string
	var currentRankGroupRole string
	for _, role := range user.Roles {

		// are they in Officer or NCO groups already
		if value, found := cavDiscord.DiscordRankGroupMap[cavDiscord.DiscordRankGroupRole(role)]; found {
			log.Infof("user: %s, found matching rankGroup role: %s, type: %s", user.Nick, role, value)
			currentRankGroupRole = string(value)
		}

		// are they in a rank specific group 'SPC' 'CPL' etc
		if value, found := cavDiscord.RoleRankMapping[cavDiscord.DiscordRankRoleId(role)]; found {
			log.Infof("user: %s, found matching rank role: %s, rank: %s\n", user.Nick, role, value.String())
			currentRankRole = role
			currentRank = value
		}
	}

	if currentRank == proto.RankType(cavUser.Rank.RankId) {
		log.Warnf("User rank already sync'd - skipping rank sync")
		skipRoleChange = true
	}

	correctGroupRole := cavDiscord.GetDiscordRankGroupRole(proto.RankType(cavUser.Rank.RankId))
	if correctGroupRole == cavDiscord.DiscordRankGroupRole(currentRankGroupRole) {
		skipRankGroupRoleChange = true
	}

	if !skipRankGroupRoleChange {
		log.Printf("Updating rank group role: %s", correctGroupRole)
		if currentRankGroupRole != "" {
			err := session.GuildMemberRoleRemove(guildId, user.User.ID, currentRankGroupRole)
			if err != nil {
				log.Printf("error removing role, user: %s, group rank role: %s, on guild: %s,  %v", user.User.Username, currentRankGroupRole, guildId, err)
				return err
			}
		}

		err := session.GuildMemberRoleAdd(guildId, user.User.ID, string(correctGroupRole))
		if err != nil {
			log.Printf("error adding role, user: %s, rank group role: %s, on guild: %s,  %v", user.User.Username, string(correctGroupRole), guildId, err)
			return err
		}
	}

	if !skipRoleChange {

		if currentRankRole != "" {
			err := session.GuildMemberRoleRemove(guildId, user.User.ID, currentRankRole)
			if err != nil {
				log.Errorf("error removing role, user: %s, rank role: %s, on guild: %s,  %v", user.User.ID, currentRankRole, guildId, err)
				return err
			}
		}

		err := session.GuildMemberRoleAdd(guildId, user.User.ID, string(rankRoleId))
		if err != nil {
			log.Errorf("error adding role, user: %s, rank role: %s, on guild: %s,  %v", user.User.Username, string(rankRoleId), guildId, err)
			return err
		}
	}

	newNick := cavDiscord.GenerateCavNickName(cavUser)
	if newNick == user.Nick {
		log.Warnf("User nick already sync'd - skipping nick sync")
		skipNickChange = true
	}

	if !skipNickChange {
		err := session.GuildMemberNickname(guildId, user.User.ID, newNick)
		if err != nil {
			log.Errorf("error updating user nick, user: %s, nick: %s, on guild: %s,  %v", user.User.ID, newNick, guildId, err)
			return err
		}
	}

	return nil
}

func ErrorWithCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Sorry, there was en error performing this command. Please try again later",
		},
	})
	if err != nil {
		return
	}
}
