package domains

type ModuleManifest struct {
	Name         string        `yaml:"module"`
	Deliverables []Deliverable `yaml:"deliverables"`
	Dependencies []string      `yaml:"dependencies"`
	Values       Values        `yaml:"values"`
}

// module level values will be override by deliverable level values
// what if Module dependencies not found????
