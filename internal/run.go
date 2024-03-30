package internal

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/userosettadev/rosetta-cli/build"
)

func Run(args []string, stdout io.Writer, stderr io.Writer) int {

	rootCmd := &cobra.Command{
		Use:   "rosetta",
		Short: "Code to OpenAPI Specification",
	}

	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	rootCmd.Version = build.Version

	rootCmd.AddCommand(GetCommandCountTokens(), GetCommandGenerateOAS(), GetCommandConfig())

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
