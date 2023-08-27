# Prometheus Operator CRDs

[Prometheus Operator CRDs](https://github.com/prometheus-operator/prometheus-operator) - Prometheus Operator CRDs

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/prometheus-operator-crds --version=v0.60.1
$ helm upgrade -i prometheus-operator-crds appscode/prometheus-operator-crds -n kube-system --create-namespace --version=v0.60.1
```

## Introduction

This chart deploys Prometheus Operator crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `prometheus-operator-crds`:

```bash
$ helm upgrade -i prometheus-operator-crds appscode/prometheus-operator-crds -n kube-system --create-namespace --version=v0.60.1
```

The command deploys Prometheus Operator crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `prometheus-operator-crds`:

```bash
$ helm uninstall prometheus-operator-crds -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


