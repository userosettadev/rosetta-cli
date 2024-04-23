package internal_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	pb "github.com/userosettadev/rosetta-cli/api"
	"github.com/userosettadev/rosetta-cli/internal"
	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/userosettadev/rosetta-cli/internal/env"
)

func TestGetCommandGenerateOAS(t *testing.T) {

	require.NotNil(t, internal.GetCommandGenerateOAS())
}

func TestGenerateOAS(t *testing.T) {

	require.NoError(t, os.Setenv(env.EnvKeyApiKey, "test-api-key"))
	require.NoError(t, internal.GenerateOAS(".", common.LangGo, "", false,
		func(string, []*pb.File, string, []byte) (string, error) {
			return "", nil
		}))
}
