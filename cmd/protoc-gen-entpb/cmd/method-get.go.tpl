func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	q := s.db.{{ $.EntMsg.Schema.Name }}.Query()
	{{ $key := index $.EntRpc.EntReq.Fields 0 -}}
	{{ $schema := schema $.EntMsg.Schema -}}
	if p, err := Get{{ $.EntMsg.Schema.Name }}Specifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	res, err := Query{{ $.EntMsg.Schema.Name }}WithEdgeIds(q).Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
{{ $query_type := print $.EntMsg.Schema.Name "Query" | $.Ent.Ident | use -}}
func Query{{ $.EntMsg.Schema.Name }}WithEdgeIds(q *{{ $query_type }}) *{{ $query_type }} {
	{{ range $.EntMsg.Fields -}}
		{{ if not .IsEdge -}}
			{{ continue -}}
		{{ end -}}
		
		{{ $name := .EntName | entname -}}
		q.With{{ $name }}(func(q *{{ print .EntMsg.Schema.Name "Query" | $.Ent.Ident | use }}){ q.Select({{ (schema .EntMsg.Schema).Ident "FieldID" | use -}}) })
	{{ end }}

	return q
}
