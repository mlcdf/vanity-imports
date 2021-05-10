package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "embed"

	"github.com/pelletier/go-toml"
)

const defaultConfigName = ".vanity-imports.toml"

const sampleConfig = `
domain = "go.example.com"

[index]
description = ""
extra_head = ""
title = "Jane's go packages"

[repos]

[repos."/foobar"]
repo = "https://github.com/jane/foobar"
`

//go:embed templates/repo.html
var repoTmpl string

//go:embed templates/index.html
var indexTmpl string

type Config struct {
	Repos         map[string]Repository `toml:"repos"`
	Domain        string                `toml:"domain"`
	Index         Index                 `toml:"index"`
	RepoTemplate  string                `toml:"repo_template"`
	IndexTemplate string                `toml:"index_template"`
	Output        string                `toml:"output" default:"dist"`
}

type Index struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	ExtraHead   string `toml:"extra_head"`
}

type Repository struct {
	URL string `toml:"repo"`
}

func (r Repository) String() string {
	return r.URL
}

func newConfig(path string) (*Config, error) {
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
	if c.Domain == "" {
		return errors.New("domain is empty or missing in config")
	}

	if c.Index == (Index{}) && c.Index.Title == "" {
		return errors.New("index.title is empty or missing in config")
	}

	for path, repo := range c.Repos {
		if repo.URL == "" {
			return fmt.Errorf("repo.%s.repo is empty or missing in config", path)
		}
	}

	return nil
}

func initSampleConfig() error {
	if _, err := os.Stat(defaultConfigName); os.IsNotExist(err) {
		return os.WriteFile(defaultConfigName, []byte(sampleConfig), 0777)
	}

	log.Printf("%s already exists\n", defaultConfigName)
	return nil
}
