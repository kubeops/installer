# Config Syncer Manager

[Config Syncer OCM AddOn Manager](https://github.com/kubeops/config-syncer) - Config Syncer OCM AddOn Manager

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/config-syncer-manager --version=v1.0.0
$ helm upgrade -i config-syncer-manager appscode/config-syncer-manager -n open-cluster-management-addon --create-namespace --version=v1.0.0
```

## Introduction

This chart deploys a Config Syncer OCM AddOn Manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.29+
- Open Cluster Management (https://open-cluster-management.io) hub

## Installing the Chart

To install/upgrade the chart with the release name `config-syncer-manager`:

```bash
$ helm upgrade -i config-syncer-manager appscode/config-syncer-manager -n open-cluster-management-addon --create-namespace --version=v1.0.0
```

The command deploys a Config Syncer OCM AddOn Manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `config-syncer-manager`:

```bash
$ helm uninstall config-syncer-manager -n open-cluster-management-addon
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `config-syncer-manager` chart and their default values.

|                Parameter                 |                                                                             Description                                                                             |                   Default                   |
|------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| nameOverride                             | Overrides name template                                                                                                                                             | <code>""</code>                             |
| fullnameOverride                         | Overrides fullname template                                                                                                                                         | <code>""</code>                             |
| registryFQDN                             | Docker registry fqdn used to pull config-syncer images on every spoke.                                                                                              | <code>ghcr.io</code>                        |
| image                                    | Manager image. Set tag to override the chart appVersion.                                                                                                            | <code>ghcr.io/appscode/config-syncer</code> |
| tag                                      |                                                                                                                                                                     | <code>""</code>                             |
| imagePullPolicy                          |                                                                                                                                                                     | <code>Always</code>                         |
| clusterName                              | Name of the hub cluster, recorded as origin cluster on objects synced to spokes.                                                                                    | <code>hub</code>                            |
| placement.create                         |                                                                                                                                                                     | <code>true</code>                           |
| placement.name                           |                                                                                                                                                                     | <code>global</code>                         |
| kubeconfigSecretName                     | Name of a Secret holding an OCM hub kubeconfig; when set, the manager uses --kubeconfig instead of the in-cluster ServiceAccount.                                   | <code>""</code>                             |
| securityContext.allowPrivilegeEscalation |                                                                                                                                                                     | <code>false</code>                          |
| securityContext.privileged               |                                                                                                                                                                     | <code>false</code>                          |
| securityContext.runAsNonRoot             |                                                                                                                                                                     | <code>true</code>                           |
| securityContext.runAsUser                |                                                                                                                                                                     | <code>65534</code>                          |
| securityContext.readOnlyRootFilesystem   |                                                                                                                                                                     | <code>true</code>                           |
| securityContext.seccompProfile.type      |                                                                                                                                                                     | <code>RuntimeDefault</code>                 |
| envFrom                                  | List of sources to populate environment variables in the container                                                                                                  | <code>[]</code>                             |
| env                                      | List of environment variables to set in the container                                                                                                               | <code>[]</code>                             |
| monitoring.agent                         | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                | <code>""</code>                             |
| monitoring.serviceMonitor.labels         | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`. | <code>{}</code>                             |
| distro.openshift                         | Set true, if installed in OpenShift                                                                                                                                 | <code>false</code>                          |
| distro.ubi                               | Set operator or all to use ubi images                                                                                                                               | <code>""</code>                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i config-syncer-manager appscode/config-syncer-manager -n open-cluster-management-addon --create-namespace --version=v1.0.0 --set registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i config-syncer-manager appscode/config-syncer-manager -n open-cluster-management-addon --create-namespace --version=v1.0.0 --values values.yaml
```
