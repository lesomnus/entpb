{{ range .Labels }}{{ . }} {{ end }}{{ .Type.String }} {{ .Name }} = {{ .Number -}}
{{- with .Options }} {{ template "field-options.go.tpl" . -}}{{- end -}}
;
{{- /**/ -}}
