package appcontroller

// import (
// 	"slices"

// 	"gonum.org/v1/gonum/graph/simple"
// 	"gonum.org/v1/gonum/graph/topo"
// )

// func (controller *appcontroller) Reconcile() error {
// 	executionOrderPlan, err := controller.getExecutionOrderPlan()
// 	if err != nil {
// 		return err
// 	}

// 	jobHandler := controller.nc.Jobs()

// 	// for _, module := range executionOrderPlan {
// 	// 	for _, job := range controller.desiredState {
// 	// 		if job.Meta["module"] != module {
// 	// 			continue
// 	// 		}

// 	// 		_, _, err := jobHandler.Register(job, nil)
// 	// 		if err != nil {
// 	// 			return err
// 	// 		}
// 	// 	}
// 	// }

// 	// if err = c.clean(); err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }

// func (controller *appcontroller) getExecutionOrderPlan() ([]Module, error) {
// 	g := simple.NewDirectedGraph()

// 	moduleToNodeID, nodeIDToModule := controller.createMappingsBetweenModuleAndNodeID()

// 	// construct the node
// 	for module := range controller.moduleDependecies {
// 		nodeID := moduleToNodeID[module]
// 		g.AddNode(simple.Node(nodeID))
// 	}

// 	// construct the edge
// 	for module, deps := range controller.moduleDependecies {
// 		for _, dep := range deps {
// 			from := moduleToNodeID[module]
// 			to := moduleToNodeID[dep]
// 			g.HasEdgeFromTo(from, to)
// 		}
// 	}

// 	orderlist, err := topo.Sort(g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	executionOrder := []string{}
// 	for _, node := range orderlist {
// 		nodeID := node.ID()
// 		module := nodeIDToModule[nodeID]

// 		executionOrder = append(executionOrder, module)
// 	}

// 	slices.Reverse(executionOrder)

// 	return executionOrder, nil
// }

// func (c *appcontroller) createMappingsBetweenModuleAndNodeID() (map[string]int64, map[int64]string) {
// 	moduleToNodeID := make(map[string]int64)
// 	nodeIDToModule := make(map[int64]string)

// 	var cnt int64 = 0
// 	for module := range c.moduleDependecies {
// 		moduleToNodeID[module] = cnt
// 		nodeIDToModule[cnt] = module
// 		cnt += 1
// 	}

// 	return moduleToNodeID, nodeIDToModule
// }

// // TODO: implement the cleanup function
// // remove job where jobName/resource with Meta['managed-by'] == 'marcocd' that is not contained in desired state
// func (c *appcontroller) clean() error {
// 	// jobHandler := s.nomadClient.Jobs()
// 	return nil
// }
