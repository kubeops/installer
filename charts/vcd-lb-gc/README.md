# vcd-lb-gc

[VCD Load Balancer Garbage Collector by AppsCode](https://github.com/kubeops/vcd-lb-gc) - Garbage collector for orphaned VMware Cloud Director load-balancer objects

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/vcd-lb-gc --version=v2026.6.25
$ helm upgrade -i vcd-lb-gc appscode/vcd-lb-gc -n vcd --create-namespace --version=v2026.6.25
```

## Introduction

This chart deploys a vcd-lb-gc controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.26+

## Installing the Chart

To install/upgrade the chart with the release name `vcd-lb-gc`:

```bash
$ helm upgrade -i vcd-lb-gc appscode/vcd-lb-gc -n vcd --create-namespace --version=v2026.6.25
```

The command deploys a vcd-lb-gc controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `vcd-lb-gc`:

```bash
$ helm uninstall vcd-lb-gc -n vcd
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `vcd-lb-gc` chart and their default values.

|         Parameter          |                                                                                                              Description                                                                                                               |                                                                                                      Default                                                                                                      |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride               | Overrides name template                                                                                                                                                                                                                | <code>""</code>                                                                                                                                                                                                   |
| fullnameOverride           | Overrides fullname template                                                                                                                                                                                                            | <code>""</code>                                                                                                                                                                                                   |
| replicaCount               | Number of vcd-lb-gc replicas to create. Two replicas + leader election: one active, one hot standby.                                                                                                                                   | <code>2</code>                                                                                                                                                                                                    |
| registryFQDN               | Docker registry fqdn used to pull docker images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                        | <code>ghcr.io</code>                                                                                                                                                                                              |
| image.registry             | Docker registry used to pull the vcd-lb-gc image                                                                                                                                                                                       | <code>appscode</code>                                                                                                                                                                                             |
| image.repository           | Name of the vcd-lb-gc container image                                                                                                                                                                                                  | <code>vcd-lb-gc</code>                                                                                                                                                                                            |
| image.tag                  | vcd-lb-gc container image tag                                                                                                                                                                                                          | <code>""</code>                                                                                                                                                                                                   |
| image.resources            | Compute Resources required by the vcd-lb-gc container                                                                                                                                                                                  | <code>{"limits":{"cpu":"200m","memory":"256Mi"},"requests":{"cpu":"20m","memory":"64Mi"}}</code>                                                                                                                  |
| image.securityContext      | Security options this container should run with                                                                                                                                                                                        | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsGroup":65532,"runAsNonRoot":true,"runAsUser":65532,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| imagePullSecrets           | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/vcd-lb-gc \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                                   |
| imagePullPolicy            | Container image pull policy                                                                                                                                                                                                            | <code>IfNotPresent</code>                                                                                                                                                                                         |
| logLevel                   | Log level for the controller                                                                                                                                                                                                           | <code>2</code>                                                                                                                                                                                                    |
| annotations                | Annotations applied to the deployment                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                                                                                   |
| podAnnotations             | Annotations passed to the controller pod(s).                                                                                                                                                                                           | <code>{}</code>                                                                                                                                                                                                   |
| nodeSelector               | Node labels for pod assignment                                                                                                                                                                                                         | <code>{"kubernetes.io/os":"linux"}</code>                                                                                                                                                                         |
| tolerations                | Tolerations for pod assignment                                                                                                                                                                                                         | <code>[]</code>                                                                                                                                                                                                   |
| affinity                   | Affinity rules for pod assignment                                                                                                                                                                                                      | <code>{}</code>                                                                                                                                                                                                   |
| podSecurityContext         | Security options the controller pod should run with.                                                                                                                                                                                   | <code>{"runAsNonRoot":true,"seccompProfile":{"type":"RuntimeDefault"}}</code>                                                                                                                                     |
| serviceAccount.create      | Specifies whether a service account should be created                                                                                                                                                                                  | <code>true</code>                                                                                                                                                                                                 |
| serviceAccount.annotations | Annotations to add to the service account                                                                                                                                                                                              | <code>{}</code>                                                                                                                                                                                                   |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                 | <code></code>                                                                                                                                                                                                     |
| config.clusterID           | CPI cluster ID, e.g. "capvcdCluster:<uuid>" (required)                                                                                                                                                                                 | <code>""</code>                                                                                                                                                                                                   |
| config.edgeGatewayID       | URN of the edge gateway hosting the LB, e.g. "urn:vcloud:gateway:<uuid>" (required)                                                                                                                                                    | <code>""</code>                                                                                                                                                                                                   |
| config.interval            | Reconcile interval                                                                                                                                                                                                                     | <code>60s</code>                                                                                                                                                                                                  |
| config.dryRun              | If true, logs orphans without deleting them. Flip to false to enable deletions once validated.                                                                                                                                         | <code>true</code>                                                                                                                                                                                                 |
| config.skipDNAT            | Skip DNAT cleanup (set when enableVirtualServiceSharedIP is on)                                                                                                                                                                        | <code>false</code>                                                                                                                                                                                                |
| vcd.endpoint               | VCD URL, e.g. https://vcd.example.com                                                                                                                                                                                                  | <code>""</code>                                                                                                                                                                                                   |
| vcd.org                    | VCD tenant org name                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                                   |
| vcd.user                   | VCD tenant user                                                                                                                                                                                                                        | <code>""</code>                                                                                                                                                                                                   |
| vcd.password               | VCD tenant password                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                                   |
| vcd.insecure               | Skip TLS verification when talking to VCD                                                                                                                                                                                              | <code>false</code>                                                                                                                                                                                                |
| vcd.existingSecret         | Use an existing Secret holding VCD_ENDPOINT/VCD_ORG/VCD_USER/VCD_PASSWORD keys instead of creating one from the values above.                                                                                                          | <code>""</code>                                                                                                                                                                                                   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i vcd-lb-gc appscode/vcd-lb-gc -n vcd --create-namespace --version=v2026.6.25 --set replicaCount=2
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i vcd-lb-gc appscode/vcd-lb-gc -n vcd --create-namespace --version=v2026.6.25 --values values.yaml
```
