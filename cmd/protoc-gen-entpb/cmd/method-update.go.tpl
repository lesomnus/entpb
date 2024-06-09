func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	id, err := {{ import "github.com/google/uuid" | ident "FromBytes" }}(req.GetId())
	if err != nil {
		return nil, {{ import "google.golang.org/grpc/status" | ident "Errorf" }}({{ import "google.golang.org/grpc/codes" | ident "InvalidArgument" }}, "id: %s", err.Error())
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
		{{ print "q.Set" (entname .EntName) "ID(@)" | to_ent . "v.GetId()" "w" }}
		{{ else -}}
		{{ print "q.Set" (entname .EntName) "(@)" | to_ent . "*v" "w" }}
		{{ end -}}
	}
	{{ end -}}
	{{- end }}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, {{ $.Runtime.Ident "EntErrorToStatus" | use }}(err)
	}

	return ToProto{{ $.EntMsg.Schema.Name }}(res), nil
}
