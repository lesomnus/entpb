{{ if ne "" .Comment -}}
// {{ .Comment }}
{{- end }}
message {{ .Name }} {
{{- range .Fields }}
	{{- if ne "" .Comment }}
	// {{ .Comment }}
	{{- end }}
	{{ if .IsRepeated}}repeated {{ end -}}
	{{ .Type.Name }} {{ .Name }} = {{ .Number -}}
	{{- if and .IsOptional (not .IsRepeated) }} [features.field_presence = EXPLICIT]{{ end }};
{{- end }}
}
