{{ $server_name := (print $.PbSvc.GoName "Server") -}}
func New{{ $server_name }}(db *{{ $.Ent.Ident "Client" | use }}) *{{ $server_name }} {
	return &{{ $server_name }}{db: db}
}
