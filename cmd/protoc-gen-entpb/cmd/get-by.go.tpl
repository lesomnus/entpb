{{ $fields := $.EntMsg.Fields -}}
{{ $isOneof := (index $fields 0).IsOneof -}}
{{ if $isOneof -}}
{{ $fields = (index $fields 0).Oneof -}}
{{ end -}}
{{ $msg_type := print $.EntMsg.Ident | $.Pb.Ident | use -}}
{{ if not $isOneof -}}
	{{ $field := index $fields 0 -}}
	func {{ $.EntMsg.Schema.Name }}ById(k {{ ent_type $field }}) *{{ $msg_type }} {
		return &{{ $msg_type }}{Id: {{ to_pb_v $field "k" }}}
	}
	{{ if not (is_symmetric $field) -}}
		func {{ $.EntMsg.Schema.Name }}ByIdV(k {{ pb_type $field }}) *{{ $msg_type }} {
			return &{{ $msg_type }}{Id: k}
		}
	{{ end -}}
{{ else -}}
	{{ range $fields -}}
		{{ if .IsTypeMessage -}}
			{{ continue -}}
		{{ end -}}

		{{ $key_name := print .Ident | pascal -}}
		func {{ $.EntMsg.Schema.Name }}By{{ $key_name }}(k {{ ent_type . }}) *{{ $msg_type }} {
			return &{{ $msg_type }}{Key: &{{ $msg_type }}_{{ $key_name }}{ {{ $key_name }}: {{ to_pb_v . "k" }} }}
		}
		{{ if not (is_symmetric .) -}}
			func {{ $.EntMsg.Schema.Name }}By{{ $key_name }}V(k {{ pb_type . }}) *{{ $msg_type }} {
				return &{{ $msg_type }}{Key: &{{ $msg_type }}_{{ $key_name }}{ {{ $key_name }}: k }}
			}
		{{ end -}}
	{{ end -}}
{{ end -}}
