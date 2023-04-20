# auto-retro

`auto-retro` is a tool supporting Scrum teams by generating reports of team members' activity on GitHub. It approaches to follow the DDD practices; and therefore, the design shall improve over the time.

## External Dependencies
[githubv4](https://github.com/shurcooL/githubv4) is a golang wrapper over the GitHub GraphQL API v4.

## Getting Started
Before running the project please set the environment configuration up. To export the evnironment variables, we use `direnv` tool, you can install it via:
```sh
$ brew install direnv
```
Copy the contents of the `.envrc-sample` file to `.envrc` and prefix variables with `export` and set the values approprietly. `GITHUB_TOKEN` is necessary to authenticate the user, the teams are searched within the `ORGANIZATION` and are filtered by `EXCLUDE_TEAMS`, e.g. if there's a team within your organization that gothers all contributors.

To interact with the core functionality of the tool, please use the Makefile commands.

Build the project:
```sh
$ make build
```

Run the project:
```sh
$ make run
```

## Contribution
Before opening the PR, please check the already existing [issues](github.com/a-millert/auto-retro/issues). If the PR improves the project in the dimension not covered by any issue, please open one and use the issue number in your branch prefix following the pattern `AR-<number>-<details>`. Additionally, the PR title should start with a prefix: `AR-<number>: details`.

### TODO:
- HTTP server
- React client
- Dockerizing
