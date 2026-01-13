package api

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ActionFile struct {
	Actions map[string]APIAction `yaml:"actions"`
}

type APIAction struct {
	Method      string                 `yaml:"method"`
	Path        string                 `yaml:"path"`
	Description string                 `yaml:"description,omitempty"`
	Body        map[string]interface{} `yaml:"body,omitempty"`
}

// LoadActions loads YAML file from path
func LoadActions(path string) (*ActionFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var af ActionFile
	if err := yaml.Unmarshal(data, &af); err != nil {
		return nil, err
	}
	return &af, nil
}
