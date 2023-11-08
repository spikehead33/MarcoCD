package applications

import (
	"fmt"

	nomad "github.com/hashicorp/nomad/api"
)

type Deployer interface {
	Deploy() error
}

type moduleDeployer struct {
	nc         *nomad.Client
	moduleName string
	renderer   TemplateRenderer
}

func NewDeployer(
	moduleName string,
	nc *nomad.Client,
	renderer TemplateRenderer) Deployer {
	return &moduleDeployer{
		moduleName: moduleName,
		nc:         nc,
		renderer:   renderer,
	}
}

func (d *moduleDeployer) Deploy() error {
	jobSpecs, err := d.renderer.Render()
	if err != nil {
		return err
	}

	jobHandler := d.nc.Jobs()

	for _, jobSpec := range jobSpecs {
		fmt.Println(jobSpec)

		job, err := jobHandler.ParseHCL(jobSpec, false)
		if err != nil {
			return err
		}

		job.SetMeta("module", d.moduleName)
		job.SetMeta("managed-by", "marcocd")
		res, _, err := jobHandler.Register(job, nil)
		if err != nil {
			return err
		}
		fmt.Println(res.EvalID)
	}

	return nil
}
