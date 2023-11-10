package applications

import (
	"errors"
	"fmt"
	"marcocd/pkg/domains"
	"slices"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type Planner interface {
	Plan() (PlanningResult, error)
}

type planner struct {
	dependencies domains.Dependencies
}

type PlanningResult struct {
	executionOrder []string
}

func (p *PlanningResult) ExecutionOrder() []string {
	return p.executionOrder
}

func (p *planner) Plan() (PlanningResult, error) {
	return plan(p.dependencies)
}

func plan(dependencies domains.Dependencies) (PlanningResult, error) {
	table := []string{}

	i := 0
	for module := range dependencies {
		table[i] = module
		i++
	}

	depGraph, err := buildGraph(table, dependencies)
	if err != nil {
		return PlanningResult{}, err
	}

	return TopologicalSort(depGraph, table)
}

func buildGraph(table []string, dependencies domains.Dependencies) (graph.Directed, error) {
	g := simple.NewDirectedGraph()
	for i := range table {
		g.AddNode(simple.Node(i))
	}

	for module, deps := range dependencies {
		from := slices.Index(table, module)
		for _, dep := range deps {
			to := slices.Index(table, dep)
			if to == -1 {
				return nil, errors.New(fmt.Sprintf("dependency %v in module %v not found", dep, module))
			}

			g.HasEdgeFromTo(int64(from), int64(to))
		}
	}

	return g, nil
}

func TopologicalSort(g graph.Directed, table []string) (PlanningResult, error) {
	sorted, err := topo.Sort(g)
	if err != nil {
		return PlanningResult{}, err
	}

	result := []string{}
	for _, node := range sorted {
		module := table[node.ID()]
		result = append(result, module)
	}

	return PlanningResult{
		executionOrder: result,
	}, nil
}
