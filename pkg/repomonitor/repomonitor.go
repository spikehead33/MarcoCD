package repomonitor

type RepoMonitor interface {
	PullRepo() error
	RenderTemplate() error
}
