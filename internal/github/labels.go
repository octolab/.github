package github

import (
	"context"

	"go.octolab.org/pointer"

	"go.octolab.org/toolkit/github/internal"
)

func (manager *manager) RepositoryWithLabels(
	ctx context.Context,
	list ...internal.FullName,
) ([]internal.Repository, error) {
	result := make([]internal.Repository, 0, len(list))
	for _, fullName := range list {
		r, _, err := manager.client.Repositories.Get(ctx, fullName.Owner(), fullName.Name())
		if err != nil {
			return nil, err
		}
		result = append(result, internal.Repository{
			ID:       pointer.ValueOfInt64(r.ID),
			FullName: internal.FullName(pointer.ValueOfString(r.FullName)),
		})
	}
	return result, nil
}
