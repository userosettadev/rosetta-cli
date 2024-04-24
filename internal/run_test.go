package internal_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal"
)

func TestRun_CountCommand(t *testing.T) {

	require.Equal(t, 0, internal.Run(cmdToArgs("rosetta count . -l go"), io.Discard, io.Discard))
}

func cmdToArgs(cmd string) []string {

	return strings.Fields(cmd)
}
