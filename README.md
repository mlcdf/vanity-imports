
# vanity-imports

[![test](https://github.com/mlcdf/vanity-imports/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/mlcdf/vanity-imports/actions/workflows/test.yml)


Generate HTML pages that allows you to set ["custom" or "vanity" import paths](https://golang.org/doc/go1.4#canonicalimports) for your Go packages using the `go-import` meta tag ([read the specs](https://golang.org/cmd/go/#hdr-Remote_import_paths)).

For example, this package import path is `go.mlcdf.fr/vanity-imports` (instead of `github.com/mlcdf/vanity-imports`).

## Highlights

- Painless to host: it's only static files
- Designed to be used in a CI environment: ship as single binary with no OS dependencies
- Easy to configure and extend via a single TOML configuration file
- Use your own template for the index and the repo pages.

## Install

- From [GitHub releases](https://github.com/mlcdf/dyndns/releases): download the binary corresponding to your OS and architecture.
- From source (make sure `$GOPATH/bin` is in your `$PATH`):

```sh
go get go.mlcdf.fr/vanity-imports
```

## Usage

```
Usage:
    vanity-imports [option]

Options:
    --init                  creates a sample .vanity-imports.toml config file
    -c, --config CONFIG     path to the config. Defaults to .vanity-imports.toml
    -V, --version           print version
```

First, create a [config file](#configuration-format).
```sh
vanity-imports --init
```

Generate the HTML pages
```sh
vanity-imports
```

Upload the content of the `dist` directory to your web server or your favorite static hosting service such as GitHub Pages, Netlify or Vercel.

## Configuration format

Format for the `.vanity-imports.toml` file.

```toml
domain = "go.mlcdf.fr" # required
output = "output" # default to dist
repo_template = """<golang template>""" # override the default template for the repo page
index_template = """<golang template>""" # override the default template for the index page

[index]
description = ""
extra_head = "" # extra html tags appended to the head
title = "" # required

[repos]

[repos."/dyndns"] # domain + "/dyndns will be your package import path
repo = "https://github.com/mlcdf/dyndns/" # required. Url to the source repository
```

## Examples

Check out the [netlify](https://github.com/mlcdf/vanity-imports/tree/netlify) branch to find the code behind https://go.mlcdf.fr/.

## Run Locally

Clone the project

```bash
git clone https://github.com/mlcdf/vanity-imports.git
```

Go to the project directory

```bash
cd vanity-imports
```

Generate the pages

```bash
go run .
```

View the output on your browser

```bash
cd dist
python3 -m http.server 8000
# open your browser to localhost:8000
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
