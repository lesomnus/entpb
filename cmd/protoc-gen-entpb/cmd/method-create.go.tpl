func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	q := s.db.{{ $.EntMsg.Schema.Name }}.Create()
	{{ range $.EntRpc.EntReq.Fields -}}
		{{ if and .IsReadOnly .IsOptional -}}
			{{ continue }}
		{{ end -}}
		{{ $field := print "req.Get" (print .Ident | pascal) "()" -}}
		{{ $entName := entname .EntName -}}
		{{ $setter := print "q.Set" (entname .EntName) -}}
		{{ if .IsEdge -}}
			{{ if .IsRepeated -}}
				for _, v := range {{ $field }} {
					{{ to_ent . "v.GetId()" "w" (print "q.Add" (.EntName | singular | pascal ) "IDs(@)") }}
				}
			{{ else if .IsRequired -}}
				{{/* TODO: required and repeated? should it have at least one element? */ -}}
				if v, err := Get{{ .EntMsg.Schema.Name }}Id(ctx, s.db, req.Get{{ print .Ident | pascal }}()); err != nil {
					return nil, err
				} else {
					{{ $setter }}ID(v)
				}
			{{ else -}}
				if v := {{ $field }}; v != nil {
					{{ to_ent . "v.GetId()" "w" (print $setter "ID(@)") }}
				}
			{{ end -}}
		{{ else if .IsEnum -}}
			{{ to_ent . $field "w" (print $setter "(@)") }}
		{{ else if .PbType.IsMessage -}}
			if v := {{ $field }}; v != nil {
				{{ to_ent . "v" "w" (print $setter "(@)") }}
			}
		{{ else if .IsOptional -}}
			if v := {{ print "req." (print .Ident | pascal) }}; v != nil {
				{{ to_ent . "*v" "w" (print $setter "(@)") }}
			}
		{{ else -}}
			{{ to_ent . $field "v" (print $setter "(@)") }}
		{{ end -}}
	{{- end }}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
