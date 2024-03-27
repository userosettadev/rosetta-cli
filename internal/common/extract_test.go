package common_test

import (
	"testing"

	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/stretchr/testify/require"
)

func TestGetLanguage_Golang(t *testing.T) {

	lang, err := common.GetLanguage("golang")
	require.NoError(t, err)
	require.Equal(t, lang, "go")
}

func TestGetLanguage_Invalid(t *testing.T) {

	_, err := common.GetLanguage("golang-ai")
	require.Error(t, err)
}
