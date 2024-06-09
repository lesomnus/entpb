func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	q := s.db.{{ $.EntMsg.Schema.Name }}.Query()
	{{ $key := index $.EntRpc.EntReq.Fields 0 -}}
	{{ $schema := schema $.EntMsg.Schema -}}
	{{ if not $key.IsOneof -}}
	{{ $in := print "req.Get" (print $key.Ident | pascal) "()" -}}
	{{ $pred := print ($key.EntName | entname) "EQ" | $schema.Ident | use  -}}
	{{ print "q.Where(" $pred "(@))" | to_ent $key $in "v" -}}
	{{ else -}}
	switch t := req.GetKey().(type) {
	{{ range $_, $key := $key.Oneof -}}
	{{ $key_name := print $key.Ident | pascal -}}
	{{ $in := print "t." $key_name -}}
	{{ $pred := print ($key.EntName | entname) "EQ" | $schema.Ident | use  -}}
	case *{{ print $.PbMethod.Input.GoIdent.GoName "_" $key_name | $.Pb.Ident | use }}:
		{{ print "q.Where(" $pred "(@))" | to_ent $key $in "v" }}
	{{ end -}}
	default:
		return nil, {{ status "InvalidArgument" "key not provided" }}
	}
	{{ end -}}

	{{ range $.EntMsg.Fields -}}
	{{ if .IsEdge -}}
		{{ $name := .EntName | entname }}
		q.With{{ $name }}(func(q *{{ print .EntRef "Query" | $.Ent.Ident | use }}){ q.Select({{ $schema.Ident "FieldID" | use -}}) })
	{{- end }}
	{{- end }}

	res, err := q.Only(ctx)
	if err != nil {
		return nil, {{ $.Runtime.Ident "EntErrorToStatus" | use }}(err)
	}

	 return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
