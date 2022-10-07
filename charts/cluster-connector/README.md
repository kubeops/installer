# Kubernetes Cluster Connector

[Kubernetes Cluster Connector by AppsCode](https://github.com/kubeops/cluster-connector) - Kubernetes Cluster Connector for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cluster-connector --version=v2022.06.14
$ helm upgrade -i cluster-connector appscode/cluster-connector -n kubeops --create-namespace --version=v2022.06.14
```

## Introduction

This chart deploys a Kubernetes Cluster Connector on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `cluster-connector`:

```bash
$ helm upgrade -i cluster-connector appscode/cluster-connector -n kubeops --create-namespace --version=v2022.06.14
```

The command deploys a Kubernetes Cluster Connector on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cluster-connector`:

```bash
$ helm uninstall cluster-connector -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-connector` chart and their default values.

|         Parameter          |                                                                                                            Description                                                                                                             |            Default             |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------|
| nameOverride               | Overrides name template                                                                                                                                                                                                            | <code>""</code>                |
| fullnameOverride           | Overrides fullname template                                                                                                                                                                                                        | <code>""</code>                |
| replicaCount               |                                                                                                                                                                                                                                    | <code>1</code>                 |
| registryFQDN               | Docker registry fqdn used to pull docker images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                     | <code>""</code>                |
| image.registry             | Docker registry used to pull operator image                                                                                                                                                                                        | <code>appscode</code>          |
| image.repository           | Name of operator container image                                                                                                                                                                                                   | <code>cluster-connector</code> |
| image.tag                  | Overrides the image tag whose default is the chart appVersion.                                                                                                                                                                     | <code>""</code>                |
| image.resources            | Compute Resources required by the operator container                                                                                                                                                                               | <code>{}</code>                |
| image.securityContext      | Security options the operator container should run with                                                                                                                                                                            | <code>{}</code>                |
| imagePullSecrets           | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                |
| imagePullPolicy            | Container image pull policy                                                                                                                                                                                                        | <code>Always</code>            |
| serviceAccount.create      | Specifies whether a service account should be created                                                                                                                                                                              | <code>true</code>              |
| serviceAccount.annotations | Annotations to add to the service account                                                                                                                                                                                          | <code>{}</code>                |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                             | <code>""</code>                |
| podAnnotations             |                                                                                                                                                                                                                                    | <code>{}</code>                |
| podSecurityContext         |                                                                                                                                                                                                                                    | <code>{}</code>                |
| nodeSelector               |                                                                                                                                                                                                                                    | <code>{}</code>                |
| tolerations                |                                                                                                                                                                                                                                    | <code>[]</code>                |
| affinity                   |                                                                                                                                                                                                                                    | <code>{}</code>                |
| linkID                     |                                                                                                                                                                                                                                    | <code>""</code>                |
| nats.addr                  |                                                                                                                                                                                                                                    | <code>""</code>                |
| nats.encodedCreds          |                                                                                                                                                                                                                                    | <code>""</code>                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cluster-connector appscode/cluster-connector -n kubeops --create-namespace --version=v2022.06.14 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cluster-connector appscode/cluster-connector -n kubeops --create-namespace --version=v2022.06.14 --values values.yaml
```
