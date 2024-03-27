package config

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

const ConfigFilename = "rosetta.yaml"

type container struct {
	filters map[string]*Filter
}

var (
	once     sync.Once
	instance *container
)

//go:embed filter.yaml
var defaultFiltersContent string

type Filter struct {
	IgnoreFileNames        map[string]bool `yaml:"ignoreFileNames"`
	IgnoreFileNameContains string          `yaml:"ignoreFileNameContains"`
	IgnoreDirs             map[string]bool `yaml:"ignoreDirs"`
	IgnoreFileNamePrefixes []string        `yaml:"ignoreFileNamePrefixes"`
	IgnoreFileNameSuffixes []string        `yaml:"ignoreFileNameSuffixes"`
	Extension              string          `yaml:"extension"`
}

type ConfigFilter struct {
	IgnoreFileNames        []string `yaml:"ignoreFileNames"`
	IgnoreFileNameContains string   `yaml:"ignoreFileNameContains"`
	IgnoreDirs             []string `yaml:"ignoreDirs"`
	IgnoreFileNamePrefixes []string `yaml:"ignoreFileNamePrefixes"`
	IgnoreFileNameSuffixes []string `yaml:"ignoreFileNameSuffixes"`
	Extension              string   `yaml:"extension"`
}

func GetInstance() *container {

	once.Do(func() {
		filters, err := load()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
		}
		instance = &container{filters: filters}
	})

	return instance
}

func load() (map[string]*Filter, error) {

	content, err := getConfigFilterContent()
	if err != nil {
		return nil, err
	}

	configFilters := make(map[string]*ConfigFilter)
	if err := yaml.Unmarshal(content, &configFilters); err != nil {
		return nil, errors.New("invalid config file format")
	}
	defaultFiltersContent = ""

	res := make(map[string]*Filter)
	for currLang, currFilter := range configFilters {
		res[currLang] = &Filter{
			IgnoreFileNames:        listToMap(currFilter.IgnoreFileNames),
			IgnoreDirs:             listToMap(currFilter.IgnoreDirs),
			IgnoreFileNameContains: currFilter.IgnoreFileNameContains,
			IgnoreFileNamePrefixes: currFilter.IgnoreFileNamePrefixes,
			IgnoreFileNameSuffixes: currFilter.IgnoreFileNameSuffixes,
			Extension:              currFilter.Extension,
		}
	}

	return res, nil
}

func CreateDefaultConfigFile() error {

	err := os.WriteFile(ConfigFilename, []byte(defaultFiltersContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create configuration file with %v", err)
	}

	return nil
}

func (c *container) BuildFilter(lang string) (*Filter, error) {

	res, ok := c.filters[lang]
	if !ok {
		return nil, fmt.Errorf("%s programming language is currently not supported", lang)
	}

	return res, nil
}

func (f *Filter) Ignore(path string, info os.FileInfo) (bool, error) {

	if info.IsDir() {
		if ok := f.IgnoreDirs[info.Name()]; ok {
			return true, filepath.SkipDir
		}
		return true, nil
	}

	if ok := f.IgnoreFileNames[info.Name()]; ok {
		return true, nil
	}

	if f.IgnoreFileNameContains != "" && strings.Contains(strings.ToLower(info.Name()), f.IgnoreFileNameContains) {
		return true, nil
	}

	for _, currPrefix := range f.IgnoreFileNamePrefixes {
		if strings.HasPrefix(info.Name(), currPrefix) {
			return true, nil
		}
	}

	for _, currSuffix := range f.IgnoreFileNameSuffixes {
		if strings.HasSuffix(info.Name(), currSuffix) {
			return true, nil
		}
	}

	return !strings.HasSuffix(path, f.Extension), nil
}

// Convert a list of strings to a map with value true
func listToMap(list []string) map[string]bool {

	res := make(map[string]bool)
	for _, item := range list {
		res[item] = true
	}

	return res
}

func getConfigFilterContent() ([]byte, error) {

	if fileExists(ConfigFilename) {
		res, err := os.ReadFile(ConfigFilename)
		if err != nil {
			return nil, fmt.Errorf("failed to read configuration file %s", ConfigFilename)
		}
		fmt.Println("Loading configuration file " + ConfigFilename)
		return res, nil
	}

	return []byte(defaultFiltersContent), nil
}

func fileExists(filename string) bool {

	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
