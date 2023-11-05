package domains

type Deliverable struct {
	Name         string   `yaml:"name"`
	Resources    []string `yaml:"resources"`
	Dependencies []string `yaml:"dependencies"`
	Values       Values   `yaml:"values"`
}
