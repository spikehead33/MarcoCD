package applications

type Pruner interface {
	Prune() error
}

type prune struct {
}
