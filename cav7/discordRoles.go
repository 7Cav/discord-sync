package cav7

import "github.com/7cav/api/proto"

type DiscordRankRoleId string

const (
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
