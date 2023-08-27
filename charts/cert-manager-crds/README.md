# cert-manager CRDs

[cert-manager CRDs](https://cert-manager.io) - Jetstack cert-manager CRDs

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cert-manager-crds --version=1.9.1
$ helm upgrade -i cert-manager-crds appscode/cert-manager-crds -n kube-system --create-namespace --version=1.9.1
```

## Introduction

This chart deploys cert-manager crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `cert-manager-crds`:

```bash
$ helm upgrade -i cert-manager-crds appscode/cert-manager-crds -n kube-system --create-namespace --version=1.9.1
```

The command deploys cert-manager crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cert-manager-crds`:

```bash
$ helm uninstall cert-manager-crds -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


