{{/* Common names and labels. */}}

{{- define "dr.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "dr.fullname" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "dr.labels" -}}
app.kubernetes.io/name: {{ include "dr.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" }}
{{- end -}}

{{- define "dr.image" -}}
{{- printf "%s:%s" .Values.image.repository (default .Chart.AppVersion .Values.image.tag) -}}
{{- end -}}

{{/*
Comma separated imagePullSecret names for the spoke agent, for the addon-manager's
--agent-image-pull-secrets flag. addon.agent.imagePullSecrets holds plain names;
when it is empty, fall back to the names in the top level imagePullSecrets list of
{name: ...} entries.
*/}}
{{- define "dr.addonAgentPullSecrets" -}}
{{- $names := list -}}
{{- range .Values.addon.agent.imagePullSecrets -}}
{{- $names = append $names . -}}
{{- end -}}
{{- if not $names -}}
{{- range .Values.imagePullSecrets -}}
{{- if .name -}}
{{- $names = append $names .name -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- join "," $names -}}
{{- end -}}

{{/*
Name of the PVC backing the control plane data directory (/.ocm). When
controlplane.persistence.existingClaim is set the chart does not create the PVC and
the Deployment mounts the pre-existing claim by that name; otherwise the chart
creates and mounts <fullname>-controlplane-data. Both the Deployment and the PVC
template resolve the name through this helper so they can never disagree.
*/}}
{{- define "dr.controlplaneClaimName" -}}
{{- with .Values.controlplane.persistence.existingClaim -}}
{{- . -}}
{{- else -}}
{{- printf "%s-controlplane-data" (include "dr.fullname" .) -}}
{{- end -}}
{{- end -}}

{{/*
Access modes for the control plane PVC. Prefers the list form
(controlplane.persistence.accessModes) and falls back to the older singular
controlplane.persistence.accessMode so existing values files keep working.
*/}}
{{- define "dr.controlplaneAccessModes" -}}
{{- $modes := .Values.controlplane.persistence.accessModes -}}
{{- if not $modes -}}
{{- $modes = list (.Values.controlplane.persistence.accessMode | default "ReadWriteOnce") -}}
{{- end -}}
{{- toYaml $modes -}}
{{- end -}}

{{- define "dr.etcdServiceName" -}}
{{- printf "%s-etcd" (include "dr.fullname" .) -}}
{{- end -}}

{{/*
etcd client endpoints. When etcd.deploy is true, enumerate the StatefulSet pods'
stable DNS. Otherwise use the configured external endpoints.
*/}}
{{- define "dr.etcdEndpoints" -}}
{{- if .Values.etcd.deploy -}}
{{- $svc := include "dr.etcdServiceName" . -}}
{{- $ns := .Values.namespace -}}
{{- $eps := list -}}
{{- range $i := until (int .Values.etcd.replicas) -}}
{{- $eps = append $eps (printf "http://%s-%d.%s.%s.svc:2379" $svc $i $svc $ns) -}}
{{- end -}}
{{- join "," $eps -}}
{{- else -}}
{{- join "," .Values.etcd.externalEndpoints -}}
{{- end -}}
{{- end -}}
