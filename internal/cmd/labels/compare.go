package labels

import (
	"context"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"

	"go.octolab.org/toolkit/github/internal"
)

func NewCompareCommand(provider Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "compare",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			list := internal.RepositoryURNFromStrings(args)
			repositories, err := provider.RepositoryWithLabels(ctx, list...)
			if err != nil {
				return err
			}

			expected, obtained := repositories[0], repositories[1]

			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Text: "Color", Align: simpletable.AlignRight},
					{Text: "Name", Align: simpletable.AlignRight},
					{Text: "<->", Align: simpletable.AlignCenter},
					{Text: "Name"},
					{Text: "Color"},
				},
			}
			delta := len(expected.Labels) - len(obtained.Labels)
			switch {
			case delta > 0:
				obtained.Labels = append(obtained.Labels, make([]internal.Label, delta)...)
			case delta < 0:
				expected.Labels = append(expected.Labels, make([]internal.Label, -delta)...)
			}
			for i := range make([]struct{}, len(expected.Labels)) {
				left, right := expected.Labels[i], obtained.Labels[i]
				table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
					{Text: left.Color},
					{Text: left.Name},
					{},
					{Text: right.Name},
					{Text: right.Color},
				})
			}
			table.SetStyle(simpletable.StyleDefault)
			cmd.Println(table.String())
			return nil
		},
	}
	return cmd
}
