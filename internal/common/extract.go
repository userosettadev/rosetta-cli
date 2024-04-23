package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pb "github.com/userosettadev/rosetta-cli/api"
	"github.com/userosettadev/rosetta-cli/internal/config"
)

const (
	LangGo         = "go"
	LangJavaScript = "javascript"
)

func GetLanguage(input string) (string, error) {

	res := strings.ToLower(input)
	if res == "golang" || res == LangGo {
		return LangGo, nil
	} else if res == "python" {
		return "python", nil
	} else if res == "java" {
		return "java", nil
	} else if res == "js" || res == LangJavaScript {
		return LangJavaScript, nil
	} else if res == "ts" || res == "typescript" {
		return "typescript", nil
	} else if res == "csharp" || res == "cs" {
		return "csharp", nil
	} else if res == "kotlin" {
		return "kotlin", nil
	} else if res == "scala" {
		return "scala", nil
	} else if res == "rust" {
		return "rust", nil
	}

	return "", fmt.Errorf("%s programming language is currently not supported", input)
}

func ExtractCode(root string, lang string, verbose bool) ([]*pb.File, error) {

	filter, err := config.GetInstance().BuildFilter(lang)
	if err != nil {
		return nil, err
	}

	code, err := visit(root, filter, verbose)
	if err != nil {
		return nil, err
	}
	if len(code) == 0 {
		return nil, fmt.Errorf("could not find %s files in %s", lang, root)
	}

	return code, nil
}

func visit(root string, filter *config.Filter, verbose bool) ([]*pb.File, error) {

	fileInfo, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("invalid path %s", root)
	}

	if fileInfo.IsDir() {
		return visitDir(root, filter, verbose)
	}

	content, err := ReadFile(root, verbose)
	if err != nil {
		return nil, err
	}
	return []*pb.File{{Path: root, Content: content}}, nil
}

func visitDir(root string, filter *config.Filter, verbose bool) ([]*pb.File, error) {

	var res []*pb.File
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to read %s", path)
		}

		if ok, err := filter.Ignore(path, info); ok {
			return err
		}

		content, err := ReadFile(path, verbose)
		if err != nil {
			return err
		}
		res = append(res, &pb.File{Path: path, Content: content})

		return nil
	})

	return res, err
}

func ReadFile(path string, verbose bool) ([]byte, error) {

	if verbose {
		fmt.Println(path)
	}

	res, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s", path)
	}

	return res, nil
}
