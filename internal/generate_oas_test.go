package internal_test

import (
	"os"
	"testing"

	"github.com/userosettadev/rosetta-cli/internal"
	pb "github.com/userosettadev/rosetta-cli/internal/api"
	"github.com/userosettadev/rosetta-cli/internal/env"
	"github.com/stretchr/testify/require"
)

func TestGenerateOAS(t *testing.T) {

	require.NoError(t, os.Setenv(env.EnvTenantKey, "test-tenant"))
	require.NoError(t, internal.GenerateOAS(".", "go", "", false,
		func(string, []*pb.File, string, []byte) (string, error) {
			return "", nil
		}))
}
