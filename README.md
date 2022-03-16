# Repomaster

A tool written in Golang to manage Go repositories that adheres the [standard project layout](https://github.com/golang-standards/project-layout).

## List 

Command: `repomaster list`

List all applications with their versions under the `cmd` folder

## Build 

Command: `repomaster build <application>`

Build the specified application using the provided build steps in `repomaster.yaml`. These environment variables will be exported into the build steps command:
- `$APP_ID` the ID of the currently building application
- `$APP_FULL_VERSION` the full semver of the application, with the format `v[major].[minor].[patch]-[prerelease]+[commit]`

The version and pre-release tag will be taken from the `app.json` file found in the app folder. 
- If `version` is not found, `repomaster` will use `v0.0.0` as the version
- If `pre-release` tag is not found, `dev` will be added automatically as the pre-release tag 

The only way to build an image without a pre-release tag is by `repomaster release <application>`

## Setting pre-release tags 

Command: `repomaster prerelease <application> <prerelease>`

Set the pre-release tag of the specified application, without touching the version. This command will automatically make a commit for the updated version!

## Setting version

Sub-command: `repomaster set <application> <version> [parameters]`

`-p --prerelease <name>`

Updates the version of the application to the provided semver version. Set the `--prerelease <name>` parameter to change the pre-release tag for this version. This command will automatically make a commit for the updated version!

## Bumping version

Command: `repomaster bump <application> <?level>`

This command will bump the version based on semver semantics. Without the `level` parameter set, it will bump only the patch version. Use `minor` to bump the minor version, `major` to bump the major version, and `prerelease` to bump the prerelease version. This command will automatically make a commit for the version bump!

## Releasing a version

Sub-command: `repomaster release <application>`

Release the current commit as a stable version, these actions will be executed by `repomaster`:
- Remove the pre-release tag and build metadata in the application's `app.json`
- Commit the tags and build metadata removal
- Tag the commit with the format `[APP_ID]:vX.Y.Z` without tags or metadata
- Build and push the application with the normal version

Note:
- After a successful `release`, you need to change the version through `repomaster version update` or `repomaster version bump` before you can make a newer build.