package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal"
	"github.com/userosettadev/rosetta-cli/internal/common"
)

func TestCountTokens(t *testing.T) {

	count, err := internal.CountTokens(".", common.LangGo, true)
	require.NoError(t, err)
	require.True(t, count > 0)
}
