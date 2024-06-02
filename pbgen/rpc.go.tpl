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
);
{{- /**/ -}}
