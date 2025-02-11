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

CERT_MANAGER_CERT_MANAGER_TAG=${CERT_MANAGER_CERT_MANAGER_TAG:-v1.15.4}
FLUXCD_HELM_CONTROLLER_TAG=${FLUXCD_HELM_CONTROLLER_TAG:-v1.0.1}
FLUXCD_SOURCE_CONTROLLER_TAG=${FLUXCD_SOURCE_CONTROLLER_TAG:-v1.3.0}
KMODULES_CUSTOM_RESOURCES_TAG=${KMODULES_CUSTOM_RESOURCES_TAG:-v0.29.1}
KMODULES_RESOURCE_METADATA_TAG=${KMODULES_RESOURCE_METADATA_TAG:-master}
KUBEOPS_CSI_DRIVER_CACERTS_TAG=${KUBEOPS_CSI_DRIVER_CACERTS_TAG:-v0.1.0}
KUBEOPS_EXTERNAL_DNS_OPERATOR_TAG=${KUBEOPS_EXTERNAL_DNS_OPERATOR_TAG:-v0.0.9}
KUBEOPS_PETSET_TAG=${KUBEOPS_PETSET_TAG:-master}
KUBEOPS_SIDEKICK_TAG=${KUBEOPS_SIDEKICK_TAG:-v0.0.9}
OPEN_POLICY_AGENT_GATEKEEPER_TAG=${OPEN_POLICY_AGENT_GATEKEEPER_TAG:-v3.14.0}
OPEN_VIZ_APIMACHINERY_TAG=${OPEN_VIZ_APIMACHINERY_TAG:-v0.0.8}
PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG=${PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG:-kube-prometheus-stack-69.2.2}
PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG=${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG:-v0.75.2}
X_HELM_APIMACHINERY_TAG=${X_HELM_APIMACHINERY_TAG:-v0.0.16}

crd-importer \
    --input=https://github.com/kubeops/csi-driver-cacerts/raw/${KUBEOPS_CSI_DRIVER_CACERTS_TAG}/crds/cacerts.csi.cert-manager.io_caproviderclasses.yaml \
    --out=./charts/cert-manager-csi-driver-cacerts/crds

crd-importer \
    --input=https://github.com/cert-manager/cert-manager/releases/download/${CERT_MANAGER_CERT_MANAGER_TAG}/cert-manager.crds.yaml \
    --gk=ClusterIssuer.cert-manager.io --gk=Issuer.cert-manager.io \
    --out=./charts/cert-manager-csi-driver-cacerts/crds

crd-importer \
    --input=https://github.com/kubeops/external-dns-operator/raw/${KUBEOPS_EXTERNAL_DNS_OPERATOR_TAG}/crds/external-dns.appscode.com_externaldnses.yaml \
    --out=./charts/external-dns-operator/crds

crd-importer \
    --input=https://github.com/open-policy-agent/gatekeeper/raw/${OPEN_POLICY_AGENT_GATEKEEPER_TAG}/charts/gatekeeper/crds/constrainttemplate-customresourcedefinition.yaml \
    --out=./charts/gatekeeper-library/crds

crd-importer \
    --input=https://github.com/kubeops/sidekick/raw/${KUBEOPS_SIDEKICK_TAG}/crds/apps.k8s.appscode.com_sidekicks.yaml \
    --out=./charts/sidekick/crds

crd-importer \
    --input=https://github.com/fluxcd/helm-controller/raw/${FLUXCD_HELM_CONTROLLER_TAG}/config/crd/bases/helm.toolkit.fluxcd.io_helmreleases.yaml \
    --input=https://github.com/fluxcd/source-controller/raw/${FLUXCD_SOURCE_CONTROLLER_TAG}/config/crd/bases/source.toolkit.fluxcd.io_helmrepositories.yaml \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/appcatalog.appscode.com_appbindings.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/management.k8s.appscode.com_projectquotas.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_clusterprofiles.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_features.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_featuresets.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_resourcedashboards.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_resourceeditors.yaml \
    --input=https://github.com/kmodules/resource-metadata/raw/${KMODULES_RESOURCE_METADATA_TAG}/crds/ui.k8s.appscode.com_resourceoutlinefilters.yaml \
    --input=https://github.com/prometheus-community/helm-charts/raw/${PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG}/charts/kube-prometheus-stack/charts/crds/crds/crd-servicemonitors.yaml \
    --input=https://github.com/x-helm/apimachinery/raw/${X_HELM_APIMACHINERY_TAG}/crds/charts.x-helm.dev_chartpresets.yaml \
    --input=https://github.com/x-helm/apimachinery/raw/${X_HELM_APIMACHINERY_TAG}/crds/charts.x-helm.dev_clusterchartpresets.yaml \
    --out=./charts/kube-ui-server/crds

crd-importer \
    --input=https://github.com/open-viz/apimachinery/raw/${OPEN_VIZ_APIMACHINERY_TAG}/crds/openviz.dev_grafanadashboards.yaml \
    --out=./charts/gatekeeper-grafana-dashboards/crds

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --input=https://github.com/prometheus-community/helm-charts/raw/${PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG}/charts/kube-prometheus-stack/charts/crds/crds/crd-servicemonitors.yaml \
    --out=./charts/panopticon/crds

crd-importer \
    --input=https://github.com/open-viz/apimachinery/raw/${OPEN_VIZ_APIMACHINERY_TAG}/crds/openviz.dev_grafanadashboards.yaml \
    --input=https://github.com/prometheus-community/helm-charts/raw/${PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG}/charts/kube-prometheus-stack/charts/crds/crds/crd-servicemonitors.yaml \
    --out=./charts/falco-ui-server/crds

crd-importer \
    --no-description \
    --input=https://github.com/kubeops/operator-shard-manager/raw/master/crds/operator.k8s.appscode.com_shardconfigurations.yaml \
    --out=./charts/operator-shard-manager/crds

crd-importer \
    --no-description \
    --input=https://github.com/kubeops/petset/raw/${KUBEOPS_PETSET_TAG}/crds/apps.k8s.appscode.com_petsets.yaml \
    --input=https://github.com/kubeops/petset/raw/${KUBEOPS_PETSET_TAG}/crds/apps.k8s.appscode.com_placementpolicies.yaml \
    --out=./charts/petset/crds

crd-importer \
    --input=https://github.com/open-viz/apimachinery/raw/${OPEN_VIZ_APIMACHINERY_TAG}/crds/openviz.dev_grafanadashboards.yaml \
    --input=https://github.com/prometheus-community/helm-charts/raw/${PROMETHEUS_COMMUNITY_HELM_CHARTS_TAG}/charts/kube-prometheus-stack/charts/crds/crds/crd-servicemonitors.yaml \
    --out=./charts/scanner/crds

# import cert-manager crds
crd-importer \
    --input=https://github.com/cert-manager/cert-manager/releases/download/${CERT_MANAGER_CERT_MANAGER_TAG}/cert-manager.crds.yaml \
    --out=./charts/cert-manager-crds/crds

# import prometheus-operator crds
crd-importer \
    --input=https://github.com/prometheus-operator/prometheus-operator/releases/download/${PROMETHEUS_OPERATOR_PROMETHEUS_OPERATOR_TAG}/stripped-down-crds.yaml \
    --out=./charts/prometheus-operator-crds/crds

crd-importer \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/appcatalog.appscode.com_appbindings.yaml \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --out=./charts/kmodules-crds/crds
rm -rf charts/kmodules-crds/crds/auditor.appscode.com_siteinfoes.yaml
rm -rf charts/kmodules-crds/crds/auditor.appscode.com_siteinfos.yaml

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
