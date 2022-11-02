{{- define "chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "chart.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}

{{- define "chart.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "chart.selectorLabels" -}}
app.kubernetes.io/name: tls-hacking
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/component: server
app.kubernetes.io/part-of: tls-hacking
{{- end -}}

{{- define "chart.labels" -}}
app.kubernetes.io/name: {{ include "chart.name" . }}
helm.sh/chart: {{ include "chart.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: "{{ include "chart.app-version" . }}"
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
chart: {{ .Chart.Name | quote }}
{{- end -}}


{{- define "chart.app-version" -}}
{{- printf "%s" .Chart.AppVersion | replace "+" "_" | replace ":" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}



{{- define "test" -}}
a-test
{{- end -}}