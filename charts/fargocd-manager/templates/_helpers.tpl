{{/*
Expand the name of the chart.
*/}}
{{- define "fargocd-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "fargocd-manager.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "fargocd-manager.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "fargocd-manager.labels" -}}
helm.sh/chart: {{ include "fargocd-manager.chart" . }}
{{ include "fargocd-manager.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "fargocd-manager.selectorLabels" -}}
app.kubernetes.io/name: {{ include "fargocd-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Returns whether the OpenShift distribution is used
*/}}
{{- define "distro.openshift" -}}
{{- or (.Capabilities.APIVersions.Has "project.openshift.io/v1/Project") .Values.distro.openshift -}}
{{- end }}

{{/*
Returns the raw Argo CD principal kubeconfig
*/}}
{{- define "argocd.kubeconfig" -}}
{{- .Values.argocd.kubeconfig }}
{{- end }}

{{/*
Returns the name of the Secret on the hub holding the Argo CD kubeconfig.
Prefers an existing Secret named via argocd.kubeconfigSecret; otherwise
derives a chart-owned name when argocd.kubeconfig content is provided.
*/}}
{{- define "argocd.kubeconfigSecretName" -}}
{{- if .Values.argocd.kubeconfigSecret }}
{{- .Values.argocd.kubeconfigSecret -}}
{{- else if .Values.argocd.kubeconfig }}
{{- printf "%s-argo-kubeconfig" (include "fargocd-manager.fullname" .) -}}
{{- end }}
{{- end }}
