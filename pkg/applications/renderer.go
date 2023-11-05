package applications

import (
	"marcocd/pkg/domains"

	nomad "github.com/hashicorp/nomad/api"
)

type TemplateRenderer interface {
	Render(templateFile string, values domains.Values) (*nomad.Job, error)
}
