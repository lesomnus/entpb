{{ range $i, $_ := . }}{{ if ne $i 0 }}
{{ end -}}
option {{ .Name }} = {{ .Value }};
{{- end -}}
