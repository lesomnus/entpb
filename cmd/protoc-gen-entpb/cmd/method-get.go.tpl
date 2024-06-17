func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	q := s.db.{{ $.EntMsg.Schema.Name }}.Query()
	{{ $key := index $.EntRpc.EntReq.Fields 0 -}}
	{{ $schema := schema $.EntMsg.Schema -}}
	if p, err := Get{{ $.EntMsg.Schema.Name }}Specifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	{{ range $.EntMsg.Fields -}}
	{{ if .IsEdge -}}
		{{ $name := .EntName | entname }}
		q.With{{ $name }}(func(q *{{ print .EntRef "Query" | $.Ent.Ident | use }}){ q.Select({{ (schema .EntMsg.Schema).Ident "FieldID" | use -}}) })
	{{- end }}
	{{- end }}

	res, err := q.Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	 return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
