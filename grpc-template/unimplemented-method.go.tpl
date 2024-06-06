{{ $serverName := (print .Service.GoName "Server")  -}}
func (s *{{ $serverName }}) {{ .GoName }}({{ ident "context" "Context" }}, *{{ goident .Input.GoIdent }}) (*{{ goident .Output.GoIdent }}, error) {
	return nil, {{ ident .Status "Errorf" }}({{ ident .Codes "Unimplemented" }}, "method {{ .GoName }} not implemented")
}
