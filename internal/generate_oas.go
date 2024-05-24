package internal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	pb "github.com/userosettadev/rosetta-cli/api"
	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/userosettadev/rosetta-cli/internal/env"
)

func GetCommandGenerateOAS() *cobra.Command {

	var lang string
	var specPath string
	var output string
	var verbose bool

	cmd := cobra.Command{
		Use:   "gen [path] [flags]",
		Short: "Generate OpenAPI Specification",
		Long:  "Generate OpenAPI Specification based on the provided source code. 'src' should be the path to the root directory of the source code",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if lang == "" {
				return errors.New("please set programming language")
			}
			// by now flags have been parsed successfully, so we don't need to show usage on any errors
			cmd.Root().SilenceUsage = true

			spec, err := GenerateOAS(args[0], lang, specPath, verbose, generate)
			if err != nil {
				return err
			}

			if output != "" {
				return os.WriteFile(output, []byte(spec), 0644)
			}
			cmd.Println(spec)

			return nil
		},
	}
	cmd.Flags().StringVarP(&lang, "lang", "l", "", "Programming language")
	cmd.Flags().StringVarP(&specPath, "spec", "s", "", "Path to old OpenAPI spec")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Write the resulting OpenAPI specification to the specified output file instead of the console. If the specified output file already exists, it will be overwritten")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	return &cmd
}

func GenerateOAS(root string, lang string, specPath string, verbose bool,
	convert func(string, []*pb.File, string, []byte) (string, error)) (string, error) {

	apiKey := env.GetInstance().GetApiKey()
	if apiKey == "" || len(apiKey) > 40 {
		return "", fmt.Errorf("please set %s environment variable", env.EnvKeyApiKey)
	}

	lang, err := common.GetLanguage(lang)
	if err != nil {
		return "", err
	}

	code, err := common.ExtractCode(root, lang, verbose)
	if err != nil {
		return "", err
	}

	var spec []byte
	if specPath != "" && specPath != "na" {
		spec, err = common.ReadFile(specPath, verbose)
		if err != nil {
			return "", err
		}
	}

	return convert(apiKey, code, lang, spec)
}

func generate(apiKey string, files []*pb.File, lang string, spec []byte) (string, error) {

	conn, err := common.BuildGRPCConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	response, err := pb.NewFileServiceClient(conn).
		CreateOAS(context.Background(), &pb.CreateOASRequest{
			ApiKey:   apiKey,
			Language: lang,
			Files:    files,
			Spec:     spec,
		})
	if err != nil {
		if strings.Contains(err.Error(), "connect: connection refused") {
			return "", fmt.Errorf(`failed to connect to Rosetta service. This could be due to several reasons, such as network connectivity issues, a firewall or proxy settings blocking the connection`)
		}
		return "", err
	}
	if response.Error != "" {
		return "", errors.New(response.Error)
	}

	return response.Spec, nil
}
