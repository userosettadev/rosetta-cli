package config_test

import (
	_ "embed"
	"os"
	"testing"
	"time"

	"github.com/userosettadev/rosetta-cli/internal/common"
	"github.com/userosettadev/rosetta-cli/internal/config"
	"github.com/stretchr/testify/require"
)

type MockFileInfo struct {
	FileName    string
	IsDirectory bool
}

func (mfi MockFileInfo) Name() string       { return mfi.FileName }
func (mfi MockFileInfo) Size() int64        { return int64(8) }
func (mfi MockFileInfo) Mode() os.FileMode  { return os.ModePerm }
func (mfi MockFileInfo) ModTime() time.Time { return time.Now() }
func (mfi MockFileInfo) IsDir() bool        { return mfi.IsDirectory }
func (mfi MockFileInfo) Sys() interface{}   { return nil }

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
		MockFileInfo{
			FileName:    "test_advertiser.py",
			IsDirectory: false,
		})
	require.NoError(t, err)
	require.True(t, ignore)
}
