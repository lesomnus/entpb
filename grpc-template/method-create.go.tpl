func (s *{{ print $.Service.GoName "Server" }}) {{ $.Method.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.Method.Input.GoIdent | use }}) (*{{ $.Method.Output.GoIdent | use }}, error) {
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
	{{ else if .IsEnum -}}
	{{ to_ent . ($field) "w" (print $setter "(@)") }}
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
		return nil, {{ $.Runtime.Ident "EntErrorToStatus" | use }}(err)
	}

	return toProto{{ $.Message.Schema.Name }}(res), nil
}
