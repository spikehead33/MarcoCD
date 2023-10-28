package app_controller

import (
	"slices"

	nomad "github.com/hashicorp/nomad/api"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type AppController interface {
	Reconcile() error
	UpdateDesiredState(map[string]string) error
	UpdateModuleDependencies(map[string][]string)
}

type service struct {
	moduleDependecies map[string][]string
	desiredState      map[string]*nomad.Job
	nomadClient       *nomad.Client
}

func (s *service) UpdateModuleDependencies(newDep map[string][]string) {
	s.moduleDependecies = newDep
}

func (s *service) UpdateDesiredState(newDesiredState map[string]string) error {
	dep := make(map[string]*nomad.Job)

	jobHandler := s.nomadClient.Jobs()
	for jobName, jobSpec := range newDesiredState {
		job, err := jobHandler.ParseHCL(jobSpec, false)
		if err != nil {
			return err
		}

		dep[jobName] = job
	}

	s.desiredState = dep

	return nil
}

func (s *service) Reconcile() error {
	moduleExecutionPlan, err := s.formulateMouduleExecutionOrdering()
	if err != nil {
		return err
	}

	jobNameToModule := make(map[string]string)
	for jobName, jobSpec := range s.desiredState {
		jobNameToModule[jobName] = jobSpec.Meta["module"]
	}

	jobHandler := s.nomadClient.Jobs()
	for _, module := range moduleExecutionPlan {
		for _, job := range s.desiredState {
			if job.Meta["module"] != module {
				continue
			}

			_, _, err := jobHandler.Register(job, nil)
			if err != nil {
				return err
			}
		}
	}

	if err = s.clean(); err != nil {
		return err
	}

	return nil
}

func (s *service) formulateMouduleExecutionOrdering() ([]string, error) {
	g := simple.NewDirectedGraph()
	moduleToNodeID, nodeIDToModule := s.createMappingsBetweenModuleAndNodeID()

	// construct the node
	for module := range s.moduleDependecies {
		nodeID := moduleToNodeID[module]
		g.AddNode(simple.Node(nodeID))
	}

	// construct the edge
	for module, deps := range s.moduleDependecies {
		for _, dep := range deps {
			from := moduleToNodeID[module]
			to := moduleToNodeID[dep]
			g.HasEdgeFromTo(from, to)
		}
	}

	orderlist, err := topo.Sort(g)
	if err != nil {
		return nil, err
	}

	executionOrder := []string{}
	for _, node := range orderlist {
		nodeID := node.ID()
		module := nodeIDToModule[nodeID]

		executionOrder = append(executionOrder, module)
	}

	slices.Reverse(executionOrder)

	return executionOrder, nil
}

func (s *service) createMappingsBetweenModuleAndNodeID() (map[string]int64, map[int64]string) {
	moduleToNodeID := make(map[string]int64)
	nodeIDToModule := make(map[int64]string)

	var cnt int64 = 0
	for module := range s.moduleDependecies {
		moduleToNodeID[module] = cnt
		nodeIDToModule[cnt] = module
		cnt += 1
	}

	return moduleToNodeID, nodeIDToModule
}

// TODO: implement the cleanup function
// remove job where jobName/resource with Meta['managed-by'] == 'marcocd' that is not contained in desired state
func (s *service) clean() error {
	// jobHandler := s.nomadClient.Jobs()
	return nil
}
