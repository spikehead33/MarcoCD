package domains

type ModuleManifest struct {
	Name              string        `yaml:"name"`
	Deliverables      []Deliverable `yaml:"deliverables"`
	ModuleLevelValues Values
	Dependencies      []string `yaml:"dependencies"`
}

// module level values will be override by deliverable level values
// what if Module dependencies not found????
