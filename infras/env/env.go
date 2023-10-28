package env

import (
	"github.com/caarlos0/env/v9"
)

type Env struct {
	GitHubAppID             int64  `env:"GITHUB_APP_ID"`
	GitHubAppInstallationID int64  `env:"GITHUB_APP_INSTALLATION_ID"`
	GitHubAppPrivateKeyPath string `env:"GITHUB_APP_PRIVATE_KEY_PATH"`

	GitOpsRepositoryUri string

	NomadClusterAddr string `env:"NOMAD_ADDR"`
}

func New() (Env, error) {
	enviornment := Env{}
	if err := env.Parse(enviornment); err != nil {
		return enviornment, err
	}

	return enviornment, nil
}
