package github

import (
	"context"
	"sync"

	"github.com/google/go-github/v29/github"
	"golang.org/x/sync/errgroup"

	"go.octolab.org/toolkit/github/internal"
	"go.octolab.org/toolkit/github/internal/errors"
)

func (manager *manager) RepositoryWithLabels(
	ctx context.Context,
	list ...internal.RepositoryURN,
) ([]internal.Repository, error) {
	result := make([]internal.Repository, 0, len(list))

	for _, urn := range list {
		var (
			mutex      sync.Mutex
			repository internal.Repository
		)

		owner, repo := urn.Split()
		group, ctx := errgroup.WithContext(ctx)

		group.Go(func() error {
			data, _, err := manager.client.Repositories.Get(ctx, owner, repo)
			if err != nil {
				return err
			}

			mutex.Lock()
			repository.ID = data.GetID()
			repository.URN = internal.RepositoryURN(data.GetFullName())
			mutex.Unlock()
			return nil
		})

		group.Go(func() error {
			opt := &github.ListOptions{}
			data, _, err := manager.client.Issues.ListLabels(ctx, owner, repo, opt)
			if err != nil {
				return err
			}
			if len(data) == 0 {
				return nil
			}

			labels := make([]internal.Label, 0, len(data))
			for _, label := range data {
				labels = append(labels, internal.Label{
					ID:    label.GetID(),
					Name:  label.GetName(),
					Color: label.GetColor(),
					Desc:  label.GetDescription(),
				})
			}

			mutex.Lock()
			repository.Labels = labels
			mutex.Unlock()
			return nil
		})

		if err := group.Wait(); err != nil {
			return nil, err
		}
		result = append(result, repository)
	}

	if len(result) != len(list) {
		return nil, errors.Inconsistent
	}
	return result, nil
}
