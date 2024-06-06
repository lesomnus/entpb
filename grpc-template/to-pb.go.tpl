func toProto{{ $.Schema.Name }}(v *{{ ident $.Ent (print $.Schema.Name) }}) *{{ ident $.Pb (print $.Ident) }} {
	m := &{{ ident $.Pb (print .Ident) }}{}
	{{ range $.Fields -}}
	{{ $pb_field := print "m." (pascal (print .Ident)) -}}
	{{ $ent_field := print "v." (entname .EntName) -}}
	{{ if .IsEdge -}}
		{{ $ent_field = print "v.Edges." (entname .EntName) -}}
		{{ $ref_type := ident $.Pb (print .PbType.Name) -}}
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
	{{ $pb_field }} = {{ to_pb . $ent_field }}
	{{ end -}}
	{{ end -}}

	return m
}
