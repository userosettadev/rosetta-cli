package internal

import (
	"github.com/spf13/cobra"
	"github.com/userosettadev/rosetta-cli/internal/config"
)

func GetCommandConfig() *cobra.Command {

	return &cobra.Command{
		Use:   "config",
		Short: "Create a configuration file",
		Long:  "Create a configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.CreateDefaultConfigFile(); err != nil {
				return err
			}
			cmd.Printf("%s created :)", config.ConfigFilename)

			return nil
		},
	}
}
