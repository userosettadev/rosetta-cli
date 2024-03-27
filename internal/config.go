package internal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/userosettadev/rosetta-cli/internal/config"
)

func GetCommandConfig() *cobra.Command {

	cmd := cobra.Command{
		Use:   "config",
		Short: "Create a configuration file",
		Long:  "Create a configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.CreateDefaultConfigFile(); err != nil {
				return err
			}
			fmt.Printf("%s created :)", config.ConfigFilename)

			return nil
		},
	}

	return &cmd
}
