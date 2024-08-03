{{ $fields := $.EntMsg.Fields -}}
{{ $isOneof := (index $fields 0).IsOneof -}}
{{ if $isOneof -}}
{{ $fields = (index $fields 0).Oneof -}}
{{ end -}}
{{ $id_field := index $fields 0 -}}
{{ range $fields -}}
	{{ if eq .Ident "id" -}}
		{{/* ID field may not be the first field */ -}}
		{{ $id_field = . -}}
		{{ break -}}
	{{ end -}}
{{ end -}}
{{ $id_type := ent_type $id_field -}}
func Get{{ $.EntMsg.Schema.Name }}Id(ctx {{ import "context" | ident "Context" }}, db *{{ $.Ent.Ident "Client" | use }}, req *{{ print $.EntMsg.Ident | $.Pb.Ident | use }}) ({{ $id_type }}, error) {
	var r {{ $id_type }}
	{{ if not $isOneof -}}
		{{ to_ent_with_rv $id_field "req.GetId()" "v" "r = @\nreturn r, nil" "r" }}
	{{ else -}}
		k := req.GetKey()
		if t, ok := k.(*{{ print $.EntMsg.Ident "_Id" | $.Pb.Ident | use }}); ok {
			{{ to_ent_with_rv $id_field "t.Id" "v" "return @, nil" "r" }}
		}

		p, err := Get{{ $.EntMsg.Schema.Name }}Specifier(req)
		if err != nil {
			return r, err
		}

		v, err := db.{{ $.EntMsg.Schema.Name }}.Query().Where(p).OnlyID(ctx)
		if err != nil {
			return r, ToStatus(err)
		}

		return v, nil
	{{ end -}}
}
