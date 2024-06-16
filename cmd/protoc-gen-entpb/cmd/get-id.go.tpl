{{ $fields := $.EntMsg.Fields -}}
{{ $isOneof := (index $fields 0).IsOneof -}}
{{ if $isOneof -}}
{{ $fields = (index $fields 0).Oneof -}}
{{ end -}}
{{ $id_type := ent_type (index $fields 0) -}}
func Get{{ $.EntMsg.Schema.Name }}Id(ctx {{ import "context" | ident "Context" }}, db *{{ $.Ent.Ident "Client" | use }}, req *{{ print $.EntMsg.Ident | $.Pb.Ident | use }}) ({{ $id_type }}, error) {
	var r {{ $id_type }}
	{{ if not $isOneof -}}
		{{ to_ent_with_rv (index $fields 0) "req.GetId()" "v" "return @, nil" "r" }}
	{{ else -}}
		k := req.GetKey()
		if t, ok := k.(*{{ print $.EntMsg.Ident "_Id" | $.Pb.Ident | use }}); ok {
			{{ to_ent_with_rv (index $fields 0) "t.Id" "v" "return @, nil" "r" }}
		}

		q := db.{{ $.EntMsg.Schema.Name }}.Query()
		switch t := k.(type) {
		{{ range (slice $fields 1) -}}
			{{ $key_name := print .Ident | pascal -}}
			{{ $in := print "t." $key_name -}}
			{{ $pred := print (.EntName | entname) "EQ" | (schema $.EntMsg.Schema).Ident | use  -}}
			case *{{ print $.EntMsg.Ident "_" $key_name | $.Pb.Ident | use }}:
				{{ print "q.Where(" $pred "(@))" | to_ent . $in "v" }}
		{{ end -}}
		case nil:
			return r, {{ status "InvalidArgument" "key not provided" }}
		default:
			return r, {{ status "Unimplemented" "unknown type of key" }}
		}
		if v, err := q.OnlyID(ctx); err != nil {
			return r, ToStatus(err)
		} else {
			return v, nil
		}
	{{ end -}}
}
