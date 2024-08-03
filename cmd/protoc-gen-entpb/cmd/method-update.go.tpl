func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	id, err := Get{{ $.EntMsg.Schema.Name }}Id(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.{{ $.EntMsg.Schema.Name }}.UpdateOneID(id)
	{{ range slice $.EntRpc.EntReq.Fields 1 -}}
		{{ if or .IsImmutable .IsReadOnly -}}
			{{ continue }}	
		{{ end -}}

		{{ if .IsRepeated -}}
			{{ $field := print "req.Get" (print .Ident | plural | pascal) "()" -}}
			for _, v := range {{ $field }} {
				{{ print "q.Add" (singular .EntName | entname) "IDs(@)" | to_ent . "v.GetId()" "w" }}
			}
		{{ else -}}
			{{ $field := print "req." (print .Ident | pascal) -}}
			if v := {{ $field }}; v != nil {
				{{ if .IsEdge -}}
				if id, err := Get{{ .EntMsg.Schema.Name }}Id(ctx, s.db, {{ $field }}); err != nil {
					return nil, err
				} else {
					q.Set{{ entname .EntName }}ID(id)
				}
				{{ else -}}
					{{ $in := "v" -}}
					{{ if and .IsDereferenceable .IsOptional -}}
						{{ $in = "*v" -}}
					{{ end -}}
					{{ print "q.Set" (entname .EntName) "(@)" | to_ent . $in "w" }}
				{{ end -}}
			}
		{{ end -}}
	{{- end }}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
