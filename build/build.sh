#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RIOTPOT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${RIOTPOT_ROOT}/build/common.sh"
source "${RIOTPOT_ROOT}/build/requirements.sh"
source "${RIOTPOT_ROOT}/build/lib/golang.sh"
source "${RIOTPOT_ROOT}/build/lib/ui.sh"

riotpot::golang::setup_env
riotpot::requirements::install
riotpot::golang::install_requirements

# [1-3-2024] TODO: building in docker fails to complete.
riotpot::ui::install_node
riotpot::ui::build