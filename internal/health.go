package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	pb "github.com/userosettadev/rosetta-cli/api"
	"github.com/userosettadev/rosetta-cli/internal/common"
)

func GetCommandHealth() *cobra.Command {

	cmd := cobra.Command{
		Use:   "health",
		Short: "Rosetta service health",
		Long:  "Rosetta service health",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Root().SilenceUsage = true

			ok, err := health()
			if err != nil {
				return err
			}
			cmd.Println(ok)

			return nil
		},
	}

	return &cmd
}

func health() (string, error) {

	conn, err := common.BuildGRPCConnection()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	response, err := pb.NewFileServiceClient(conn).Check(context.Background(), &pb.HealthCheckRequest{Service: "rosetta-cli"})
	if err != nil {
		if strings.Contains(err.Error(), "connect: connection refused") {
			return "", fmt.Errorf(`failed to connect to Rosetta service. This could be due to several reasons, such as network connectivity issues, a firewall or proxy settings blocking the connection`)
		}
		return "", err
	}

	return response.GetStatus().String(), nil
}
