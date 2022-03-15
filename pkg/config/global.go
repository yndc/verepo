package config

type Container struct {
	Build   []string `yaml:"build"`
	DryRun  bool
	Verbose bool
}

var Global Container = Container{}
