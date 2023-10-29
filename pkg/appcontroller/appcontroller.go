package appcontroller

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

type appcontroller struct {
	moduleDependecies map[string][]string
	desiredState      map[string]*nomad.Job
	nomadClient       *nomad.Client
}

func New(client *nomad.Client) AppController {
	return &appcontroller{
		nomadClient: client,
	}
}

func (c *appcontroller) UpdateModuleDependencies(newDep map[string][]string) {
	c.moduleDependecies = newDep
}

func (c *appcontroller) UpdateDesiredState(newDesiredState map[string]string) error {
	dep := make(map[string]*nomad.Job)

	jobHandler := c.nomadClient.Jobs()
	for jobName, jobSpec := range newDesiredState {
		job, err := jobHandler.ParseHCL(jobSpec, false)
		if err != nil {
			return err
		}

		dep[jobName] = job
	}

	c.desiredState = dep

	return nil
}

func (c *appcontroller) Reconcile() error {
	moduleExecutionPlan, err := c.formulateMouduleExecutionOrdering()
	if err != nil {
		return err
	}

	jobNameToModule := make(map[string]string)
	for jobName, jobSpec := range c.desiredState {
		jobNameToModule[jobName] = jobSpec.Meta["module"]
	}

	jobHandler := c.nomadClient.Jobs()
	for _, module := range moduleExecutionPlan {
		for _, job := range c.desiredState {
			if job.Meta["module"] != module {
				continue
			}

			_, _, err := jobHandler.Register(job, nil)
			if err != nil {
				return err
			}
		}
	}

	if err = c.clean(); err != nil {
		return err
	}

	return nil
}

func (c *appcontroller) formulateMouduleExecutionOrdering() ([]string, error) {
	g := simple.NewDirectedGraph()
	moduleToNodeID, nodeIDToModule := c.createMappingsBetweenModuleAndNodeID()

	// construct the node
	for module := range c.moduleDependecies {
		nodeID := moduleToNodeID[module]
		g.AddNode(simple.Node(nodeID))
	}

	// construct the edge
	for module, deps := range c.moduleDependecies {
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

func (c *appcontroller) createMappingsBetweenModuleAndNodeID() (map[string]int64, map[int64]string) {
	moduleToNodeID := make(map[string]int64)
	nodeIDToModule := make(map[int64]string)

	var cnt int64 = 0
	for module := range c.moduleDependecies {
		moduleToNodeID[module] = cnt
		nodeIDToModule[cnt] = module
		cnt += 1
	}

	return moduleToNodeID, nodeIDToModule
}

// TODO: implement the cleanup function
// remove job where jobName/resource with Meta['managed-by'] == 'marcocd' that is not contained in desired state
func (c *appcontroller) clean() error {
	// jobHandler := s.nomadClient.Jobs()
	return nil
}
