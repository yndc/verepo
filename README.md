# Repomaster

A tool written in Golang to manage Go repositories that adheres the [standard project layout](https://github.com/golang-standards/project-layout).

## List versions

Command: `repomaster list`

List all applications under the `cmd` folder with their versions stored in git tags.

## Get version

Command: `repomaster get <application>`

Get the current version of the specified app

## Setting pre-release tags 

Command: `repomaster prerelease <application> <prerelease>`

Set the pre-release tag of the specified application, without touching the version.

## Setting version

Sub-command: `repomaster set <application> <version> [parameters]`

`-p --prerelease <name>`

Updates the version of the application to the provided semver version. Set the `--prerelease <name>` parameter to change the pre-release tag for this version.

## Bumping version

Command: `repomaster bump <application> <?level>`

This command will bump the version based on semver semantics. Without the `level` parameter set, it will bump only the patch version. Use `minor` to bump the minor version, `major` to bump the major version, and `prerelease` to bump the prerelease version. This command will automatically make a commit for the version bump!

## Releasing a version

Sub-command: `repomaster release <application>`

These actions will be executed:
- Tag the current commit without any pre-release tags.
- Move the changes under unreleased changes to the current release version