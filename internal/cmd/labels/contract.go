package labels

import (
	"context"

	"go.octolab.org/toolkit/github/internal"
)

type Provider interface {
	RepositoryWithLabels(context.Context, ...internal.RepositoryURN) ([]internal.Repository, error)
}
