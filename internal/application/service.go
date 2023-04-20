package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/shurcooL/githubv4"

	application "github.com/a-millert/auto-retro/internal/application/graphql"
	core "github.com/a-millert/auto-retro/internal/domain/model"
	"github.com/a-millert/auto-retro/internal/pkg/fail"
)

type Service interface {
	Login() *core.User
	Teams() *core.User
}

type Application struct {
	ctx          *context.Context
	client       *githubv4.Client
	org          string
	excludeTeams []string
}

func New(ctx *context.Context,
	ghClient *githubv4.Client,
	org string,
	excludeTeams string,
) *Application {
	return &Application{
		ctx:          ctx,
		client:       ghClient,
		org:          org,
		excludeTeams: strings.Split(excludeTeams, ","),
	}
}

func (a *Application) Login() (*core.User, error) {
	var loginResponse application.LoginResponse
	err := a.client.Query(*a.ctx, &loginResponse, nil)
	if err != nil || len(loginResponse.Viewer.Login) == 0 {
		return nil, fail.New(err, "Failed to authenticate with the token").Abort()
	}
	fmt.Println("Successfully logged in as:", loginResponse.Viewer.Login)

	return loginResponse.ToDomain(), nil
}

func (a *Application) Teams() ([]*core.Team, error) {
	var teamsResponse application.TeamsResponse
	vars := teamsResponse.InputVariables(a.org)
	err := a.client.Query(*a.ctx, &teamsResponse, vars)
	if err != nil {
		return nil, fail.New(err, "Failed to get the teams").Abort()
	} else if len(teamsResponse.Organization.Teams.Nodes) == 0 {
		fmt.Println("User doesn't belong to any teams")
		return nil, nil
	}

	return teamsResponse.ToDomain(a.excludeTeams), nil
}
