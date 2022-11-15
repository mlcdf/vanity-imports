package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"
)

type vanityTemplateData struct {
	Root    string
	VCS     string
	RepoURL string
}

// Build the pages
func build(config *Config) error {
	t, err := template.New("").Parse(`{{define "T"}}` + config.RepoTemplate + "{{end}}")
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for root, repo := range config.Repos {
		err := os.MkdirAll(path.Join(wd, config.Output, root), os.ModePerm)
		if err != nil {
			return err
		}

		dest := path.Join(wd, config.Output, root, "index.html")
		log.Printf("%s %s\n", info("Building ["+root+"]"), dest)

		f, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}

		data := vanityTemplateData{
			Root:    path.Join(config.Domain, root),
			VCS:     repo.VCS,
			RepoURL: repo.URL,
		}

		err = t.ExecuteTemplate(f, "T", data)

		if err != nil {
			return err
		}
	}

	return nil
}

// Build the index page
func buildIndex(config *Config) error {
	funcMap := template.FuncMap{
		"unescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"replace": func(input, from, to string) string {
			return strings.Replace(input, from, to, -1)
		}}

	t, err := template.New("").Funcs(funcMap).Parse(`{{define "T"}}` + config.IndexTemplate + "{{end}}")
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	dest := path.Join(wd, config.Output, "index.html")
	log.Printf("%s %s\n", info("Building [index]"), dest)

	f, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(f, "T", config)
	if err != nil {
		return err
	}
	return nil
}

func info(txt string) string {
	return fmt.Sprintf("\033[0;34m%s\033[m", txt)
}
