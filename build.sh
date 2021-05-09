#! /usr/bin/env bash
set -e

# download latest release of dyndns
LOCATION=$(curl -s https://api.github.com/repos/mlcdf/vanity-imports/releases/latest | grep browser_download_url | grep linux-amd64 | cut -d '"' -f 4)
echo $LOCATION
curl -L "${LOCATION}" -o vanity-imports

# Print the version for debugging
vanity-imports --version

# Build the pages using .vanity-imports.toml
vanity-imports
