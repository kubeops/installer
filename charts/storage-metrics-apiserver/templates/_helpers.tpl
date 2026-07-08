{{/*
Expand the name of the chart.
*/}}
{{- define "storage-metrics-apiserver.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Fully qualified app name.
*/}}
{{- define "storage-metrics-apiserver.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "storage-metrics-apiserver.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "storage-metrics-apiserver.labels" -}}
helm.sh/chart: {{ include "storage-metrics-apiserver.chart" . }}
{{ include "storage-metrics-apiserver.selectorLabels" . }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "storage-metrics-apiserver.selectorLabels" -}}
app.kubernetes.io/name: {{ include "storage-metrics-apiserver.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "storage-metrics-apiserver.serviceAccountName" -}}
{{ include "storage-metrics-apiserver.fullname" . }}
{{- end -}}

{{/*
Returns the registry used for the storage-metrics-apiserver docker image.
*/}}
{{- define "storage-metrics-apiserver.image.registry" -}}
{{- list .Values.registryFQDN .Values.image.registry | compact | join "/" }}
{{- end -}}
