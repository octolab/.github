package main

import (
	"os"

	"go.octolab.org/toolkit/cli/cobra"

	"go.octolab.org/toolkit/github/internal/cmd"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func main() {
	root := cmd.New(os.Getenv("GITHUB_TOKEN"))
	root.SetOut(os.Stdout)
	root.SetErr(os.Stderr)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(version, date, commit))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
