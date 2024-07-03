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
			{{ $key_ident := .Ident -}}
			{{ $key_name := print .Ident | pascal -}}
			{{ $in := print "t." $key_name -}}
			case *{{ print $.EntMsg.Ident "_" $key_name | $.Pb.Ident | use }}:
			{{ if not .IsTypeMessage -}}
				{{ $pred := print (.EntName | entname) "EQ" | $schema.Ident | use  -}}
				{{ print "return " $pred "(@), nil" | to_ent . $in "v" }}
			{{ else -}}
				{{ $fields := .EntMsg.Fields -}}
				ps := make([]{{ $.Pred.Ident $.EntMsg.Schema.Name | use }}, 0, {{ len $fields }})
				{{ range $fields -}}
					{{ $in := print $in ".Get" (print .Ident | pascal) "()" -}}
					{{ if not .IsTypeMessage -}}
						{{ $pred := print (.EntName | entname) "EQ" | $schema.Ident | use  -}}
						{{ print "ps = append(ps, " $pred "(@))" | to_ent . $in "v" }}
					{{ else -}}
						if p, err := Get{{ .EntMsg.Schema.Name }}Specifier({{ $in }}); err != nil {
							s, _ := status.FromError(err)
							return nil, {{ status_errf "InvalidArgument" (print $key_ident ".%s") "s.Message()" }}
						} else {
							ps = append(ps, {{ print .Ident | pascal | printf "Has%sWith" | $schema.Ident | use }}(p))
						}
					{{ end -}}
				{{ end -}}
				return {{ $schema.Ident "And" | use }}(ps...), nil
			{{ end -}}
		{{ end -}}
		case nil:
			return nil, {{ status "InvalidArgument" "key not provided" }}
		default:
			return nil, {{ status "Unimplemented" "unknown type of key" }}
		}
	{{ end -}}
}
