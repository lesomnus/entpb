{{ $serverName := (print $.GoName "Server")  -}}
type {{ $serverName }} struct {
	db *{{ $.Ent.Ident "Client" | use }}
	{{ print "Unimplemented" $serverName | $.Pb.Ident | use }}
}
