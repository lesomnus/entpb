{{ if eq 1 (len .) }}
{{- with index . 0 -}}
[{{ .Name }} = {{ .Value }}]
{{- end -}}
{{- else -}}
[
{{- range $i, $_ := . }}{{ if ne $i 0 }},{{ end }}
	{{ .Name }} = {{ .Value }}
{{- end }}
]
{{- end -}}
