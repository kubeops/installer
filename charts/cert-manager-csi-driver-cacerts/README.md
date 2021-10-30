# ca-certs CSI driver

[Cert-manager ca-certificates CSI Driver by AppsCode](https://github.com/kubeops/csi-driver-cacerts) - A Helm chart for cert-manager-csi-driver-cacerts

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager
```

## Introduction

This chart deploys a CSI driver on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `cert-manager-csi-driver-cacerts`:

```console
$ helm install cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager
```

The command deploys a CSI driver on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `cert-manager-csi-driver-cacerts`:

```console
$ helm delete cert-manager-csi-driver-cacerts -n cert-manager
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cert-manager-csi-driver-cacerts` chart and their default values.

|           Parameter            |                                                            Description                                                             |                                        Default                                        |
|--------------------------------|------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| nameOverride                   |                                                                                                                                    | `""`                                                                                  |
| fullnameOverride               |                                                                                                                                    | `""`                                                                                  |
| app.logLevel                   | Verbosity of cert-manager-csi-driver-cacerts logging.                                                                              | `2 # 1-5`                                                                             |
| app.driver.name                | Name of the driver which will be registered with Kubernetes.                                                                       | `cacerts.csi.cert-manager.io`                                                         |
| app.livenessProbe.port         | The port that will expose the livness of the csi-driver                                                                            | `9809`                                                                                |
| registryFQDN                   | Docker registry fqdn used to pull CSI driver image. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | `""`                                                                                  |
| driver.registry                | Docker registry used to pull node driver image                                                                                     | `appscode`                                                                            |
| driver.repository              | Name of driver container image                                                                                                     | `csi-driver-cacerts`                                                                  |
| driver.tag                     | driver container image tag                                                                                                         | `v0.0.1`                                                                              |
| driver.pullPolicy              | Kubernetes imagePullPolicy on the driver container                                                                                 | `IfNotPresent`                                                                        |
| driver.resources               | Compute Resources required by the driver container                                                                                 | `{}`                                                                                  |
| nodeDriverRegistrar.registry   | Docker registry used to pull node-driver-registrar image                                                                           | `k8s.gcr.io/sig-storage`                                                              |
| nodeDriverRegistrar.repository | Name of node-driver-registrar container image                                                                                      | `csi-node-driver-registrar`                                                           |
| nodeDriverRegistrar.tag        | node-driver-registrar container image tag                                                                                          | `v2.3.0`                                                                              |
| nodeDriverRegistrar.pullPolicy | Kubernetes imagePullPolicy on node-driver-registrar                                                                                | `IfNotPresent`                                                                        |
| nodeDriverRegistrar.resources  | Compute Resources required by the node-driver-registrar container                                                                  | `{}`                                                                                  |
| livenessProbe.registry         | Docker registry used to pull livenessprobe image                                                                                   | `k8s.gcr.io/sig-storage`                                                              |
| livenessProbe.repository       | Name of livenessprobe container image                                                                                              | `livenessprobe`                                                                       |
| livenessProbe.tag              | livenessprobe container image tag                                                                                                  | `v2.4.0`                                                                              |
| livenessProbe.pullPolicy       | Kubernetes imagePullPolicy on liveness probe.                                                                                      | `IfNotPresent`                                                                        |
| livenessProbe.resources        | Compute Resources required by the livenessprobe container                                                                          | `{"limits":{"cpu":"100m","memory":"100Mi"},"requests":{"cpu":"10m","memory":"20Mi"}}` |
| imagePullSecrets               |                                                                                                                                    | `[]`                                                                                  |
| serviceAccount.create          | Specifies whether a service account should be created                                                                              | `true`                                                                                |
| serviceAccount.annotations     | Annotations to add to the service account                                                                                          | `{}`                                                                                  |
| serviceAccount.name            | The name of the service account to use. If not set and create is true, a name is generated using the fullname template             | `""`                                                                                  |
| annotations                    | Annotations applied to csi driver daemonset                                                                                        | `{}`                                                                                  |
| podAnnotations                 | Annotations passed to csi driver pod(s).                                                                                           | `{}`                                                                                  |
| tolerations                    | Tolerations for pod assignment                                                                                                     | `[]`                                                                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --set app.logLevel=2 # 1-5
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --values values.yaml
```
