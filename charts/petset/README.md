# Petset Operator

[Petset Operator by AppsCode](https://github.com/kubeops/petset) - Petset Operator by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/petset --version=v2025.7.31
$ helm upgrade -i petset appscode/petset -n kubeops --create-namespace --version=v2025.7.31
```

## Introduction

This chart deploys Petset operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `petset`:

```bash
$ helm upgrade -i petset appscode/petset -n kubeops --create-namespace --version=v2025.7.31
```

The command deploys Petset operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `petset`:

```bash
$ helm uninstall petset -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `petset` chart and their default values.

|             Parameter             |                                                                                                             Description                                                                                                             |                                                                                            Default                                                                                             |
|-----------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                      | Overrides name template                                                                                                                                                                                                             | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                  | Overrides fullname template                                                                                                                                                                                                         | <code>""</code>                                                                                                                                                                                |
| replicaCount                      | Number of stash operator replicas to create (only 1 is supported)                                                                                                                                                                   | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                      | Docker registry fqdn used to pull Stash related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                              | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                    | Docker registry used to pull operator image                                                                                                                                                                                         | <code>appscode</code>                                                                                                                                                                          |
| image.repository                  | Name of operator container image                                                                                                                                                                                                    | <code>petset</code>                                                                                                                                                                            |
| image.tag                         | Operator container image tag                                                                                                                                                                                                        | <code>""</code>                                                                                                                                                                                |
| image.resources                   | Compute Resources required by the operator container                                                                                                                                                                                | <code>{"requests":{"cpu":"100m"}}</code>                                                                                                                                                       |
| image.securityContext             | Security options this container should run with                                                                                                                                                                                     | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| imagePullSecrets                  | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/petset \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                |
| imagePullPolicy                   | Container image pull policy                                                                                                                                                                                                         | <code>IfNotPresent</code>                                                                                                                                                                      |
| criticalAddon                     | If true, installs petset as critical addon                                                                                                                                                                                          | <code>false</code>                                                                                                                                                                             |
| logLevel                          | Log level for operator                                                                                                                                                                                                              | <code>3</code>                                                                                                                                                                                 |
| annotations                       | Annotations applied to operator deployment                                                                                                                                                                                          | <code>{}</code>                                                                                                                                                                                |
| podAnnotations                    | Annotations passed to operator pod(s).                                                                                                                                                                                              | <code>{}</code>                                                                                                                                                                                |
| podLabels                         | Labels passed to operator pod(s)                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| nodeSelector                      | Node labels for pod assignment                                                                                                                                                                                                      | <code>{"kubernetes.io/os":"linux"}</code>                                                                                                                                                      |
| tolerations                       | Tolerations for pod assignment                                                                                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| affinity                          | Affinity rules for pod assignment                                                                                                                                                                                                   | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext                | Security options the operator pod should run with.                                                                                                                                                                                  | <code>{"fsGroup":65535}</code>                                                                                                                                                                 |
| serviceAccount.create             | Specifies whether a service account should be created                                                                                                                                                                               | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations        | Annotations to add to the service account                                                                                                                                                                                           | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name               | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                              | <code></code>                                                                                                                                                                                  |
| apiserver.enableMutatingWebhook   | If true, mutating webhook is configured for petset CRDs                                                                                                                                                                             | <code>true</code>                                                                                                                                                                              |
| apiserver.enableValidatingWebhook | If true, validating webhook is configured for petset CRDs                                                                                                                                                                           | <code>true</code>                                                                                                                                                                              |
| apiserver.healthcheck.enabled     | If true, enables the readiness and liveliness probes for the operator pod.                                                                                                                                                          | <code>false</code>                                                                                                                                                                             |
| apiserver.servingCerts.generate   | If true, generates on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate operators pods. Otherwise specify certs in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`.   | <code>true</code>                                                                                                                                                                              |
| apiserver.servingCerts.caCrt      | CA certficate used by serving certificate of webhook server.                                                                                                                                                                        | <code>""</code>                                                                                                                                                                                |
| apiserver.servingCerts.serverCrt  | Serving certficate used by webhook server.                                                                                                                                                                                          | <code>""</code>                                                                                                                                                                                |
| apiserver.servingCerts.serverKey  | Private key for the serving certificate used by webhook server.                                                                                                                                                                     | <code>""</code>                                                                                                                                                                                |
| monitoring.agent                  | Name of monitoring agent (either "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                               | <code>"none"</code>                                                                                                                                                                            |
| monitoring.serviceMonitor.labels  | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                 | <code>{}</code>                                                                                                                                                                                |
| networkPolicy.enabled             |                                                                                                                                                                                                                                     | <code>false</code>                                                                                                                                                                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i petset appscode/petset -n kubeops --create-namespace --version=v2025.7.31 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i petset appscode/petset -n kubeops --create-namespace --version=v2025.7.31 --values values.yaml
```
