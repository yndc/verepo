# VeRepo

A simple tool written in Golang to manage project versioning and changelog.

## Multi-mode

`verepo` supports repositories with multiple projects in them. Currently it only supports Golang projects with `cmd/[project]` folder architecture.

### List projects

Command: `verepo --list`

List all projects under the `cmd` folder with their versions.

### Get version

Command: `verepo <project> version`

Get the current version of the specified project.

### Adding a new entry to the changelog 

Command: `verepo <project> log <message> [flags]`

Flags:
- `--add` 
- `--fix` 
- `--remove` 

Add a new entry into the unreleased section of the changelog. By default, the message will be added into the `Changes` section. Use the flags to control which section should the log be placed.

### Setting pre-release tags 

Command: `verepo <project> prerelease <tag>`

Set the pre-release tag of the specified project, without touching the version. The new pre-release version must be on a higher precedence than the current pre-release.

### Setting version

Sub-command: `verepo <project> set <version> [flags]`

Flags:
- `-p --prerelease <name>`

Updates the version of the project to the provided semver version. Set the `--prerelease <name>` parameter to change the pre-release tag for this version.

### Bumping version

Command: `verepo <project> bump [flags]`

This command will bump the version based on semver semantics. Without the `level` parameter set, it will bump only the patch version. Use `minor` to bump the minor version, `major` to bump the major version.

### Releasing a version

Sub-command: `verepo <project> release`

These actions will be executed:
- Tag the current commit without any pre-release tags.
- Move the changes under unreleased changes to the current release version