rpc {{ .Name }} (
{{- with .Request -}}
	{{ if .Stream }}stream {{ end -}}
	{{ .Type -}}
{{- end -}}
) returns (
{{- with .Response -}}
	{{ if .Stream }}stream {{ end -}}
	{{ .Type -}}
{{- end -}}
)
{{- if eq 0 (len .Options) -}}
;
{{- else }} {
	{{ include "options" .Options | indent 1 }}
}
{{- end -}}
