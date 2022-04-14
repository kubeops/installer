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

crd-importer \
    --input=https://github.com/kubeops/csi-driver-cacerts/raw/v0.0.1/crds/cacerts.csi.cert-manager.io_caproviderclasses.yaml \
    --out=./charts/cert-manager-csi-driver-cacerts/crds

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/kubernetes-1.21.1/crds/appcatalog.appscode.com_appbindings.yaml \
    --out=./charts/kube-ui-server/crds

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/kubernetes-1.21.1/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --input=https://github.com/prometheus-operator/prometheus-operator/raw/v0.54.1/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml \
    --out=./charts/panopticon/crds

{
    supervisor_dir=${1:-}/supervisor/crds

    supervisor_repo_url=https://github.com/kubeops/supervisor.git
    supervisor_repo_tag=${KUBEOPS_SUPERVISOR_TAG:-master}

    if [ "$#" -ne 1 ]; then
        if [ "${supervisor_repo_tag}" == "master" ]; then
            echo "Error: missing path_to_input_crds_directory"
            echo "Usage: import-crds.sh <path_to_input_crds_directory>"
            exit 1
        fi

        tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
        # always cleanup temp dir
        # ref: https://opensource.com/article/20/6/bash-trap
        trap \
            "{ rm -rf "${tmp_dir}"; }" \
            SIGINT SIGTERM ERR EXIT

        mkdir -p ${tmp_dir}
        pushd $tmp_dir
        git clone $supervisor_repo_url
        repo_dir=$(ls -b1)
        cd $repo_dir
        git checkout $supervisor_repo_tag
        popd
        supervisor_dir=${tmp_dir}/${repo_dir}/crds
    fi

    crd-importer \
        --input=${supervisor_dir} \
        --out=./charts/supervisor/crds
}
