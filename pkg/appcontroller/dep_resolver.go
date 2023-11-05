package appcontroller

type Dependable interface {
	Dependencies() []Dependable
}

type DependenciesResolver interface {
	Resolve() ([]interface{}, error)
}
