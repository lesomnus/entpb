edition = "{{ .Edition }}";

{{- with .Package }}

package {{ . }};
{{- end }}

{{- range $i, $_ := .Imports }}{{ if eq $i 0 }}
{{ end }}
import {{ with .Visibility }}{{ . }} {{ end }}"{{ .Name }}";
{{- end }}

{{- with .Options }}

{{ template "options.go.tpl" . }}
{{- end }}

{{- range .TopLevelDefinitions }}

{{ include .TemplateName . }}
{{- end }}

{{- /**/ -}}
