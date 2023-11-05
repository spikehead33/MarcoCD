package applications

import (
	nomad "github.com/hashicorp/nomad/api"
)

type TemplateRenderer interface {
	Render() ([]*nomad.Job, error)
}
