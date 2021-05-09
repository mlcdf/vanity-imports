
# vanity-imports

Generate HTML pages that allows you to set ["custom" or "vanity" import paths](https://golang.org/doc/go1.4#canonicalimports) for your Go packages using the `go-import` meta tag ([read the specs](https://golang.org/cmd/go/#hdr-Remote_import_paths)).

For example, this package is import path is `go.mlcdf.fr/vanity-imports` (instead of `github.com/mlcdf/vanity-imports`).

## Highlights

- Painless to host: it's only static files
- Designed to be used in CI environment: ship as single binary with no OS dependencies
- Easy to configure and extend via a single TOML configuration file

## Install

- From [GitHub releases](https://github.com/mlcdf/dyndns/releases): download the binary corresponding to your OS and architecture.
- From source (make sure `$GOPATH/bin` is in your `$PATH`):

```sh
go get go.mlcdf.fr/vanity-imports
```

## Usage

```
Usage:
    vanity-imports [OPTION]

Options:
    -c, --config CONFIG     path to the config. Defaults to .vanity-imports.toml
    -V, --version           print version
```

First, create a [config file](#configuration-format).
```sh
touch .vanity-imports.toml
```

Generate the HTML pages
```sh
vanity-imports
```

Upload the content of the `dist` directory to your web server.

## Configuration format

Format for the `.vanity-imports.toml` file.

```toml
basePath = "go.mlcdf.fr" # required
output = "output" # default to dist
repoTemplate = """<golang template>""" # override the default template for the repo page
indexTemplate = """<golang template>""" # override the default template for the index page

[index]
description = ""
extra_head = "" # extra html tags appended to the head
title = "" # required

[repos]

[repos."/dyndns"] # basePath + "/dyndns will be your package import path
name = "dyndns" # required
repo = "https://github.com/mlcdf/dyndns/" # required. Url to the source repository
```

## Examples

Checkout the [netlify](https://github.com/mlcdf/vanity-imports/tree/netlify) branch to find the code behind https://go.mlcdf.fr/.

## Run Locally

Clone the project

```bash
  git clone https://github.com/mlcdf/vanity-imports.git
```

Go to the project directory

```bash
  cd vanity-imports
```

Start the app

```bash
  go run .
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
