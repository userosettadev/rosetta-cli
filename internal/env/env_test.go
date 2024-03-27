package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal/env"
)

func TestEnv_GetKey(t *testing.T) {

	const value = "test-me"
	require.NoError(t, os.Setenv(env.EnvTenantKey, value))
	require.Equal(t, value, env.GetInstance().GetTenant())
}
