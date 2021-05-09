package main

import (
	"errors"
	"fmt"
	"os"

	_ "embed"

	"github.com/pelletier/go-toml"
)

//go:embed templates/repo.html
var repoTmpl string

//go:embed templates/index.html
var indexTmpl string

type Config struct {
	Repos         map[string]Repository `toml:"repos"`
	BasePath      string                `toml:"basePath"`
	Index         Index                 `toml:"index"`
	RepoTemplate  string                `toml:"repoTemplate"`
	IndexTemplate string                `toml:"indexTemplate"`
	Output        string                `toml:"output" default:"dist"`
}

type Index struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	ExtraHead   string `toml:"extra_head"`
}

type Repository struct {
	URL         string `toml:"repo"`
	Description string `toml:"description"`
	Name        string `toml:"name"`
}

func (r Repository) String() string {
	return r.URL
}

func NewConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	if config.RepoTemplate == "" {
		config.RepoTemplate = repoTmpl
	}

	if config.IndexTemplate == "" {
		config.IndexTemplate = indexTmpl
	}

	err = config.isValid()
	return config, err
}

func (c Config) isValid() error {
	if c.BasePath == "" {
		return errors.New("basePath is empty or missing in config")
	}

	if c.Index == (Index{}) || c.Index.Title == "" {
		return errors.New("index.title is empty or missing in config")
	}

	for path, repo := range c.Repos {
		if repo.Name == "" {
			return fmt.Errorf("repo.%s.name is empty or missing in config", path)
		}
		if repo.URL == "" {
			return fmt.Errorf("repo.%s.repo is empty or missing in config", path)
		}
	}

	return nil
}
