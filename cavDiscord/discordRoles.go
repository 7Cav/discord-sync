package cavDiscord

import (
	"fmt"
	"github.com/7cav/api/proto"
	"log"
)

type DiscordRankRoleId string

const (
	discord7CavRCT DiscordRankRoleId = "899328824871882752"
	discord7CavPVT DiscordRankRoleId = "899328617081864202"
	discord7CavPFC DiscordRankRoleId = "899328498013966417"
	discord7CavSPC DiscordRankRoleId = "899328418766815283"
	discord7CavCPL DiscordRankRoleId = "899328353511813160"
	discord7CavSGT DiscordRankRoleId = "899328273752928318"
	discord7CavSSG DiscordRankRoleId = "899328187820044359"
	discord7CavSFC DiscordRankRoleId = "899328106366660638"
	discord7CavMSG DiscordRankRoleId = "899328027538907216"
	discord7Cav1SG DiscordRankRoleId = "899327878615957535"
	discord7CavSGM DiscordRankRoleId = "899327773091459213"
	discord7CavCSM DiscordRankRoleId = "879186756937846796"
	discord7CavWO1 DiscordRankRoleId = "899327096697024572"
	discord7CavCW2 DiscordRankRoleId = "899327005122764852"
	discord7CavCW3 DiscordRankRoleId = "899326922381746206"
	discord7CavCW4 DiscordRankRoleId = "899326840487940137"
	discord7CavCW5 DiscordRankRoleId = "899326766664003604"
	discord7Cav2LT DiscordRankRoleId = "899326460974759966"
	discord7Cav1LT DiscordRankRoleId = "899326360185610271"
	discord7CavCPT DiscordRankRoleId = "899326238685024267"
	discord7CavMAJ DiscordRankRoleId = "899326126936190986"
	discord7CavLTC DiscordRankRoleId = "899326048590766100"
	discord7CavCOL DiscordRankRoleId = "899325943179538432"
	discord7CavBG  DiscordRankRoleId = "899325493600473088"
	discord7CavMG  DiscordRankRoleId = "899325397391523920"
	discord7CavLTG DiscordRankRoleId = "899325154402914315"
	discord7CavGEN DiscordRankRoleId = "899325051936079892"
	discord7CavGOA DiscordRankRoleId = "899324897925414993"
)

var RankRoleMapping = map[proto.RankType]DiscordRankRoleId{
	proto.RankType_RANK_TYPE_RCT: discord7CavRCT,
	proto.RankType_RANK_TYPE_PVT: discord7CavPVT,
	proto.RankType_RANK_TYPE_PFC: discord7CavPFC,
	proto.RankType_RANK_TYPE_SPC: discord7CavSPC,
	proto.RankType_RANK_TYPE_CPL: discord7CavCPL,
	proto.RankType_RANK_TYPE_SGT: discord7CavSGT,
	proto.RankType_RANK_TYPE_SSG: discord7CavSSG,
	proto.RankType_RANK_TYPE_SFC: discord7CavSFC,
	proto.RankType_RANK_TYPE_MSG: discord7CavMSG,
	proto.RankType_RANK_TYPE_1SG: discord7Cav1SG,
	proto.RankType_RANK_TYPE_SGM: discord7CavSGM,
	proto.RankType_RANK_TYPE_CSM: discord7CavCSM,
	proto.RankType_RANK_TYPE_WO1: discord7CavWO1,
	proto.RankType_RANK_TYPE_CW2: discord7CavCW2,
	proto.RankType_RANK_TYPE_CW3: discord7CavCW3,
	proto.RankType_RANK_TYPE_CW4: discord7CavCW4,
	proto.RankType_RANK_TYPE_CW5: discord7CavCW5,
	proto.RankType_RANK_TYPE_2LT: discord7Cav2LT,
	proto.RankType_RANK_TYPE_1LT: discord7Cav1LT,
	proto.RankType_RANK_TYPE_CPT: discord7CavCPT,
	proto.RankType_RANK_TYPE_MAJ: discord7CavMAJ,
	proto.RankType_RANK_TYPE_LTC: discord7CavLTC,
	proto.RankType_RANK_TYPE_COL: discord7CavCOL,
	proto.RankType_RANK_TYPE_BG:  discord7CavBG,
	proto.RankType_RANK_TYPE_MG:  discord7CavMG,
	proto.RankType_RANK_TYPE_LTG: discord7CavLTG,
	proto.RankType_RANK_TYPE_GEN: discord7CavGEN,
	proto.RankType_RANK_TYPE_GOA: discord7CavGOA,
}

var RoleRankMapping = map[DiscordRankRoleId]proto.RankType{
	discord7CavRCT: proto.RankType_RANK_TYPE_RCT,
	discord7CavPVT: proto.RankType_RANK_TYPE_PVT,
	discord7CavPFC: proto.RankType_RANK_TYPE_PFC,
	discord7CavSPC: proto.RankType_RANK_TYPE_SPC,
	discord7CavCPL: proto.RankType_RANK_TYPE_CPL,
	discord7CavSGT: proto.RankType_RANK_TYPE_SGT,
	discord7CavSSG: proto.RankType_RANK_TYPE_SSG,
	discord7CavSFC: proto.RankType_RANK_TYPE_SFC,
	discord7CavMSG: proto.RankType_RANK_TYPE_MSG,
	discord7Cav1SG: proto.RankType_RANK_TYPE_1SG,
	discord7CavSGM: proto.RankType_RANK_TYPE_SGM,
	discord7CavCSM: proto.RankType_RANK_TYPE_CSM,
	discord7CavWO1: proto.RankType_RANK_TYPE_WO1,
	discord7CavCW2: proto.RankType_RANK_TYPE_CW2,
	discord7CavCW3: proto.RankType_RANK_TYPE_CW3,
	discord7CavCW4: proto.RankType_RANK_TYPE_CW4,
	discord7CavCW5: proto.RankType_RANK_TYPE_CW5,
	discord7Cav2LT: proto.RankType_RANK_TYPE_2LT,
	discord7Cav1LT: proto.RankType_RANK_TYPE_1LT,
	discord7CavCPT: proto.RankType_RANK_TYPE_CPT,
	discord7CavMAJ: proto.RankType_RANK_TYPE_MAJ,
	discord7CavLTC: proto.RankType_RANK_TYPE_LTC,
	discord7CavCOL: proto.RankType_RANK_TYPE_COL,
	discord7CavBG:  proto.RankType_RANK_TYPE_BG,
	discord7CavMG:  proto.RankType_RANK_TYPE_MG,
	discord7CavLTG: proto.RankType_RANK_TYPE_LTG,
	discord7CavGEN: proto.RankType_RANK_TYPE_GEN,
	discord7CavGOA: proto.RankType_RANK_TYPE_GOA,
}

type DiscordRosterRole string

const (
	Discord7CavActive    DiscordRosterRole = "437748324960043009"
	Discord7CavReserve   DiscordRosterRole = "690899750425329666"
	Discord7CavELOA      DiscordRosterRole = "937082349848526848"
	Discord7CavRet       DiscordRosterRole = "437748982400417792"
	Discord7CavDisch     DiscordRosterRole = "437749895785480193"
	Discord7CavWOH       DiscordRosterRole = "690899500457525308"
	Discord7CavArlington DiscordRosterRole = "937084272068661289"
)

var RoleRosterMapping = map[DiscordRosterRole]proto.RosterType{
	Discord7CavActive:  proto.RosterType_ROSTER_TYPE_COMBAT,
	Discord7CavReserve: proto.RosterType_ROSTER_TYPE_RESERVE,
	Discord7CavELOA:    proto.RosterType_ROSTER_TYPE_ELOA,
	Discord7CavWOH:     proto.RosterType_ROSTER_TYPE_WALL_OF_HONOR,
	Discord7CavRet:     proto.RosterType_ROSTER_TYPE_PAST_MEMBERS,
	Discord7CavDisch:   proto.RosterType_ROSTER_TYPE_PAST_MEMBERS,
}

var RosterRoleMapping = map[proto.RosterType]DiscordRosterRole{
	proto.RosterType_ROSTER_TYPE_COMBAT:        Discord7CavActive,
	proto.RosterType_ROSTER_TYPE_RESERVE:       Discord7CavReserve,
	proto.RosterType_ROSTER_TYPE_ELOA:          Discord7CavELOA,
	proto.RosterType_ROSTER_TYPE_WALL_OF_HONOR: Discord7CavWOH,
	proto.RosterType_ROSTER_TYPE_ARLINGTON:     Discord7CavArlington,
	proto.RosterType_ROSTER_TYPE_PAST_MEMBERS:  Discord7CavDisch,
	//proto.RosterType_ROSTER_TYPE_PAST_MEMBERS: discord7CavRet,
}

// SpecialRETRoleCheck special function because `proto.RosterType_ROSTER_TYPE_PAST_MEMBERS` maps to both DISCH and RET
// really, this just confirms the role string is the retired role
func SpecialRETRoleCheck(role string) bool {
	return DiscordRosterRole(role) == Discord7CavRet
}

const RETIRED_POSITION_TITLE = "Retired"

type DiscordRankGroupRole string

const (
	discord7CavOfficer DiscordRankGroupRole = "437750376415100928"
	discord7CavNCO     DiscordRankGroupRole = "437750614374613012"
)

var DiscordRankGroupMap = map[DiscordRankGroupRole]DiscordRankGroupRole{
	discord7CavOfficer: discord7CavOfficer,
	discord7CavNCO:     discord7CavNCO,
}

func GetDiscordRankGroupRole(rank proto.RankType) DiscordRankGroupRole {
	log.Printf("Getting discord rank group role for rank: %s", rank)
	switch rank {
	case proto.RankType_RANK_TYPE_GOA,
		proto.RankType_RANK_TYPE_GEN,
		proto.RankType_RANK_TYPE_LTG,
		proto.RankType_RANK_TYPE_MG,
		proto.RankType_RANK_TYPE_BG,
		proto.RankType_RANK_TYPE_COL,
		proto.RankType_RANK_TYPE_LTC,
		proto.RankType_RANK_TYPE_MAJ,
		proto.RankType_RANK_TYPE_CPT,
		proto.RankType_RANK_TYPE_1LT,
		proto.RankType_RANK_TYPE_2LT:
		log.Println("returning officer rank group")
		return discord7CavOfficer
	case proto.RankType_RANK_TYPE_CW5,
		proto.RankType_RANK_TYPE_CW4,
		proto.RankType_RANK_TYPE_CW3,
		proto.RankType_RANK_TYPE_CW2,
		proto.RankType_RANK_TYPE_WO1,
		proto.RankType_RANK_TYPE_CSM,
		proto.RankType_RANK_TYPE_SGM,
		proto.RankType_RANK_TYPE_1SG,
		proto.RankType_RANK_TYPE_MSG,
		proto.RankType_RANK_TYPE_SFC,
		proto.RankType_RANK_TYPE_SSG,
		proto.RankType_RANK_TYPE_SGT,
		proto.RankType_RANK_TYPE_CPL:
		log.Println("returning NCO rank group")
		return discord7CavNCO
	}

	log.Println("returning no rank group")
	return ""
}

func GenerateCavNickName(cavUser *proto.Profile) string {
	prefix := ""
	switch cavUser.Roster {
	case proto.RosterType_ROSTER_TYPE_ELOA:
		prefix = "[ELOA] "
		break
	case proto.RosterType_ROSTER_TYPE_RESERVE:
		prefix = "[AR] "
		break
	case proto.RosterType_ROSTER_TYPE_PAST_MEMBERS:
		if cavUser.Primary.PositionTitle == RETIRED_POSITION_TITLE {
			prefix = "[RET] "
			break
		}
		prefix = "[DISCH] "
		break
	}

	// lol
	if cavUser.User.Username == "Jarvis.A" {
		return fmt.Sprintf("%s%s.Jarvis", prefix, cavUser.Rank.RankShort)
	}

	return fmt.Sprintf("%s%s.%s", prefix, cavUser.Rank.RankShort, cavUser.User.Username)
}
