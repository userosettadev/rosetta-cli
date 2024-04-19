package token

import (
	"log/slog"

	"github.com/pkoukk/tiktoken-go"
	tiktoken_loader "github.com/pkoukk/tiktoken-go-loader"
	pb "github.com/userosettadev/rosetta-cli/api"
)

func CountMultipleFiles(files []*pb.File) int {

	res := 0
	for _, currFile := range files {
		res += Count(string(currFile.Content))
	}

	return res
}

func Count(text string) int {

	tiktoken.SetBpeLoader(tiktoken_loader.NewOfflineLoader())
	tke, err := tiktoken.GetEncoding("cl100k_base")
	if err != nil {
		slog.Error("failed to load encoding, using estimation", "error", err)
		return len([]rune(text)) / 4
	}

	return len(tke.Encode(text, nil, nil))
}
