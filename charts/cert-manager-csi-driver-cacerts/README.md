# ca-certs CSI driver

[Cert-manager ca-certificates CSI Driver by AppsCode](https://github.com/kubeops/csi-driver-cacerts) - A Helm chart for cert-manager-csi-driver-cacerts

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cert-manager-csi-driver-cacerts --version=v2022.04.04
$ helm upgrade -i cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --create-namespace --version=v2022.04.04
```

## Introduction

This chart deploys a CSI driver on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `cert-manager-csi-driver-cacerts`:

```bash
$ helm upgrade -i cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --create-namespace --version=v2022.04.04
```

The command deploys a CSI driver on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cert-manager-csi-driver-cacerts`:

```bash
$ helm uninstall cert-manager-csi-driver-cacerts -n cert-manager
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cert-manager-csi-driver-cacerts` chart and their default values.

|           Parameter            |                                                            Description                                                             |                                             Default                                              |
|--------------------------------|------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| nameOverride                   |                                                                                                                                    | <code>""</code>                                                                                  |
| fullnameOverride               |                                                                                                                                    | <code>""</code>                                                                                  |
| app.logLevel                   | Verbosity of cert-manager-csi-driver-cacerts logging.                                                                              | <code>2 # 1-5</code>                                                                             |
| app.driver.name                | Name of the driver which will be registered with Kubernetes.                                                                       | <code>cacerts.csi.cert-manager.io</code>                                                         |
| app.livenessProbe.port         | The port that will expose the livness of the csi-driver                                                                            | <code>9809</code>                                                                                |
| registryFQDN                   | Docker registry fqdn used to pull CSI driver image. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | <code>""</code>                                                                                  |
| driver.registry                | Docker registry used to pull node driver image                                                                                     | <code>appscode</code>                                                                            |
| driver.repository              | Name of driver container image                                                                                                     | <code>csi-driver-cacerts</code>                                                                  |
| driver.tag                     | driver container image tag                                                                                                         | <code>""</code>                                                                                  |
| driver.pullPolicy              | Kubernetes imagePullPolicy on the driver container                                                                                 | <code>IfNotPresent</code>                                                                        |
| driver.resources               | Compute Resources required by the driver container                                                                                 | <code>{}</code>                                                                                  |
| nodeDriverRegistrar.registry   | Docker registry used to pull node-driver-registrar image                                                                           | <code>k8s.gcr.io/sig-storage</code>                                                              |
| nodeDriverRegistrar.repository | Name of node-driver-registrar container image                                                                                      | <code>csi-node-driver-registrar</code>                                                           |
| nodeDriverRegistrar.tag        | node-driver-registrar container image tag                                                                                          | <code>v2.3.0</code>                                                                              |
| nodeDriverRegistrar.pullPolicy | Kubernetes imagePullPolicy on node-driver-registrar                                                                                | <code>IfNotPresent</code>                                                                        |
| nodeDriverRegistrar.resources  | Compute Resources required by the node-driver-registrar container                                                                  | <code>{}</code>                                                                                  |
| livenessProbe.registry         | Docker registry used to pull livenessprobe image                                                                                   | <code>k8s.gcr.io/sig-storage</code>                                                              |
| livenessProbe.repository       | Name of livenessprobe container image                                                                                              | <code>livenessprobe</code>                                                                       |
| livenessProbe.tag              | livenessprobe container image tag                                                                                                  | <code>v2.4.0</code>                                                                              |
| livenessProbe.pullPolicy       | Kubernetes imagePullPolicy on liveness probe.                                                                                      | <code>IfNotPresent</code>                                                                        |
| livenessProbe.resources        | Compute Resources required by the livenessprobe container                                                                          | <code>{"limits":{"cpu":"100m","memory":"100Mi"},"requests":{"cpu":"10m","memory":"20Mi"}}</code> |
| imagePullSecrets               |                                                                                                                                    | <code>[]</code>                                                                                  |
| serviceAccount.create          | Specifies whether a service account should be created                                                                              | <code>true</code>                                                                                |
| serviceAccount.annotations     | Annotations to add to the service account                                                                                          | <code>{}</code>                                                                                  |
| serviceAccount.name            | The name of the service account to use. If not set and create is true, a name is generated using the fullname template             | <code>""</code>                                                                                  |
| annotations                    | Annotations applied to csi driver daemonset                                                                                        | <code>{}</code>                                                                                  |
| podAnnotations                 | Annotations passed to csi driver pod(s).                                                                                           | <code>{}</code>                                                                                  |
| tolerations                    | Tolerations for pod assignment                                                                                                     | <code>[]</code>                                                                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --create-namespace --version=v2022.04.04 --set app.logLevel=2 # 1-5
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cert-manager-csi-driver-cacerts appscode/cert-manager-csi-driver-cacerts -n cert-manager --create-namespace --version=v2022.04.04 --values values.yaml
```
