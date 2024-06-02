edition = "2023";

{{ range .Imports -}}
import public "{{ . }}";
{{- end }}

{{ if ne "" .GoPackage -}}
option go_package = "{{ .GoPackage }}";
{{- end }}
