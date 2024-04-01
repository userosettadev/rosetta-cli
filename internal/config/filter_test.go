package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/userosettadev/rosetta-cli/internal/config"
)

type mockFileInfo struct {
	FileName    string
	IsDirectory bool
}

func (m mockFileInfo) Name() string       { return m.FileName }
func (m mockFileInfo) Size() int64        { return int64(8) }
func (m mockFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (m mockFileInfo) ModTime() time.Time { return time.Now() }
func (m mockFileInfo) IsDir() bool        { return m.IsDirectory }
func (m mockFileInfo) Sys() interface{}   { return nil }

func TestBuildFilter(t *testing.T) {

	lang, err := common.GetLanguage("js")
	require.NoError(t, err)
	filter, err := config.GetInstance().BuildFilter(lang)
	require.NoError(t, err)
	require.Equal(t, ".js", filter.Extension)
}

func TestFilter_Ignore(t *testing.T) {

	lang, err := common.GetLanguage("python")
	require.NoError(t, err)
	filter, err := config.GetInstance().BuildFilter(lang)
	require.NoError(t, err)
	ignore, err := filter.Ignore("cpa-network/api/tests/test_advertiser.py",
		mockFileInfo{
			FileName:    "test_advertiser.py",
			IsDirectory: false,
		})
	require.NoError(t, err)
	require.True(t, ignore)
}
