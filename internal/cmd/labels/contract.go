package labels

import (
	"context"

	"go.octolab.org/toolkit/github/internal/entity"
)

type Provider interface {
	RepositoryWithLabels(context.Context, ...entity.RepositoryURN) ([]entity.Repository, error)
}
