# Gatekeeper Grafana Dashboards

[Gatekeeper Grafana Dashboards by AppsCode](https://github.com/kubeops/ui-server) - Gatekeeper Grafana Dashboards for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/gatekeeper-grafana-dashboards --version=v2023.10.1
$ helm upgrade -i gatekeeper-grafana-dashboards appscode/gatekeeper-grafana-dashboards -n kubeops --create-namespace --version=v2023.10.1
```

## Introduction

This chart deploys Gatekeeper Grafana Dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `gatekeeper-grafana-dashboards`:

```bash
$ helm upgrade -i gatekeeper-grafana-dashboards appscode/gatekeeper-grafana-dashboards -n kubeops --create-namespace --version=v2023.10.1
```

The command deploys Gatekeeper Grafana Dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `gatekeeper-grafana-dashboards`:

```bash
$ helm uninstall gatekeeper-grafana-dashboards -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `gatekeeper-grafana-dashboards` chart and their default values.

|            Parameter            |                            Description                             |      Default       |
|---------------------------------|--------------------------------------------------------------------|--------------------|
| nameOverride                    | Overrides name template                                            | <code>""</code>    |
| fullnameOverride                | Overrides fullname template                                        | <code>""</code>    |
| dashboard.enabled               |                                                                    | <code>true</code>  |
| dashboard.folderID              | ID of Grafana folder where these dashboards will be applied        | <code>0</code>     |
| dashboard.overwrite             | If true, dashboard with matching uid will be overwritten           | <code>true</code>  |
| dashboard.templatize.title      | If true, datasource will be prefixed to dashboard name             | <code>false</code> |
| dashboard.templatize.datasource | If true, datasource will be hardcoded in the dashboard             | <code>false</code> |
| grafana.name                    | Name of Grafana Appbinding where these dashboards are applied      | <code>""</code>    |
| grafana.namespace               | Namespace of Grafana Appbinding where these dashboards are applied | <code>""</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i gatekeeper-grafana-dashboards appscode/gatekeeper-grafana-dashboards -n kubeops --create-namespace --version=v2023.10.1 --set dashboard.folderID=0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i gatekeeper-grafana-dashboards appscode/gatekeeper-grafana-dashboards -n kubeops --create-namespace --version=v2023.10.1 --values values.yaml
```
