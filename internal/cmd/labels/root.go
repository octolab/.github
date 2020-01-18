package labels

import "github.com/spf13/cobra"

func New(provider Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "labels",
		Short: "Manage labels",
		Long:  "Manage labels.",
	}
	cmd.AddCommand(NewListCommand(provider))
	return cmd
}
