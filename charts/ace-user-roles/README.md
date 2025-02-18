# ACE User Roles

[ACE User Roles by AppsCode](https://github.com/kubeops/ui-server) - ACE User Roles for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/ace-user-roles --version=v2024.9.30
$ helm upgrade -i ace-user-roles appscode/ace-user-roles -n kubeops --create-namespace --version=v2024.9.30
```

## Introduction

This chart deploys ACE User Roles on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `ace-user-roles`:

```bash
$ helm upgrade -i ace-user-roles appscode/ace-user-roles -n kubeops --create-namespace --version=v2024.9.30
```

The command deploys ACE User Roles on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `ace-user-roles`:

```bash
$ helm uninstall ace-user-roles -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `ace-user-roles` chart and their default values.

|               Parameter                |         Description         |               Default                |
|----------------------------------------|-----------------------------|--------------------------------------|
| nameOverride                           | Overrides name template     | <code>""</code>                      |
| fullnameOverride                       | Overrides fullname template | <code>""</code>                      |
| enableClusterRoles.ace                 |                             | <code>true</code>                    |
| enableClusterRoles.appcatalog          |                             | <code>true</code>                    |
| enableClusterRoles.catalog             |                             | <code>true</code>                    |
| enableClusterRoles.cert-manager        |                             | <code>true</code>                    |
| enableClusterRoles.kubedb              |                             | <code>true</code>                    |
| enableClusterRoles.kubedb-ui           |                             | <code>true</code>                    |
| enableClusterRoles.kubestash           |                             | <code>true</code>                    |
| enableClusterRoles.kubevault           |                             | <code>true</code>                    |
| enableClusterRoles.license-proxyserver |                             | <code>true</code>                    |
| enableClusterRoles.metrics             |                             | <code>true</code>                    |
| enableClusterRoles.prometheus          |                             | <code>true</code>                    |
| enableClusterRoles.stash               |                             | <code>true</code>                    |
| annotations.helm.sh/hook               |                             | <code>pre-install,pre-upgrade</code> |
| annotations.helm.sh/hook-delete-policy |                             | <code>before-hook-creation</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i ace-user-roles appscode/ace-user-roles -n kubeops --create-namespace --version=v2024.9.30 --set annotations.helm.sh/hook=pre-install,pre-upgrade
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i ace-user-roles appscode/ace-user-roles -n kubeops --create-namespace --version=v2024.9.30 --values values.yaml
```
