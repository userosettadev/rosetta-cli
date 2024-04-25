package internal_test

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal"
)

func TestRun_CountCommand(t *testing.T) {

	buf := new(bytes.Buffer)

	require.Equal(t, 0, internal.Run(cmdToArgs("rosetta count . -l go"), buf, io.Discard))

	tokens, err := strconv.Atoi(strings.Replace(buf.String(), "\n", "", -1))
	require.NoError(t, err)
	require.True(t, tokens > 0)
}

func TestRun_CountCommand_NoLang(t *testing.T) {

	require.Equal(t, 1, internal.Run(cmdToArgs("rosetta count ."), io.Discard, io.Discard))
}

func cmdToArgs(cmd string) []string {

	return strings.Fields(cmd)
}
