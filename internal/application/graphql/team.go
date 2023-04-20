package application

import (
	"github.com/shurcooL/githubv4"
)

const (
	memberRole = githubv4.TeamRoleMember
	topMembers = githubv4.Int(12)
	pageSize   = githubv4.Int(100)
)

type TeamsResponse struct {
	Organization struct {
		Teams struct {
			Nodes []struct {
				Name    string
				Members struct {
					Nodes []struct {
						Name  string
						Login string
					}
				} `graphql:"members(first: $firstMembers)"`
			}
		} `graphql:"teams(first: $first, role: $role)"`
	} `graphql:"organization(login: $orgName)"`
}

func (resp TeamsResponse) InputVariables(organization string) map[string]interface{} {
	return map[string]interface{}{
		"orgName":      githubv4.String(organization),
		"first":        pageSize,
		"role":         memberRole,
		"firstMembers": topMembers,
	}
}
