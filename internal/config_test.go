package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal"
)

func TestGetCommandConfig(t *testing.T) {

	require.NotNil(t, internal.GetCommandConfig())
}
