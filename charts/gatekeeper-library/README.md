# Gatekeeper Library

[OPA Gatekeeper Library](https://github.com/open-policy-agent/gatekeeper-library) - OPA Gatekeeper Library for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/gatekeeper-library --version=v2023.10.1
$ helm upgrade -i gatekeeper-library appscode/gatekeeper-library -n kubeops --create-namespace --version=v2023.10.1
```

## Introduction

This chart deploys OPA Gatekeeper Library on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `gatekeeper-library`:

```bash
$ helm upgrade -i gatekeeper-library appscode/gatekeeper-library -n kubeops --create-namespace --version=v2023.10.1
```

The command deploys OPA Gatekeeper Library on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `gatekeeper-library`:

```bash
$ helm uninstall gatekeeper-library -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `gatekeeper-library` chart and their default values.

|     Parameter     | Description |                 Default                  |
|-------------------|-------------|------------------------------------------|
| nameOverride      |             | <code>""</code>                          |
| fullnameOverride  |             | <code>""</code>                          |
| enable            |             | <code>"templates" # "constraints"</code> |
| enableConstraints |             | <code>{}</code>                          |
| enforcementAction |             | <code>"warn" # "deny" "dryrun"</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i gatekeeper-library appscode/gatekeeper-library -n kubeops --create-namespace --version=v2023.10.1 --set enable="templates" # "constraints"
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i gatekeeper-library appscode/gatekeeper-library -n kubeops --create-namespace --version=v2023.10.1 --values values.yaml
```
