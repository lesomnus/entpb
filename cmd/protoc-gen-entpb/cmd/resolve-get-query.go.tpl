{{ $fields := $.EntMsg.Fields -}}
{{ $isOneof := (index $fields 0).IsOneof -}}
{{ if $isOneof -}}
{{ $fields = (index $fields 0).Oneof -}}
{{ end -}}
{{ $func_name := print "ResolveGet" $.EntMsg.Schema.Name "Query" }}
{{ $get_req := print $.EntMsg.Ident | $.Pb.Ident | use }}
func {{ $func_name }}(req *{{ $get_req }}) (*{{ $get_req }}, error) {
	t, ok := req.Key.(*{{ print $.EntMsg.Ident "_Query" | $.Pb.Ident | use }})
	if !ok {
		return req, nil
	}

	q := t.Query
	{{/* Expect first field to be "id" anGet{{ $.EntMsg.Schema.Name }}Specifierd last field to be "query" */}}
	{{ range slice (reverse_fields (slice $fields 1)) 1 -}}
		{{ if .IsTypeMessage -}}
			{{/* skip indices */}}
			{{ continue -}}
		{{ end -}}
		if v, ok := {{ import "strings" | ident "CutPrefix" }}(q, {{ printf "%q" .QueryPrefix }}); ok {
			return {{ print $.EntMsg.Schema.Name "By" (print .Ident | pascal) | $.Pb.Ident | use }}(v), nil
		}
	{{ end -}}

	{{ $id_field := index $fields 0 -}}
	{{ if eq $id_field.EntInfo.Type 4 -}}
		v, err := {{ import "github.com/google/uuid" | ident "Parse" }}(q)
		if err != nil {
			return nil, {{ status_errf "InvalidArgument" "invalid query string: %s" "err" }}
		}
		return {{ print $.EntMsg.Schema.Name "ById" | $.Pb.Ident | use }}(v), nil
	{{ else -}}
		return {{ print $.EntMsg.Schema.Name "ById" | $.Pb.Ident | use }}(q), nil
	{{ end -}}
}
