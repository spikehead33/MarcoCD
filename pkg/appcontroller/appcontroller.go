package appcontroller

import (
	nomad "github.com/hashicorp/nomad/api"
)

type AppController interface {
	Reconcile() error
	UpdateDesiredState() error
	UpdateModuleDependencies() error
}

type appcontroller struct {
	moduleDependencies map[Module][]Module
	desiredState       []*nomad.Job
	nc                 *nomad.Client
}

func New(nc *nomad.Client) AppController {
	return nil
	// return &appcontroller{
	// 	nc: nc,
	// }
}
