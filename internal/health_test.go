package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal"
)

func TestGetCommandHealth(t *testing.T) {

	require.NotNil(t, internal.GetCommandHealth())
}
