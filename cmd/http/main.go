package main

import (
	"os"

	service "github.com/a-millert/auto-retro/internal/application"
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

	// Build a service and run it.
	return service.New(env.GitHubToken, env.Organization, env.ExcludeTeams).Run()
}
