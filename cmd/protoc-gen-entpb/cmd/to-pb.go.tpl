{{ $pb_type := print $.EntMsg.Ident | $.Pb.Ident | use -}}
func ToProto{{ $.EntMsg.Schema.Name }}(v *{{ $.Ent.Ident $.EntMsg.Schema.Name | use }}) *{{ $pb_type }} {
	m := &{{ $pb_type }}{}
	{{ range $.EntMsg.Fields -}}
	{{ $pb_field := print "m." (pascal (print .Ident)) -}}
	{{ $ent_field := print "v." (entname .EntName) -}}
	{{ if .IsEdge -}}
		{{ $ent_field = print "v.Edges." (entname .EntName) -}}
		{{ $ref_type := print .PbType.Ident | $.Pb.Ident | use -}}
		{{ if .IsRepeated -}}
		for _, v := range {{ $ent_field }} {
			{{ $pb_field }} = append({{ $pb_field }}, &{{ $ref_type }}{Id: v.ID[:]})
		}
		{{ else -}}
		if v := {{ $ent_field }}; v != nil {
			{{ $pb_field }} = &{{ $ref_type }}{Id: v.ID[:]}
		}
		{{ end -}}
	{{ else -}}
	{{ to_pb . $ent_field $pb_field }}
	{{ end -}}
	{{ end -}}

	return m
}
