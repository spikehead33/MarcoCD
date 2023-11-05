package domains

type Deliverable struct {
	Name         string   `yaml:"name"`
	Resources    []string `yaml:"resources"`
	Values       Values
	Dependencies []string `yaml:"dependencies"`
}
