package bot

import (
	"fmt"
	"github.com/7cav/api/proto"
	"github.com/7cav/discord-sync/cavAPI"
	"github.com/7cav/discord-sync/cavDiscord"
	"github.com/7cav/discord-sync/keycloak"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"log"
)

type CavDiscordServer struct {
	serverId    string
	roleMapping map[string]roleMapping
}

type roleMapping struct {
	discordRoleId    string
	milpacPositionId uint64
}

var SyncCommandName = "sync"

func SyncCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        SyncCommandName,
		Description: "Sync account to 7Cav.us",
	}
}

func HandleSync(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("Sync command")

	log.Printf("Attempt get kc user for %s\n", i.Member.User.ID)

	kcUser, err := keycloak.KCUserViaDiscordID(i.Member.User.ID)

	if err != nil {
		AskToConnectDiscord(s, i)
		return
	}

	cavUser := cavAPI.GetUserViaKCID(kcUser.ID)

	err = syncRankOnCoreDiscord(s, i.Member, cavUser)
	if err != nil {
		ErrorWithCommand(s, i)
		return
	}

	err = syncRosterOnCoreDiscord(s, i.Member, cavUser)
	if err != nil {
		ErrorWithCommand(s, i)
		return
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

func syncRosterOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {

	var correctRosterRole cavDiscord.DiscordRosterRole

	if rosterRole, found := cavDiscord.RosterRoleMapping[cavUser.Roster]; !found {
		return fmt.Errorf("no matching discord role for roster %s", cavUser.Roster)
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
			log.Fatalf("error removing role, user: %s, roster role: %s, on guild: %s,  %v", user.User.ID, currentRosterRole, guildId, err)
			return err
		}
	}

	err := session.GuildMemberRoleAdd(guildId, user.User.ID, string(correctRosterRole))
	if err != nil {
		log.Fatalf("error adding role, user: %s, roster role: %s, on guild: %s,  %v", user.User.ID, string(correctRosterRole), guildId, err)
		return err
	}

	return nil
}

func syncRankOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {

	skipRoleChange := false
	skipNickChange := false

	// Sync correct rank
	rankRoleId := cavDiscord.RankRoleMapping[proto.RankType(cavUser.Rank.RankId)]

	if rankRoleId == "" {
		return fmt.Errorf("no matching discord role for rank %s", cavUser.Rank.RankShort)
	}
	guildId := viper.GetString("discord.guild-id")

	var currentRank proto.RankType
	var currentRankRole string
	for _, role := range user.Roles {
		if value, found := cavDiscord.RoleRankMapping[cavDiscord.DiscordRankRoleId(role)]; found {
			log.Printf("user: %s, found matching rank role: %s, rank: %s\n", user.Nick, role, value.String())
			currentRankRole = role
			currentRank = value
			break
		}
	}

	if currentRank == proto.RankType(cavUser.Rank.RankId) {
		log.Println("User rank already sync'd - skipping rank sync")
		skipRoleChange = true
	}

	if !skipRoleChange {

		if currentRankRole != "" {
			err := session.GuildMemberRoleRemove(guildId, user.User.ID, currentRankRole)
			if err != nil {
				log.Fatalf("error removing role, user: %s, rank role: %s, on guild: %s,  %v", user.User.ID, currentRankRole, guildId, err)
				return err
			}
		}

		err := session.GuildMemberRoleAdd(guildId, user.User.ID, string(rankRoleId))
		if err != nil {
			log.Fatalf("error adding role, user: %s, rank role: %s, on guild: %s,  %v", user.User.ID, string(rankRoleId), guildId, err)
			return err
		}
	}

	newNick := cavDiscord.GenerateCavNickName(cavUser)
	if newNick == user.Nick {
		log.Println("User nick already sync'd - skipping nick sync")
		skipNickChange = true
	}

	if !skipNickChange {
		err := session.GuildMemberNickname(guildId, user.User.ID, newNick)
		if err != nil {
			log.Fatalf("error updating user nick, user: %s, nick: %s, on guild: %s,  %v", user.User.ID, newNick, guildId, err)
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
