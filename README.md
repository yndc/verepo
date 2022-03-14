# Repomaster

A tool written in Golang to manage Go repositories that adheres the [standard project layout](https://github.com/golang-standards/project-layout).

## Affected

Command: `repomaster affected`

Flags:


This command will list all applications with affected

## Application versioning

Command: `repomaster version`

This tool assumes that each folder inside the `/cmd` directory represents a service or an executable, with its own version.

### Updating Versions

Sub-command: `repomaster version update <application> [parameters]`

`-a --all`
`-p --patch`
`-n --minor`
`-j --major`

Updates the version
