# Kmodules CRDs

[Kmodules CRDs](https://github.com/Kmodules) - Shared Kubernetes Custom Resource Definitions for AppsCode Projects

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kmodules-crds --version=v0.25.1
$ helm upgrade -i kmodules-crds appscode/kmodules-crds -n kube-system --create-namespace --version=v0.25.1
```

## Introduction

This chart deploys Kmodules crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kmodules-crds`:

```bash
$ helm upgrade -i kmodules-crds appscode/kmodules-crds -n kube-system --create-namespace --version=v0.25.1
```

The command deploys Kmodules crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kmodules-crds`:

```bash
$ helm uninstall kmodules-crds -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


