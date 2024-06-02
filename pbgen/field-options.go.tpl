[
{{- range $i, $_ := . }}{{ if ne $i 0 }},{{ end }}
	{{ .Name }} = {{ .Value }}
{{- end }}
]
{{- /**/ -}}
