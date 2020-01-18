package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	"go.octolab.org/toolkit/github/internal/cmd/labels"
	"go.octolab.org/toolkit/github/internal/github"
)

// New returns the new root command.
func New(token string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "github",
		Short:         "GitHub manager",
		Long:          "GitHub manager for all OctoLab's projects.",
		SilenceErrors: false,
		SilenceUsage:  true,
	}
	source := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client := oauth2.NewClient(context.TODO(), source)
	cmd.AddCommand(labels.New(github.New(client)))
	return cmd
}
