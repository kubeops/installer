# Application CRDs

[Application CRDs](https://github.com/kmodules/application) - Kubernetes Application CRDs

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/application-crds --version=v0.8.3
$ helm upgrade -i application-crds appscode/application-crds -n kube-system --create-namespace --version=v0.8.3
```

## Introduction

This chart deploys Application crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `application-crds`:

```bash
$ helm upgrade -i application-crds appscode/application-crds -n kube-system --create-namespace --version=v0.8.3
```

The command deploys Application crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `application-crds`:

```bash
$ helm uninstall application-crds -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


