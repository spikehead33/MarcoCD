package applications

import (
	"marcocd/pkg/domains"

	nomad "github.com/hashicorp/nomad/api"
)

type Reconciler interface {
	Reconcile() error
	UpdateDesiredState([]string) error
}

type reconciler struct {
	nc             *nomad.Client
	desiredState   []*nomad.Job
	moduleManifest *domains.ModuleManifest
}

func (rec *reconciler) Reconcile() error {
	return nil
}

func (rec *reconciler) UpdateDesiredState(states []string) error {
	jobs := rec.nc.Jobs()

	desiredState := make([]*nomad.Job, 512)
	for _, state := range states {
		job, err := jobs.ParseHCL(state, false)
		if err != nil {
			return err
		}

		desiredState = append(desiredState, job)
	}

	rec.desiredState = desiredState
	return nil
}
