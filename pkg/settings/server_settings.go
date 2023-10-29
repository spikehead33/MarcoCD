package settings

type ServerSettings struct {
	GitHubAppID             int64
	GitHubAppInstallationID int64
	GitHubAppPrivateKeyPath string

	GitOpsRepoPath string
}
