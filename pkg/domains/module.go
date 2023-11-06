package domains

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ModuleManifest struct {
	Name         string        `yaml:"module"`
	Deliverables []Deliverable `yaml:"deliverables"`
	Dependencies []string      `yaml:"dependencies"`
	Values       Values        `yaml:"values"`
}

// module level values will be override by deliverable level values
// what if Module dependencies not found????

func NewModuleManifestFromFile(manifestPath string) (*ModuleManifest, error) {
	f, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, err
	}

	var moduleManifest ModuleManifest

	if err := yaml.Unmarshal(f, &moduleManifest); err != nil {
		return nil, err
	}

	return nil, nil
}
