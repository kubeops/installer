{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "fargocd.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "fargocd.fullname" -}}
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
{{- define "fargocd.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "fargocd.labels" -}}
helm.sh/chart: {{ include "fargocd.chart" . }}
{{ include "fargocd.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "fargocd.selectorLabels" -}}
app.kubernetes.io/name: {{ include "fargocd.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "fargocd.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "fargocd.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- .Values.license }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "image.registry" -}}
{{- list .Values.registryFQDN .Values.image.registry | compact | join "/" }}
{{- end }}

{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{/*
Returns the enabled monitoring agent name
*/}}
{{- define "monitoring.agent" -}}
{{- .Values.monitoring.agent }}
{{- end }}

{{/*
Returns whether the ServiceMonitor will be labeled with custom label
*/}}
{{- define "monitoring.apply-servicemonitor-label" -}}
{{- ternary "false" "true" ( empty .Values.monitoring.serviceMonitor.labels ) -}}
{{- end }}

{{/*
Returns the ServiceMonitor labels
*/}}
{{- define "monitoring.servicemonitor-label" -}}
{{- range $key, $val := .Values.monitoring.serviceMonitor.labels }}
{{ $key }}: {{ $val }}
{{- end }}
{{- end }}

{{/*
Prepare certs
*/}}
{{- define "fargocd.prepare-certs" -}}
{{- if not ._caCrt }}
{{- $caCrt := "" }}
{{- $serverCrt := "" }}
{{- $serverKey := "" }}
{{- if .Values.apiserver.servingCerts.generate }}
{{- $ca := genCA "ca" 3650 }}
{{- $cn := include "fargocd.fullname" . -}}
{{- $altName1 := printf "%s.%s" $cn .Release.Namespace }}
{{- $altName2 := printf "%s.%s.svc" $cn .Release.Namespace }}
{{- $server := genSignedCert $cn nil (list $altName1 $altName2) 3650 $ca }}
{{- $caCrt =  b64enc $ca.Cert }}
{{- $serverCrt = b64enc $server.Cert }}
{{- $serverKey = b64enc $server.Key }}
{{- else }}
{{- $caCrt = required "Required when apiserver.servingCerts.generate is false" .Values.apiserver.servingCerts.caCrt }}
{{- $serverCrt = required "Required when apiserver.servingCerts.generate is false" .Values.apiserver.servingCerts.serverCrt }}
{{- $serverKey = required "Required when apiserver.servingCerts.generate is false" .Values.apiserver.servingCerts.serverKey }}
{{- end }}

{{ $_ := set $ "_caCrt" $caCrt }}
{{ $_ := set $ "_serverCrt" $serverCrt }}
{{ $_ := set $ "_serverKey" $serverKey }}

{{- end }}
{{- end }}
