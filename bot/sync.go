package bot

import (
	"fmt"
	"github.com/7cav/api/proto"
	"github.com/7cav/discord-sync/cav7"
	"github.com/7cav/discord-sync/cavAPI"
	"github.com/7cav/discord-sync/keycloak"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CavDiscordServer struct {
	serverId    string
	roleMapping map[string]roleMapping
}

type roleMapping struct {
	discordRoleId    string
	milpacPositionId uint64
}

const discord7CavActive = "437748324960043009"

var SyncCommandName = "sync"

func SyncCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        SyncCommandName,
		Description: "Sync account to 7Cav.us",
	}
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

	err = syncRankOnCoreDiscord(s, i.Member, cavUser)

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

//func syncRosterOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {
//
//}

func syncRankOnCoreDiscord(session *discordgo.Session, user *discordgo.Member, cavUser *proto.Profile) error {

	skipRoleChange := false
	skipNickchange := false

	// Sync correct rank
	rankRoleId := cav7.RankRoleMapping[proto.RankType(cavUser.Rank.RankId)]

	if rankRoleId == "" {
		return fmt.Errorf("no matching discord role for rank %s", cavUser.Rank.RankShort)
	}
	guildId := viper.GetString("discord.guild-id")

	var currentRank proto.RankType
	var currentRankRole string
	for _, role := range user.Roles {
		if value, ok := cav7.RoleRankMapping[cav7.DiscordRankRoleId(role)]; ok {
			log.Infof("user: %s, found matching rank role: %s, rank: %s\n", user.Nick, role, value.String())
			currentRankRole = role
			currentRank = value
			break
		}
	}

	if currentRank == proto.RankType(cavUser.Rank.RankId) {
		log.Warnf("User rank already sync'd - skipping rank sync")
		skipRoleChange = true
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
			log.Errorf("error adding role, user: %s, rank role: %s, on guild: %s,  %v", user.User.ID, string(rankRoleId), guildId, err)
			return err
		}
	}

	newNick := generateCavNickName(cavUser)
	if newNick == user.Nick {
		log.Warnf("User nick already sync'd - skipping nick sync")
		skipNickchange = true
	}

	if !skipNickchange {
		err := session.GuildMemberNickname(guildId, user.User.ID, newNick)
		if err != nil {
			log.Errorf("error updating user nick, user: %s, nick: %s, on guild: %s,  %v", user.User.ID, newNick, guildId, err)
			return err
		}
	}

	return nil
}

func generateCavNickName(cavUser *proto.Profile) string {
	// lol
	if cavUser.User.Username == "Jarvis.A" {
		return fmt.Sprintf("%s.Jarvis", cavUser.Rank.RankShort)
	}

	return fmt.Sprintf("%s.%s", cavUser.Rank.RankShort, cavUser.User.Username)
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
