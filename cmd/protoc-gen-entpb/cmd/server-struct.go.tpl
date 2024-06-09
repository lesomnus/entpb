{{ $server_name := (print $.PbSvc.GoName "Server") -}}
type {{ $server_name }} struct {
	db *{{ $.Ent.Ident "Client" | use }}
	{{ print "Unimplemented" $server_name | $.Pb.Ident | use }}
}
