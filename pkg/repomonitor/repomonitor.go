package repomonitor

type RepoMonitor interface {
	UpdateRepo() error
}

type repomonitor struct {
	githubAppID             int64
	githubAppInstallationID int64
	githubAppPrivateKeyPath string
}

func New(gID, giID int64, pkPath string) RepoMonitor {
	return &repomonitor{
		githubAppID:             gID,
		githubAppInstallationID: giID,
		githubAppPrivateKeyPath: pkPath,
	}
}

func (mon *repomonitor) UpdateRepo() error {
	return nil
}

func (mon *repomonitor) RenderTemplate() error {
	return nil
}
