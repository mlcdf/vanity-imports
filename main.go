package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Version = "(devel)"

const usage = `Usage:
    vanity-urls [OPTION]

Options:
    --config CONFIG         path to the config. Defaults to .vanity-imports.toml
    -V, --version           print version
`

func main() {
	log.SetFlags(0)
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }

	var (
		configFlag  string = ".vanity-imports.toml"
		versionFlag bool
	)

	flag.StringVar(&configFlag, "config", configFlag, "path to the config")

	flag.BoolVar(&versionFlag, "version", versionFlag, "print the version")
	flag.BoolVar(&versionFlag, "V", versionFlag, "print the version")

	flag.Parse()

	if versionFlag {
		log.Println(Version)
		return
	}

	config, err := NewConfig(configFlag)
	if err != nil {
		log.Fatal(err)
	}

	err = build(config)
	if err != nil {
		log.Fatal(err)
	}

	err = buildIndex(config)
	if err != nil {
		log.Fatal(err)
	}
}
