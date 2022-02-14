# Kubernetes Cluster Connector

[Kubernetes Cluster Connector by AppsCode](https://github.com/kubeops/cluster-connector) - Kubernetes Cluster Connector for Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install cluster-connector appscode/cluster-connector -n kubeops
```

## Introduction

This chart deploys a Kubernetes Cluster Connector on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `cluster-connector`:

```console
$ helm install cluster-connector appscode/cluster-connector -n kubeops
```

The command deploys a Kubernetes Cluster Connector on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `cluster-connector`:

```console
$ helm delete cluster-connector -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cluster-connector` chart and their default values.

|         Parameter          |                                                      Description                                                       |                 Default                 |
|----------------------------|------------------------------------------------------------------------------------------------------------------------|-----------------------------------------|
| replicaCount               |                                                                                                                        | <code>1</code>                          |
| image.repository           |                                                                                                                        | <code>appscode/cluster-connector</code> |
| image.pullPolicy           |                                                                                                                        | <code>IfNotPresent</code>               |
| image.tag                  | Overrides the image tag whose default is the chart appVersion.                                                         | <code>"v0.0.1"</code>                   |
| imagePullSecrets           |                                                                                                                        | <code>[]</code>                         |
| nameOverride               |                                                                                                                        | <code>""</code>                         |
| fullnameOverride           |                                                                                                                        | <code>""</code>                         |
| serviceAccount.create      | Specifies whether a service account should be created                                                                  | <code>true</code>                       |
| serviceAccount.annotations | Annotations to add to the service account                                                                              | <code>{}</code>                         |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template | <code>""</code>                         |
| podAnnotations             |                                                                                                                        | <code>{}</code>                         |
| podSecurityContext         |                                                                                                                        | <code>{}</code>                         |
| securityContext            |                                                                                                                        | <code>{}</code>                         |
| resources                  |                                                                                                                        | <code>{}</code>                         |
| nodeSelector               |                                                                                                                        | <code>{}</code>                         |
| tolerations                |                                                                                                                        | <code>[]</code>                         |
| affinity                   |                                                                                                                        | <code>{}</code>                         |
| user.name                  |                                                                                                                        | <code>''</code>                         |
| user.email                 |                                                                                                                        | <code>''</code>                         |
| user.product               |                                                                                                                        | <code>console-enterprise</code>         |
| user.tos                   |                                                                                                                        | <code>'true'</code>                     |
| user.token                 |                                                                                                                        | <code>''</code>                         |
| linkID                     |                                                                                                                        | <code>''</code>                         |
| nats.address               |                                                                                                                        | <code>''</code>                         |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install cluster-connector appscode/cluster-connector -n kubeops --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install cluster-connector appscode/cluster-connector -n kubeops --values values.yaml
```
