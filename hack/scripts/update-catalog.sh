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

CI_VALUES=charts/config-syncer/ci/ci-values.yaml

# Set a fake license so config-syncer templates render for image discovery,
# then reset it back to an empty license afterwards.
reset_license() {
    sed -i.bak 's/^license: .*/license: ""/' "$CI_VALUES" && rm -f "${CI_VALUES}.bak"
}
trap reset_license EXIT

sed -i.bak 's/^license: .*/license: fake-license/' "$CI_VALUES" && rm -f "${CI_VALUES}.bak"

image-packer list --root-dir=charts --output-dir=catalog

reset_license
trap - EXIT

# image-packer generate-scripts --insecure --allow-nondistributable-artifacts \
#     --output-dir=catalog \
#     --src=catalog/imagelist.yaml

make add-license fmt
