package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	application "github.com/a-millert/auto-retro/internal/application/graphql"
	"github.com/a-millert/auto-retro/internal/pkg/fail"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

func main() {
	// since os.Exit can not handle `defer`.
	// DON'T call `os.Exit` in the any other place.
	if err := realMain(); err != nil {
		os.Exit(exitError)
	}
	os.Exit(exitOK)
}

func realMain() error {
	// Read configurations from environmental variables.
	env, err := ReadFromEnv()
	if err != nil {
		return fail.New(err, "Failed to read environment variables").Abort()
	}

	ctx := context.Background()

	// Set up the GitHub GraphQL API v4 client with the HTTP OAuth2 client.
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: env.GitHubToken},
	)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	ghClient := githubv4.NewClient(httpClient)

	// Sanity-check: Authenticate the user.
	var loginQuery application.LoginQuery
	err = ghClient.Query(ctx, &loginQuery, nil)
	if err != nil || len(loginQuery.Viewer.Login) == 0 {
		return fail.New(err, "Failed to authenticate with the token").Abort()
	}
	fmt.Println("Successfully logged in as:", loginQuery.Viewer.Login)

	// application builder with DI

	return nil
}
