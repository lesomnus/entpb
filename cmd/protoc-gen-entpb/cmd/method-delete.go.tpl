func (s *{{ print $.PbSvc.GoName "Server" }}) {{ $.PbMethod.GoName }}(ctx {{ import "context" | ident "Context" }}, req *{{ $.PbMethod.Input.GoIdent | use }}) (*{{ $.PbMethod.Output.GoIdent | use }}, error) {
	p, err := Get{{ $.EntMsg.Schema.Name }}Specifier(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.{{ $.EntMsg.Schema.Name }}.Delete().Where(p).Exec(ctx); err != nil {
		return nil, ToStatus(err)
	}

	 return &{{ $.PbMethod.Output.GoIdent | use }}{}, nil
}
