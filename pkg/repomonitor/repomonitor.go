package repomonitor

// import (
// 	"context"
// 	"marcocd/pkg/gitauth"
// 	"os"

// 	"github.com/go-git/go-git/v5"
// 	"github.com/go-git/go-git/v5/plumbing"
// )

// type RepoMonitor interface {
// 	UpdateRepo(context.Context) error
// }

// type repomonitor struct {
// 	repoPath      string
// 	refName       string
// 	authenticator *gitauth.GitAuthenticator
// 	workTree      *git.Worktree
// }

// func New(gID, giID int64, pkPath string) (RepoMonitor, error) {
// 	key, err := os.ReadFile(pkPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	authenticator, err := gitauth.New(gID, giID, key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &repomonitor{authenticator: authenticator}, nil
// }

// func (mon *repomonitor) UpdateRepo(ctx context.Context) error {
// 	auth, err := mon.authenticator.GetGitAuth(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	if mon.workTree == nil {
// 		repo, err := git.PlainClone(mon.repoPath, false, &git.CloneOptions{
// 			Auth:          auth,
// 			SingleBranch:  true,
// 			ReferenceName: plumbing.ReferenceName(mon.refName),
// 		})

// 		if err != nil {
// 			return err
// 		}

// 		workTree, err := repo.Worktree()
// 		if err != nil {
// 			return err
// 		}

// 		mon.workTree = workTree

// 		return nil
// 	}

// 	mon.workTree.Pull()

// 	return nil
// }

// func (mon *repomonitor) RenderTemplate() error {
// 	return nil
// }
