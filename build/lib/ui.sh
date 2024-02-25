#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function riotpot::ui::build() {
    cd ui
    npm install
    npm run build
}

riotpot::ui::install_nvm() {
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm

}

riotpot::ui::install_node() {
    nvm install 18.18.0
    nvm use 18.18.0
}

function riotpot::ui::serve_dev() {
    serve -s ./ui/build
}
