{{ $fields := $.EntMsg.Fields -}}
{{ $isOneof := (index $fields 0).IsOneof -}}
{{ if $isOneof -}}
{{ $fields = (index $fields 0).Oneof -}}
{{ end -}}
func Get{{ $.EntMsg.Schema.Name }}Specifier(req *{{ print $.EntMsg.Ident | $.Pb.Ident | use }}) ({{ $.Pred.Ident $.EntMsg.Schema.Name | use }}, error) {
	{{ $schema := schema $.EntMsg.Schema -}}
	{{ if not $isOneof -}}
		{{ to_ent (index $fields 0) "req.GetId()" "v" (print "return " ($schema.Ident "IDEQ" | use) "(@), nil") }}
	{{ else -}}
		switch t := req.GetKey().(type) {
		{{ range $fields -}}
			{{ $key_name := print .Ident | pascal -}}
			{{ $in := print "t." $key_name -}}
			{{ $pred := print (.EntName | entname) "EQ" | (schema $.EntMsg.Schema).Ident | use  -}}
			case *{{ print $.EntMsg.Ident "_" $key_name | $.Pb.Ident | use }}:
				{{ print "return " $pred "(@), nil" | to_ent . $in "v" }}
		{{ end -}}
		case nil:
			return nil, {{ status "InvalidArgument" "key not provided" }}
		default:
			return nil, {{ status "Unimplemented" "unknown type of key" }}
		}
	{{ end -}}
}
