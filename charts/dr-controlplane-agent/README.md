# DR Control Plane Agent

[DR Control Plane Agent by AppsCode](https://github.com/kubeops/dr-controlplane) - Per data center agent of the DC failover service

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/dr-controlplane-agent --version=v0.1.0
$ helm upgrade -i dr-controlplane-agent appscode/dr-controlplane-agent -n dc-failover --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys a DR Control Plane Agent on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.26+

## Installing the Chart

To install/upgrade the chart with the release name `dr-controlplane-agent`:

```bash
$ helm upgrade -i dr-controlplane-agent appscode/dr-controlplane-agent -n dc-failover --create-namespace --version=v0.1.0
```

The command deploys a DR Control Plane Agent on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `dr-controlplane-agent`:

```bash
$ helm uninstall dr-controlplane-agent -n dc-failover
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `dr-controlplane-agent` chart and their default values.

|              Parameter              |                                                                                                                                                                                                                                 Description                                                                                                                                                                                                                                  |                     Default                     |
|-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| clusterName                         | clusterName is the ManagedCluster name. Injected by the addon manager and used as the agent --dc-name. Required at render time.                                                                                                                                                                                                                                                                                                                                              | <code>""</code>                                 |
| namespace                           | namespace where the agent runs and the coordination Leases live on the spoke. It must be the namespace that already holds the coordination credential Secret (agent.coordKubeconfigSecret), otherwise the agent pod never leaves ContainerCreating. The addon manager pins this to its --agent-install-namespace and also sets the addon-framework install namespace to the same value.                                                                                      | <code>dc-failover</code>                        |
| addonInstallNamespace               | addonInstallNamespace is a built in value the addon-framework always sets to the effective addon install namespace. It wins over namespace when present, so the rendered objects can never drift from where the framework installs the release.                                                                                                                                                                                                                              | <code>""</code>                                 |
| createNamespace                     | createNamespace makes the chart emit the target Namespace object, which makes the addon's ManifestWork OWN that namespace. Default false on purpose: the agent shares the pre existing dc-failover namespace with the coordination credential Secret, the active DC markers and, on some clusters, etcd. If the addon owned that namespace, deleting the addon would garbage collect the entire DR install with it. Only turn this on for a namespace nothing else lives in. | <code>false</code>                              |
| registryFQDN                        | Docker registry fqdn used to pull the dr-controlplane image. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${repository}. Injected by the addon manager from --agent-registry-fqdn.                                                                                                                                                                                                                                                                  | <code>ghcr.io</code>                            |
| image.registry                      | Docker registry used to pull the image                                                                                                                                                                                                                                                                                                                                                                                                                                       | <code>appscode</code>                           |
| image.repository                    | Name of the container image                                                                                                                                                                                                                                                                                                                                                                                                                                                  | <code>dr-controlplane</code>                    |
| image.tag                           | Container image tag                                                                                                                                                                                                                                                                                                                                                                                                                                                          | <code>"" # defaults to .Chart.AppVersion</code> |
| imagePullPolicy                     | Container image pull policy                                                                                                                                                                                                                                                                                                                                                                                                                                                  | <code>IfNotPresent</code>                       |
| imagePullSecrets                    | imagePullSecrets are Secret references that must exist in .Values.namespace on each spoke. Injected by the addon manager from --agent-image-pull-secrets.                                                                                                                                                                                                                                                                                                                    | <code>[]</code>                                 |
| replicas                            | replicas is the agent pod count on this spoke. Injected by the addon manager from --agent-replicas. More than one is pod crash insurance only: every replica shares the DC identity (same --dc-name), renews the same Leases, and projects the same markers, so extras are harmless but carry no quorum weight; quorum lives in the etcd members alone.                                                                                                                      | <code>1</code>                                  |
| agent.coordKubeconfigData           | base64 kubeconfig injected by the addon manager; empty means an operator provisioned Secret of the same name is used instead.                                                                                                                                                                                                                                                                                                                                                | <code>""</code>                                 |
| agent.coordKubeconfigSecret         | Kubeconfig Secret for the coordination control plane (key: kubeconfig). Empty disables the mount and the agent uses the in cluster config.                                                                                                                                                                                                                                                                                                                                   | <code>coord-kubeconfig</code>                   |
| agent.metricsAddr                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>":8080"</code>                            |
| agent.health.leaseDurationSeconds   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>15</code>                                 |
| agent.health.renewIntervalSeconds   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>5</code>                                  |
| agent.election.leaseDurationSeconds |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>45</code>                                 |
| agent.election.renewDeadlineSeconds |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>40</code>                                 |
| agent.election.retryPeriodSeconds   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>5</code>                                  |
| agent.resources                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>{}</code>                                 |
| nodeSelector                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>{}</code>                                 |
| tolerations                         |                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | <code>[]</code>                                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i dr-controlplane-agent appscode/dr-controlplane-agent -n dc-failover --create-namespace --version=v0.1.0 --set namespace=dc-failover
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i dr-controlplane-agent appscode/dr-controlplane-agent -n dc-failover --create-namespace --version=v0.1.0 --values values.yaml
```
