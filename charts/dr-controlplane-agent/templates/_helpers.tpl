{{/* Common names and labels for the dr-controlplane agent addon chart. */}}

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
The namespace every rendered object lands in. addonInstallNamespace is a built in
value the addon-framework sets to the effective addon install namespace and that
getValuesFuncs cannot override, so prefer it and fall back to .Values.namespace
for a standalone helm render.
*/}}
{{- define "dr.namespace" -}}
{{- default .Values.namespace .Values.addonInstallNamespace -}}
{{- end -}}
