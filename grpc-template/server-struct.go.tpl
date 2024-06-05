{{ $serverName := (print .GoName "Server")  -}}
type {{ $serverName }} struct {
	db *{{ ident .Ent "Client" }}
	{{ ident .Grpc (print "Unimplemented" $serverName ) }}
}
