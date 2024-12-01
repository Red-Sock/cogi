package cogi

import (
	"context"
	"net/http"
	"path"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/google/go-github/v67/github"
	"golang.org/x/sync/errgroup"
	"gopkg.in/src-d/go-git.v4"
)

type githubManager struct {
	cl *github.Client
}

func newGithub() githubManager {
	return githubManager{
		cl: github.NewClient(nil),
	}
}

func (g *githubManager) Clone(tempDirPath string, origin Origin) (*CloneResults, error) {
	ctx := context.Background()
	repos, _, err := g.cl.Repositories.ListByOrg(ctx, origin.Dir, nil)
	if err != nil {
		if !isNotFoundInGithub(err) {
			return nil, errors.Wrap(err, "error listing repos by org")
		}
	}

	if len(repos) == 0 {
		repos, _, err = g.cl.Repositories.ListByUser(ctx, origin.Dir, nil)
		if err != nil {
			return nil, errors.Wrap(err, "error listing repos by user")
		}
	}

	cr := &CloneResults{}

	eg := errgroup.Group{}
	for _, repo := range repos {
		folderToCloneTo := path.Join(tempDirPath, *repo.Name)
		repo := repo
		eg.Go(func() error {
			err = g.cloneRepo(folderToCloneTo, repo)
			if err != nil {
				cr.AddFail(repo.GetCloneURL(), err)
				return nil
			}

			cr.AddSuccess(repo.GetCloneURL())
			return nil
		})
	}

	_ = eg.Wait()

	return cr, nil
}

func (g *githubManager) cloneRepo(toFs string, repo *github.Repository) error {
	cloneOpts := &git.CloneOptions{
		URL: repo.GetCloneURL(),
	}

	_, err := git.PlainClone(toFs, false, cloneOpts)
	if err != nil {
		return errors.Wrap(err, "error when tried to clone\""+repo.GetCloneURL()+"\"repo")
	}

	return nil
}

func isNotFoundInGithub(err error) bool {
	ghErr, ok := err.(*github.ErrorResponse)
	if !ok {
		return false
	}

	return ghErr.Response.StatusCode == http.StatusNotFound
}
