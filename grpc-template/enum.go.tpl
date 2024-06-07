func toPb{{ $.PbEnum.Desc.Name }}(v {{ import $.EntEnum.GoType.PkgPath | ident $.EntEnum.GoType.Name }}) {{ $.PbEnum.GoIdent | use }} {
	switch v {
	{{ range $.EntEnum.Fields -}}
	case "{{ .Value }}" :
		return {{ .Number }}
	{{ end -}}
	default:
		return 0
	}
}

func toEnt{{ $.EntEnum.GoType.Name }}(v {{ $.PbEnum.GoIdent | use }}) {{ import $.EntEnum.GoType.PkgPath | ident $.EntEnum.GoType.Name }} {
	switch v {
	{{ range $.EntEnum.Fields -}}
	case {{ .Number }} :
		return "{{ .Value }}"
	{{ end -}}
	default:
		return ""
	}
}
