package common

import (
	"github.com/pkoukk/tiktoken-go"
	pb "github.com/userosettadev/rosetta-cli/api"
)

func CountTokensMultipleText(files []*pb.File) (int, error) {

	res := 0
	for _, currFile := range files {
		currTokens, err := CountTokens(string(currFile.Content))
		if err != nil {
			return -1, err
		}
		res += currTokens
	}

	return res, nil
}

func CountTokens(text string) (int, error) {

	const encoding = "cl100k_base"
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		return -1, err
	}

	return len(tke.Encode(text, nil, nil)), nil
}
