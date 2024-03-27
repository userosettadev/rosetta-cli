package internal

import (
	"errors"
	"fmt"

	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/spf13/cobra"
)

func GetCommandCountToken() *cobra.Command {

	var lang string
	var verbose bool

	cmd := cobra.Command{
		Use:   "count [src] [flags]",
		Short: "Count the number of tokens",
		Long:  "Count the number of tokens based on the provided source code. 'src' should be the path to the root directory of the source code",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if lang == "" {
				return errors.New("please set programming language")
			}
			// by now flags have been parsed successfully, so we don't need to show usage on any errors
			cmd.Root().SilenceUsage = true

			lang, err := common.GetLanguage(lang)
			if err != nil {
				return err
			}

			code, err := common.ExtractCode(args[0], lang, verbose)
			if err != nil {
				return err
			}
			tokens, err := common.CountTokensMultipleText(code)
			if err != nil {
				return err
			}
			fmt.Println(tokens)

			return nil
		},
	}
	cmd.Flags().StringVarP(&lang, "lang", "l", "", "Programming language")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	return &cmd
}
