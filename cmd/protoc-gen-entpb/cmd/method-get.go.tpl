func (s *{{ print $.Service.GoName "Server" }}) {{ $.Method.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.Method.Input.GoIdent | use }}) (*{{ $.Method.Output.GoIdent | use }}, error) {

}
