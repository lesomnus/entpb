{{ $serverName := (print $.Service.GoName "Server")  -}}
func (s *{{ $serverName }}) {{ $.Method.GoName }}(ctx {{ ident "context" "Context" }}, req *{{ goident $.Method.Input.GoIdent }}) (*{{ goident $.Method.Output.GoIdent }}, error) {
	q := s.db.{{ $.Message.Schema.Name }}.Create()
	{{ range $.Message.Fields -}}
	{{ if .IsReadOnly -}}
		{{ continue }}
	{{ end -}}
	{{ $field := print "req." (pascal (print .Ident)) -}}
	{{ $entName := entname .EntName -}}
	{{ $setter := print "q.Set" (entname .EntName) -}}
	{{ if .IsEdge -}}
		{{ if .IsRepeated -}}
		for _, v := range {{ $field }} {
			{{ to_ent . "v.GetId()" "w" (print "q.Add" (.EntName | singular | pascal ) "IDs(@)") }}
		}
		{{ else -}}
		{{ to_ent . (print $field ".GetId()" ) "v" (print $setter "ID(@)") }}
		{{ end -}}
	{{ else if .PbType.IsMessage -}}
	if v := {{ $field }}; v != nil {
		{{ to_ent . "v" "w" (print $setter "(@)") }}
	}
	{{ else if .IsOptional -}}
	if v := {{ $field }}; v != nil {
		{{ to_ent . "*v" "w" (print $setter "(@)") }}
	}
	{{ else -}}
	{{ to_ent . $field "v" (print $setter "(@)") }}
	{{ end -}}
	{{ end -}}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, {{ ident $.Runtime "EntErrorToStatus" }}(err)
	}

	return toProto{{ $.Message.Schema.Name }}(res), nil
}
