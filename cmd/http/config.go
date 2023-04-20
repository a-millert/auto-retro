package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Env struct defines all required environment variables for the app to run.
type Env struct {
	GitHubToken string `envconfig:"GITHUB_TOKEN" default:""`

	Organization string `envconfig:"ORGANIZATION" default:""`

	ExcludeTeams string `envconfig:"EXCLUDE_TEAMS" default:""`
}

// validate is used to check the invariants of values defined by Env.
func (e *Env) validate() error {
	checks := []struct {
		bad    bool
		errMsg error
	}{
		{
			len(e.GitHubToken) == 0,
			fmt.Errorf("empty GitHub token provided"),
		},
	}

	for _, check := range checks {
		if check.bad {
			return check.errMsg
		}
	}

	return nil
}

// ReadFromEnv reads configuration from environmental variables defined by Env.
func ReadFromEnv() (*Env, error) {
	var env Env

	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}
	if err := env.validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return &env, nil
}
