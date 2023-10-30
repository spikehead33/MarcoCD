package gitauth

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/go-git/go-git/plumbing/transport"
	git_http "github.com/go-git/go-git/v5/plumbing/transport/http"
)

type GitAuthenticator struct {
	tr *ghinstallation.Transport
}

func New(githubAppID, githubAppInstallationID int64, githubAppPrivateKey []byte) (*GitAuthenticator, error) {
	tr, err := ghinstallation.New(http.DefaultTransport, githubAppID, githubAppInstallationID, githubAppPrivateKey)
	if err != nil {
		return nil, err
	}

	return &GitAuthenticator{tr}, nil
}

func (g *GitAuthenticator) GetGitAuth(ctx context.Context) (transport.AuthMethod, error) {
	token, err := g.tr.Token(ctx)
	if err != nil {
		return nil, nil
	}
	return &git_http.BasicAuth{Username: "x-access-token", Password: token}, nil
}
