package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
)

// Build the pages
func build(config *Config) error {

	t, err := template.New("foo").Parse(`{{define "T"}}` + config.RepoTemplate + "{{end}}")
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _path, repo := range config.Repos {
		os.MkdirAll(path.Join(wd, config.Output, _path), os.ModePerm)

		dest := path.Join(wd, config.Output, _path, "index.html")
		log.Printf("%s %s\n", info("Building ["+_path+"]"), dest)

		f, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}

		err = t.ExecuteTemplate(f, "T", struct {
			Path   string
			Repo   Repository
			Config Config
		}{_path, repo, *config})

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
		}}

	t, err := template.New("foo").Funcs(funcMap).Parse(`{{define "T"}}` + config.IndexTemplate + "{{end}}")
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
