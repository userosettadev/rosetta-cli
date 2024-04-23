package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal/common"
)

func TestGetLanguage_Golang(t *testing.T) {

	lang, err := common.GetLanguage("golang")
	require.NoError(t, err)
	require.Equal(t, lang, common.LangGo)
}

func TestGetLanguage_Invalid(t *testing.T) {

	_, err := common.GetLanguage("golang-ai")
	require.Error(t, err)
}

func TestExtractCode(t *testing.T) {

	_, err := common.ExtractCode(".", common.LangGo, true)
	require.NoError(t, err)
}

func TestExtractCode_Empty(t *testing.T) {

	_, err := common.ExtractCode(".", common.LangJavaScript, true)
	require.Error(t, err)
}
