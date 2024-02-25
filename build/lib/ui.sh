#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function riotpot::ui::build() {
    cd ui
    npm install
    npm run build
}

function riotpot::ui::install_node() {
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash
    nvm install 18.18.0
    nvm use 18.18.0

}

function riotpot::ui::serve_dev() {
    serve -s ./ui/build
}
