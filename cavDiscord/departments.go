package cavDiscord

import (
	"fmt"
	"github.com/7cav/api/milpacs"
)

type discordRoleDepartmentLevel int

const (
	STAFF_ROLE = 0
	LEADS_ROLE = 1
	HQ_ROLE    = 2
)

type Department interface {
	GetRolesForBillet(billet *milpacs.Position) *[]string
}

type departmentGrouping struct {
	hqRoleId      string
	snrLeadRoleId string
	staffRoleId   string
}

type S1Department struct {
	departmentGrouping
}

type S6Department struct {
	departmentGrouping
}

func S6() *S6Department {
	return &S6Department{
		departmentGrouping{
			hqRoleId:      "asd",
			snrLeadRoleId: "asd",
			staffRoleId:   "asd",
		},
	}
}

func (dept S6Department) GetRolesForBillet(billet *milpacs.Position) *[]string {
	var roles []string

	tier := func() discordRoleDepartmentLevel {
		switch billet.PositionId {
		case 50, 51:
			return HQ_ROLE // 1IC, 2IC
		case 52, 53, 56, 57:
			return LEADS_ROLE // Lead Dev, Snr Dev, Lead Games, Snr Games
		default:
			return STAFF_ROLE
		}
	}()

	switch tier {
	case HQ_ROLE:
		roles = append(roles, dept.hqRoleId, dept.staffRoleId);
		break
	case LEADS_ROLE:
		roles = append(roles, dept.snrLeadRoleId, dept.staffRoleId);
		break
	default:
		roles = append(roles, dept.staffRoleId);
		break
	}

	return &roles
}

var positionIdsToDepartment = map[uint64]Department{
	50: S6(),
	51: S6(),
	52: S6(),
	53: S6(),
	54: S6(),
	55: S6(),
	56: S6(),
	57: S6(),
	58: S6(),
	59: S6(),
}

func GetDepartment(billet *milpacs.Position) (Department, error) {
	if _, found := positionIdsToDepartment[billet.PositionId]; !found {
		return nil, fmt.Errorf("could not find matching department for position ID %d", billet.PositionId)
	}

	return positionIdsToDepartment[billet.PositionId], nil
}
