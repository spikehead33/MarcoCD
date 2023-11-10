package applications

import "github.com/go-git/go-billy/v5/memfs"

type RepoMonitor interface {
	Update() error
}

type repomonitor struct {
	storer *memfs.Memory
}

func (repo *repomonitor) Update() error {
	return nil
}
