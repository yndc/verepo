package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, &Global)
}
