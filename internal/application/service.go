package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	application "github.com/a-millert/auto-retro/internal/application/graphql"
	core "github.com/a-millert/auto-retro/internal/domain/model"
	"github.com/a-millert/auto-retro/internal/pkg/fail"
)

type Service interface {
	Login() *core.User
	Teams() *core.User
}

type Application struct {
	ctx          context.Context
	client       *githubv4.Client
	org          string
	excludeTeams []string
}

func New(
	githubToken string,
	organization string,
	excludeTeams string,
) *Application {
	ctx := context.Background()

	// Set up the GitHub GraphQL API v4 client with the HTTP OAuth2 client.
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	ghClient := githubv4.NewClient(httpClient)

	return &Application{
		ctx:          ctx,
		client:       ghClient,
		org:          organization,
		excludeTeams: strings.Split(excludeTeams, ","),
	}
}

func (a *Application) Run() error {
	// Sanity check: can log in.
	_, err := a.login()
	if err != nil {
		return err
	}
	teams, err := a.teams()
	if err != nil {
		return err
	}
	for _, x := range teams {
		fmt.Printf("Team %s with members: %+v\n", x.Name, x.Members)
	}
	return nil
}

func (a *Application) login() (*core.User, error) {
	var loginResponse application.LoginResponse
	err := a.client.Query(a.ctx, &loginResponse, nil)
	if err != nil || len(loginResponse.Viewer.Login) == 0 {
		return nil, fail.New(err, "Failed to authenticate with the token").Abort()
	}
	fmt.Println("Successfully logged in as:", loginResponse.Viewer.Login)

	return loginResponse.ToDomain(), nil
}

func (a *Application) teams() ([]*core.Team, error) {
	var teamsResponse application.TeamsResponse
	vars := teamsResponse.InputVariables(a.org)
	err := a.client.Query(a.ctx, &teamsResponse, vars)
	if err != nil {
		return nil, fail.New(err, "Failed to get the teams").Abort()
	} else if len(teamsResponse.Organization.Teams.Nodes) == 0 {
		fmt.Println("User doesn't belong to any teams")
		return nil, nil
	}

	return teamsResponse.ToDomain(a.excludeTeams), nil
}
