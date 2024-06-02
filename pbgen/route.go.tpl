{{- with .TemplateName -}}
{{- if eq . "comment" -}}
{{- template "comment.go.tpl" $ -}}
{{- else if eq . "message-field" -}}
{{- template "message-field.go.tpl" $ -}}
{{- else if eq . "message" -}}
{{- template "message.go.tpl" $ -}}
{{- end -}}
{{- end -}}
