package appcontroller

type Module interface {
	ID() string
	Dependencies() []string
	Deliverables() []Deliverable
}

type module struct {
	id           string
	dependencies []string
	deliverables []Deliverable
}

func (m *module) ID() string {
	return m.id
}

func (m *module) Dependencies() []string {
	return m.dependencies
}

func (m *module) Deliverable() []Deliverable {
	return m.deliverables
}
