package config

type Container struct {
	DryRun  bool
	Verbose bool
}

var Global Container = Container{}
