package applications

import (
	"fmt"

	nomad "github.com/hashicorp/nomad/api"
)

type Deployer interface {
	Deploy(string) error
}

type deployer struct {
	nc       *nomad.Client
	renderer TemplateRenderer
}

func NewDeployer(
	nc *nomad.Client,
	renderer TemplateRenderer) Deployer {
	return &deployer{
		nc:       nc,
		renderer: renderer,
	}
}

func (d *deployer) Deploy(module string) error {
	jobSpecs, err := d.renderer.Render()
	if err != nil {
		return err
	}

	jobHandler := d.nc.Jobs()

	for _, jobSpec := range jobSpecs {
		job, err := jobHandler.ParseHCL(jobSpec, false)
		if err != nil {
			return err
		}

		job.SetMeta("module", module)
		job.SetMeta("managed-by", "marcocd")
		res, _, _ := jobHandler.Register(job, nil)
		fmt.Println(res.EvalID)
	}

	return nil
}
