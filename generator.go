package entpb

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen"
	"github.com/mitchellh/mapstructure"
)

type generator struct {
	out_dir string // Expected it to be a absolute path.

	files map[string]*ProtoFile
	enums map[string]*ProtoFile // Holds proto file that contains the enum definition.
}

func (g *generator) generate(graph *gen.Graph) error {
	d, ok := decodeAnnotation(&ProtoFiles{}, graph.Annotations)
	if !ok {
		return nil
	}

	if err := g.parseGraph(graph); err != nil {
		return fmt.Errorf("parse graph: %w", err)
	}

	// files := []pbgen.ProtoFile{}
	for p, f := range *d {
		proto_file := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Package: pbgen.ParseFullIndent(f.pbPackage),
		}
		if f.goPackage != "" {
			proto_file.Options = append(proto_file.Options, pbgen.OptionGoPackage(f.goPackage))
		}
		for _, enum := range f.enums {
			d := pbgen.Enum{Name: enum.t.Name()}
			for k, v := range enum.vs {
				d.Body = append(d.Body, pbgen.EnumField{Name: k, Number: v})
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		os_path := filepath.Join(g.out_dir, p)
		w, err := os.Create(os_path)
		if err != nil {
			return fmt.Errorf(`create file at "%s": %w`, os_path, err)
		}
		if err := pbgen.Execute(w, proto_file); err != nil {
			return fmt.Errorf(`generate proto file for "%s": %w`, p, err)
		}
	}

	return nil
}

func (g *generator) parseGraph(graph *gen.Graph) error {
	var d map[string]*ProtoFile
	if a, ok := graph.Annotations[ProtoFilesAnnotation]; !ok {
		return nil
	} else if err := mapstructure.Decode(a, &d); err != nil {
		panic(fmt.Errorf("decode %s: %w", ProtoFilesAnnotation, err))
	}

	for p, f := range d {
		if l, ok := d[f.alias]; ok {
			if f == l {
				// Note that the map is mutated while it is iterated,
				// so newly added element can be visited.
				continue
			}
			return fmt.Errorf(`duplicated alias "%s" for "%s"`, f.alias, p)
		}

		for name := range f.enums {
			if _, ok := g.enums[name]; ok {
				return fmt.Errorf(`multiple definition of enum for same Go type "%s"`, name)
			}

			g.enums[name] = f
		}

		f.path = p
		g.files[p] = f
		g.files[f.alias] = f
	}

	errs := []error{}
	for _, s := range graph.Schemas {
		// Note that `parseMessage` does not parse their fields but only its name
		// since there may be dependencies between messages.
		if err := g.parseMessage(s); err != nil {
			errs = append(errs, fmt.Errorf(`schema "%s": %w`, s.Name, err))
			continue
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse messages: %w", errors.Join(errs...))
	}

	errs = []error{}
	for _, file := range g.files {
		for _, msg := range file.messages {
			if err := g.parseFields(msg); err != nil {
				errs = append(errs, fmt.Errorf(`schema "%s": %w`, msg.ref.Name, err))
				continue
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse fields: %w", errors.Join(errs...))
	}

	return nil
}

func (g *generator) parseMessage(r *load.Schema) error {
	d, ok := decodeAnnotation(&messageAnnotation{}, r.Annotations)
	if !ok {
		return nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = a.Value
	} else {
		d.name = r.Name
	}
	if a, ok := decodeAnnotation(&schema.CommentAnnotation{}, r.Annotations); ok {
		d.comment = a.Text
	}

	f, ok := g.files[d.Filepath]
	if !ok {
		return fmt.Errorf(`message "%s" references non-exists proto file "%s"`, d.name, d.Filepath)
	}

	if _, ok := f.messages[d.name]; ok {
		return fmt.Errorf(`message name "%s" duplicated with proto file "%s"`, d.name, d.Filepath)
	}

	d.ref = r
	f.messages[d.name] = d
	return nil
}

func (g *generator) parseFields(m *messageAnnotation) error {
	errs := []error{}
	for _, field := range m.ref.Fields {
		d, err := g.parseEntField(field)
		if err != nil {
			errs = append(errs, fmt.Errorf(`field "%s": %w`, field.Name, err))
			continue
		}

		m.fields = append(m.fields, d)
	}
	for _, edge := range m.ref.Edges {
		d, err := g.parseEntEdge(edge)
		if err != nil {
			errs = append(errs, fmt.Errorf(`edge "%s": %w`, edge.Name, err))
			continue
		}

		m.fields = append(m.fields, d)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (g *generator) parseEntField(r *load.Field) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = a.Value
	} else {
		d.name = r.Name
	}

	if r.Info.Type == field.TypeEnum {
		name := globalTypeNameFromEntTypeInfo(r.Info)
		f, ok := g.enums[name]
		if !ok {
			return nil, fmt.Errorf("unregistered enum type: %s", name)
		}

		enum, ok := f.enums[name]
		if !ok {
			panic(fmt.Errorf(`invalid state of generator: enum does not exist on file: "%s"`, name))
		}

		d.pb_type = PbType{
			Name:   enum.t.Name(),
			Import: f.path,
		}
	} else if t := pb_types[int(r.Info.Type)]; t.Name == "" {
		return nil, fmt.Errorf("unsupported type: %s", r.Info.Type.String())
	} else {
		d.pb_type = t
	}

	d.comment = r.Comment
	d.isOptional = r.Nillable

	return d, nil
}

func (g *generator) parseEntEdge(r *load.Edge) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if a, ok := decodeAnnotation(&nameAnnotation{}, r.Annotations); ok {
		d.name = a.Value
	} else {
		d.name = r.Name
	}

	// TODO: resolve pb_type

	d.comment = r.Comment
	d.isOptional = !r.Required
	d.isRepeated = !r.Unique

	return nil, nil
}

// func (g *generator) generateMessage(s *load.Schema) error {
// 	var d MessageDescriptor
// 	if a, ok := s.Annotations[MessageAnnotation]; !ok {
// 		return nil
// 	} else if err := mapstructure.Decode(a, &d); err != nil {
// 		panic(fmt.Errorf("decode proto message: %w", err))
// 	}

// 	if d.PbFilepath == "" {
// 		return fmt.Errorf(`path to proto file not provided`)
// 	}

// 	d.name = s.Name
// 	if a, ok := s.Annotations["Comment"]; ok {
// 		var c schema.CommentAnnotation
// 		if err := mapstructure.Decode(a, &c); err != nil {
// 			panic(fmt.Errorf("decode comment annotation: %w", err))
// 		}
// 		d.comment = c.Text
// 	}

// 	o, err := g.getMessageOutput(&d)
// 	if err != nil {
// 		return fmt.Errorf("get output: %w", err)
// 	}
// 	if o.GoPackage == "" {
// 		o.GoPackage = d.GoPackage
// 	}

// 	for _, field := range s.Fields {
// 		v, err := g.parseField(field)
// 		if err != nil {
// 			return fmt.Errorf("generate proto field for Ent field %s: %w", field.Name, err)
// 		}
// 		if v == nil {
// 			continue
// 		}

// 		d.fields = append(d.fields, v)
// 		if v.Type.Import != "" {
// 			o.PackageDependencies[v.Type.Import] = struct{}{}
// 		}
// 	}

// 	edges := s.Edges[:]
// 	for _, edge := range s.Edges {
// 		if edge.Ref == nil {
// 			continue
// 		}
// 		if edge.Ref.Type == edge.Type {
// 			edges = append(edges, edge.Ref)
// 		}
// 	}
// 	for _, edge := range edges {
// 		v, err := g.parseEdge(edge)
// 		if err != nil {
// 			return fmt.Errorf("generate proto field for Ent edge %s: %w", edge.Name, err)
// 		}
// 		if v == nil {
// 			continue
// 		}

// 		d.fields = append(d.fields, v)
// 	}
// 	slices.SortFunc(d.fields, func(l, r *FieldDescriptor) int {
// 		return cmp.Compare(l.Number, r.Number)
// 	})

// 	if err := tpl.ExecuteTemplate(o, "message-def.tpl", d); err != nil {
// 		return fmt.Errorf("execute template for message: %w", err)
// 	}

// 	return nil
// }

// func (g *generator) parseField(f *load.Field) (*FieldDescriptor, error) {
// 	d := &FieldDescriptor{}
// 	if a, ok := f.Annotations[FieldAnnotation]; !ok {
// 		return nil, nil
// 	} else if err := mapstructure.Decode(a, d); err != nil {
// 		panic(fmt.Errorf("decode proto field: %w", err))
// 	}

// 	if t := pb_types[int(f.Info.Type)]; t.Name == "" {
// 		return nil, fmt.Errorf("unsupported type: %s", f.Info.Type.String())
// 	} else {
// 		d.Type = t
// 	}

// 	d.Name = f.Name
// 	d.comment = f.Comment
// 	d.isOptional = f.Nillable

// 	return d, nil
// }

// func (g *generator) parseEdge(e *load.Edge) (*FieldDescriptor, error) {
// 	d := &FieldDescriptor{}
// 	if a, ok := e.Annotations[FieldAnnotation]; !ok {
// 		return nil, nil
// 	} else if err := mapstructure.Decode(a, d); err != nil {
// 		panic(fmt.Errorf("decode proto field: %w", err))
// 	}

// 	d.Name = e.Name
// 	d.Type = PbType{Name: e.Type}
// 	d.comment = e.Comment
// 	d.isOptional = !e.Required
// 	d.isRepeated = !e.Unique

// 	return d, nil
// }

// func (g *generator) getMessageOutput(d *MessageDescriptor) (*MessageOut, error) {
// 	p := d.PbFilepath
// 	v, ok := g.msg_out[p]
// 	if !ok {
// 		o, err := g.open(p)
// 		if err != nil {
// 			return nil, fmt.Errorf("open for proto output: %w", err)
// 		}

// 		v, err = NewMessageOutput(o, d)
// 		if err != nil {
// 			return nil, fmt.Errorf("create printer: %w", err)
// 		}

// 		g.msg_out[p] = v
// 	}

// 	return v, nil
// }
