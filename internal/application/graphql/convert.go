package application

import (
	"sort"

	core "github.com/a-millert/auto-retro/internal/domain/model"
	misc "github.com/a-millert/auto-retro/internal/pkg"
)

func (resp *LoginResponse) ToDomain() *core.User {
	return &core.User{Username: resp.Viewer.Login}
}

func (resp *TeamsResponse) ToDomain(excludeTeams []string) []*core.Team {
	var teams []*core.Team

	excludeTeamsSet := misc.Deduplicate(excludeTeams)

	for _, team := range resp.Organization.Teams.Nodes {
		// skip excluded teams
		if _, ok := excludeTeamsSet[team.Name]; ok {
			continue
		}

		uniqueMembers := make(map[string]bool)
		var members []core.User

		for _, member := range team.Members.Nodes {
			uniqueMembers[member.Login] = true
		}
		for memberName := range uniqueMembers {
			members = append(members, core.User{
				Username: memberName,
			})
		}
		sort.Slice(members, func(i, j int) bool {
			return members[i].Username < members[j].Username
		})

		teams = append(teams, &core.Team{
			Name:    team.Name,
			Members: members,
		})
	}

	return teams
}
