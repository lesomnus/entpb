{{ $serverName := (print .GoName "Server")  -}}
type {{ $serverName }} struct {
	db *{{ ident .Ent "Client" }}
	{{ ident $.Pb (print "Unimplemented" $serverName ) }}
}
