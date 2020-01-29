package github

import (
	"context"
	"sort"

	"go.octolab.org/toolkit/github/internal"
)

func (manager *manager) CompareLabels(
	ctx context.Context,
	expected,
	obtained internal.RepositoryURN,
) error {
	repositories, err := manager.RepositoryWithLabels(ctx, expected, obtained)
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		sort.Slice(repository.Labels, func(i, j int) bool {
			return repository.Labels[i].Name < repository.Labels[j].Name
		})
	}

	src, dst := repositories[0], repositories[1]
	_, _ = src, dst
	return nil
}

var conversion = map[string]struct {
	color string
	desc  string
	to    internal.Label
}{
	"bug": {
		color: "d73a4a",
		desc:  "Something isn't working",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"documentation": {
		color: "0075ca",
		desc:  "Improvements or additions to documentation",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"duplicate": {
		color: "cfd3d7",
		desc:  "This issue or pull request already exists",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"enhancement": {
		color: "a2eeef",
		desc:  "New feature or request",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"good first issue": {
		color: "7057ff",
		desc:  "Good for newcomers",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"help wanted": {
		color: "008672",
		desc:  "Extra attention is needed",
		to: internal.Label{
			Name:  "invalid",
			Color: "e4e669",
			Desc:  "This doesn't seem right",
		},
	},
	"invalid": {
		color: "e4e669",
		desc:  "This doesn't seem right",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"question": {
		color: "d876e3",
		desc:  "Further information is requested",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
	"wontfix": {
		color: "ffffff",
		desc:  "This will not be worked on",
		to: internal.Label{
			Name:  "",
			Color: "",
			Desc:  "",
		},
	},
}
