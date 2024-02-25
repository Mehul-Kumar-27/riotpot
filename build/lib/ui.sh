#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function riotpot::ui::build() {
    cd ui
    npm install
    npm run build
}

riotpot::ui::install_node() {
    # Check if nvm is installed
    if ! command -v nvm &> /dev/null; then
        # nvm not found, install it
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash

        # Ensure nvm is available in the current session
        export NVM_DIR="/usr/local/share/nvm"
        [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
        [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
    fi

    # Load nvm in the current session
    export NVM_DIR="/usr/local/share/nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

    # Install Node.js version 18.18.0 and set it as default
    nvm install 18.18.0
    nvm use 18.18.0
}


function riotpot::ui::serve_dev() {
    serve -s ./ui/build
}
