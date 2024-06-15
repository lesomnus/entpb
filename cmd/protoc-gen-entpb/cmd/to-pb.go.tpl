{{ $pb_type := print $.EntMsg.Ident | $.Pb.Ident | use -}}
func ToProto{{ $.EntMsg.Schema.Name }}(v *{{ $.Ent.Ident $.EntMsg.Schema.Name | use }}) *{{ $pb_type }} {
	m := &{{ $pb_type }}{}
	{{ range $.EntMsg.Fields -}}
	{{ $pb_field := print "m." (pascal (print .Ident)) -}}
	{{ $ent_field := print "v." (entname .EntName) -}}
	{{ if .IsEdge -}}
		{{ $ent_field = print "v.Edges." (entname .EntName) -}}
		{{ $converter := print "ToProto" .EntRef -}} 
		{{ if .IsRepeated -}}
		for _, v := range {{ $ent_field }} {
			{{ $pb_field }} = append({{ $pb_field }}, {{ $converter }}(v))
		}
		{{ else -}}
		if v := {{ $ent_field }}; v != nil {
			{{ $pb_field }} = {{ $converter }}(v)
		}
		{{ end -}}
	{{ else -}}
	{{ to_pb . $ent_field $pb_field }}
	{{ end -}}
	{{ end -}}

	return m
}
