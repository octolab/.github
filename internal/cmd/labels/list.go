package labels

import (
	"context"

	_ "github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"

	"go.octolab.org/toolkit/github/internal"
)

func NewListCommand(provider Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			repositories, err := provider.RepositoryWithLabels(ctx, convert(args)...)
			if err != nil {
				return err
			}

			for _, repository := range repositories {
				cmd.Printf("%s:\n\n---\n\n", repository.FullName)
			}
			return nil
		},
	}
	return cmd
}

func convert(list []string) []internal.FullName {
	converted := make([]internal.FullName, 0, len(list))
	for _, str := range list {
		converted = append(converted, internal.FullName(str))
	}
	return converted
}
