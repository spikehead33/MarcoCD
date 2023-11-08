package applications

type RepoMonitor interface {
	Update() error
}

type repomonitor struct {
}
