{{ range .Labels }}{{ . }} {{ end }}{{ .Type.String }} {{ .Name }} = {{ .Number -}}
{{- if .Options }} {{ template "field-options.go.tpl" .Options -}}{{- end -}}
;
{{- /**/ -}}
