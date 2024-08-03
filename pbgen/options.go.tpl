{{ range $i, $_ := . }}{{ if ne $i 0 }}
{{ end -}}
option {{ .Name }} = {{ with .Value }}{{ include .TemplateName . }};{{ end -}}
{{- end -}}
