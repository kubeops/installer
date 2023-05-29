#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eou pipefail

SCRIPT_ROOT=$(realpath $(dirname "${BASH_SOURCE[0]}")/../..)
SCRIPT_NAME=$(basename "${BASH_SOURCE[0]}")
pushd $SCRIPT_ROOT/charts/gatekeeper-library

# http://redsymbol.net/articles/bash-exit-traps/
function cleanup() {
    popd
}
trap cleanup EXIT

git clone https://github.com/open-policy-agent/gatekeeper-library.git
rm -rf library
mv gatekeeper-library/library .
mv gatekeeper-library/LICENSE library/LICENSE
rm -rf gatekeeper-library
find library -type f \( -path '**/template.yaml' -o -path '**/constraint.yaml' \) -prune -o -name '*.yaml' -exec rm -rf {} \;
