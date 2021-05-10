package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Version = "(devel)"

const usage = `Usage:
    vanity-imports [option]

Options:				
    --init                  creates a sample .vanity-imports.toml config file
    -c, --config CONFIG     path to the config. Defaults to .vanity-imports.toml
    -V, --version           print version
`

func main() {
	log.SetFlags(0)
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }

	var (
		configFlag  string = defaultConfigName
		versionFlag bool
		initFlag    bool
	)

	flag.StringVar(&configFlag, "c", configFlag, "path to the config")
	flag.StringVar(&configFlag, "config", configFlag, "path to the config")

	flag.BoolVar(&initFlag, "init", initFlag, "creates a sample .vanity-imports.toml config file")

	flag.BoolVar(&versionFlag, "version", versionFlag, "print the version")
	flag.BoolVar(&versionFlag, "V", versionFlag, "print the version")

	flag.Parse()

	if versionFlag {
		log.Println(Version)
		return
	}

	if initFlag {
		err := initSampleConfig()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	config, err := newConfig(configFlag)
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
