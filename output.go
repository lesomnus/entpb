package entpb

// var (
// 	//go:embed template/*
// 	templates embed.FS

// 	tpl *template.Template
// )

// type MessageOut struct {
// 	io.WriteCloser

// 	PackageDependencies map[string]struct{}

// 	GoPackage string
// }

// func NewMessageOutput(o io.WriteCloser, d *MessageDescriptor) (*MessageOut, error) {
// 	p := &MessageOut{
// 		WriteCloser:         o,
// 		PackageDependencies: map[string]struct{}{},
// 	}

// 	if err := tpl.ExecuteTemplate(o, "message-head.tpl", d); err != nil {
// 		return nil, err
// 	}

// 	return p, nil
// }

// func init() {
// 	t, err := template.ParseFS(templates, "template/*")
// 	if err != nil {
// 		panic(err)
// 	}

// 	tpl = t
// }
