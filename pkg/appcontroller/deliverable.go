package appcontroller

import nomad "github.com/hashicorp/nomad/api"

type Deliverable interface {
	ID() string
	Dependencies() []Deliverable
	Resources() []*nomad.Job
}

type deliverable struct {
	id           string
	dependencies []Deliverable
	resources    []*nomad.Job
}

func (d *deliverable) ID() string {
	return d.id
}

func (d *deliverable) Dependencies() []Deliverable {
	return d.dependencies
}

func (d *deliverable) Resources() []*nomad.Job {
	return d.resources
}
