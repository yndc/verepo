# Repomaster

A tool written in Golang to manage Go repositories that adheres the [standard project layout](https://github.com/golang-standards/project-layout).

## List 

Command: `repomaster list`

List all applications with their versions under the `cmd` folder

## Build 

Command: `repomaster build`

Build all applications that has code changes. The image will be tagged with the format `[app]:[version]-[prerelease]+[commit]`. The version and pre-release tag will be taken from the `app.json` file found in the app folder. 
- If `version` is not found, `repomaster` will use `v0.0.0` as the version
- If `pre-release` tag is not found, `dev` will be added automatically as the pre-release tag 

The only way to build an image with version as the only tag is by `repomaster release`

## Updating versions

### Manual

Sub-command: `repomaster version update <application> <version> [parameters]`

`-a --all`
`-p --prerelease <name>`

Updates the version of the application to the provided semver version. Set the `--prerelease <name>` parameter to change the pre-release tag for this version. This command will automatically make a commit for the updated version!

### Bumping the version

Command: `repomaster version bump <application>`

`-n --minor`
`-j --major`
`-p --prerelease <name>`

This command will bump the version based on semver semantics. Without a parameter, it will bump only the patch version. Use `--minor` to bump the minor version and `--major` to bump the major version. Set the `--prerelease <name>` parameter to change the pre-release tag for this version. This command will automatically make a commit for the version bump!

## Releasing a version

Sub-command: `repomaster release <application>`

Release the current commit as a stable version, these actions will be executed by `repomaster`:
- Remove the pre-release tag and build metadata in the application's `app.json`
- Commit the tags and build metadata removal
- Tag the commit with the format `Release application:vX.Y.Z` without tags or metadata
- Build and push the application with the normal version

Note:
- After a successful `release`, you need to change the version through `repomaster version update` or `repomaster version bump` before you can make a newer build.