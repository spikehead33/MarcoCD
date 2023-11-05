package manifest_reader

import (
	"marcocd/pkg/domains"
	"os"

	"gopkg.in/yaml.v3"
)

type ModuleManifestReader interface {
	Read(string) (*domains.ModuleManifest, error)
}

type moduleManifestReader struct{}

func NewModuleManifestReader() ModuleManifestReader {
	return &moduleManifestReader{}
}

func (reader *moduleManifestReader) Read(manifest string) (*domains.ModuleManifest, error) {
	f, err := os.ReadFile(manifest)
	if err != nil {
		return nil, err
	}

	var moduleManifest domains.ModuleManifest

	if err := yaml.Unmarshal(f, &moduleManifest); err != nil {
		return nil, err
	}

	return &moduleManifest, nil
}
