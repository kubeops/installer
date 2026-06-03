# FluxCD to ArgoCD Bridge Manager

[FluxCD to ArgoCD Bridge OCM AddOn Manager](https://github.com/kubeops/fargocd) - FluxCD to ArgoCD Bridge OCM AddOn Manager

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/fargocd-manager --version=v2026.5.22
$ helm upgrade -i fargocd-manager appscode/fargocd-manager -n open-cluster-management-fargocd --create-namespace --version=v2026.5.22
```

## Introduction

This chart deploys a FluxCD to ArgoCD Bridge OCM AddOn Manager on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.29+
- Open Cluster Management (https://open-cluster-management.io) hub

## Installing the Chart

To install/upgrade the chart with the release name `fargocd-manager`:

```bash
$ helm upgrade -i fargocd-manager appscode/fargocd-manager -n open-cluster-management-fargocd --create-namespace --version=v2026.5.22
```

The command deploys a FluxCD to ArgoCD Bridge OCM AddOn Manager on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `fargocd-manager`:

```bash
$ helm uninstall fargocd-manager -n open-cluster-management-fargocd
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `fargocd-manager` chart and their default values.

|                Parameter                 |                                                                                                                                                                                                                   Description                                                                                                                                                                                                                   |                    Default                    |
|------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| nameOverride                             | Overrides name template                                                                                                                                                                                                                                                                                                                                                                                                                         | <code>""</code>                               |
| fullnameOverride                         | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>                               |
| registryFQDN                             | Docker registry fqdn used to pull fargocd docker images on every spoke (propagated as the chart-level registryFQDN value of the spoke chart).                                                                                                                                                                                                                                                                                                   | <code>ghcr.io</code>                          |
| image                                    | Manager image. Set tag to override the chart appVersion.                                                                                                                                                                                                                                                                                                                                                                                        | <code>ghcr.io/appscode/fargocd</code>         |
| tag                                      |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>""</code>                               |
| imagePullPolicy                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>Always</code>                           |
| placement.create                         |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>true</code>                             |
| placement.name                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>global</code>                           |
| argocd.mode                              | How fargocd on spoke clusters talks to Argo CD. One of: in-cluster, autonomous, managed                                                                                                                                                                                                                                                                                                                                                         | <code>in-cluster</code>                       |
| argocd.namespace                         | Namespace on the Argo CD cluster where Applications are written. Auto-discovered when empty.                                                                                                                                                                                                                                                                                                                                                    | <code>""</code>                               |
| argocd.destServer                        | Application.spec.destination.server. Defaults to the in-cluster API server.                                                                                                                                                                                                                                                                                                                                                                     | <code>"https://kubernetes.default.svc"</code> |
| argocd.destName                          | Application.spec.destination.name. Set when Argo CD references the workload cluster by symbolic name (typical in managed mode).                                                                                                                                                                                                                                                                                                                 | <code>""</code>                               |
| argocd.project                           | Argo CD Project assigned to generated Applications.                                                                                                                                                                                                                                                                                                                                                                                             | <code>default</code>                          |
| argocd.kubeconfig                        | Raw kubeconfig content for the Argo CD principal cluster (managed mode). When set, the chart creates a Secret in the manager namespace holding this kubeconfig and mounts it into the manager pod; the manager then propagates the kubeconfig to every spoke. Mutually exclusive with kubeconfigSecret. <br> Example: <br> `helm upgrade -i fargocd-manager appscode/fargocd-manager \` <br> `--set-file argocd.kubeconfig=/path/to/kubeconfig` | <code>""</code>                               |
| argocd.kubeconfigSecret                  | Name of an existing Secret in the manager namespace whose data key `kubeconfig` holds the Argo CD principal kubeconfig (managed mode). Mutually exclusive with argocd.kubeconfig.                                                                                                                                                                                                                                                               | <code>""</code>                               |
| argocd.kubeconfigSpokeSecret             | Name of a pre-created Secret on each spoke holding the principal kubeconfig. When set, the manager propagates this name to spokes instead of pushing the inline kubeconfig.                                                                                                                                                                                                                                                                     | <code>""</code>                               |
| kubeconfigSecretName                     | Name of a Secret holding a kubeconfig for the OCM hub. When set, the manager pod mounts the Secret and uses --kubeconfig instead of the in-cluster ServiceAccount.                                                                                                                                                                                                                                                                              | <code>""</code>                               |
| securityContext.allowPrivilegeEscalation |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>false</code>                            |
| securityContext.privileged               |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>false</code>                            |
| securityContext.runAsNonRoot             |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>true</code>                             |
| securityContext.runAsUser                |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>65534</code>                            |
| securityContext.readOnlyRootFilesystem   |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>true</code>                             |
| securityContext.seccompProfile.type      |                                                                                                                                                                                                                                                                                                                                                                                                                                                 | <code>RuntimeDefault</code>                   |
| envFrom                                  | List of sources to populate environment variables in the container                                                                                                                                                                                                                                                                                                                                                                              | <code>[]</code>                               |
| env                                      | List of environment variables to set in the container                                                                                                                                                                                                                                                                                                                                                                                           | <code>[]</code>                               |
| monitoring.agent                         | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                                                                                            | <code>""</code>                               |
| monitoring.serviceMonitor.labels         | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                                                                                             | <code>{}</code>                               |
| distro.openshift                         | Set true, if installed in OpenShift                                                                                                                                                                                                                                                                                                                                                                                                             | <code>false</code>                            |
| distro.ubi                               | Set operator or all to use ubi images                                                                                                                                                                                                                                                                                                                                                                                                           | <code>""</code>                               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i fargocd-manager appscode/fargocd-manager -n open-cluster-management-fargocd --create-namespace --version=v2026.5.22 --set registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i fargocd-manager appscode/fargocd-manager -n open-cluster-management-fargocd --create-namespace --version=v2026.5.22 --values values.yaml
```
