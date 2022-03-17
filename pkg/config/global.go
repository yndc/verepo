package config

type Container struct {
	App     string
	DryRun  bool
	Verbose bool
}

var Global Container = Container{}
