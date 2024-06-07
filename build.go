package entpb

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/iancoleman/strcase"
	"github.com/lesomnus/entpb/pbgen/ident"
)

type Build struct {
	// Filepath or alias to proto file to be output.
	files map[string]*ProtoFile

	// Key is global name of the Go type that bound to enum.
	// e.g. for enum "Role" that bound to Go type "Role" in package "github.com/lesomnus/entpb/example",
	// the key would be its global name, "github.com/lesomnus/entpb/example:example.Role".
	// Global name can be built using following functions:
	//   - globalTypeName
	//   - globalTypeNameFromReflect
	//   - globalTypeNameFromEntTypeInfo
	enums map[string]*enum

	// Holds message definitions.
	// Key is name of Ent Schema.
	// e.g. User, Identity, ...
	messages map[string]*messageAnnotation
}

func NewBuild(graph *gen.Graph) (*Build, error) {
	b := &Build{
		files:    map[string]*ProtoFile{},
		enums:    map[string]*enum{},
		messages: map[string]*messageAnnotation{},
	}
	if err := b.parse(graph); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Build) parse(graph *gen.Graph) error {
	if d, ok := decodeAnnotation(&ProtoFiles{}, graph.Annotations); !ok {
		return nil
	} else {
		for p, f := range *d {
			for name, enum := range f.enums {
				if _, ok := b.enums[name]; ok {
					return fmt.Errorf(`multiple definition of enum for same Go type "%s"`, name)
				}

				enum.File = f
				b.enums[name] = enum
			}

			f.path = p
			b.files[p] = f
		}
	}

	errs := []error{}
	for _, s := range graph.Schemas {
		// Note that `parseMessage` does not parse their fields but only its name
		// since there may be dependencies between messages.
		if err := b.parseMessage(s); err != nil {
			errs = append(errs, fmt.Errorf(`schema "%s": %w`, s.Name, err))
			continue
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse messages declarations: %w", errors.Join(errs...))
	}

	errs = []error{}
	for _, enum := range b.enums {
		if err := b.normalizeEnum(enum); err != nil {
			errs = append(errs, fmt.Errorf(`normalize enum "%s": %w`, enum.Ident, err))
		}
	}
	for _, msg := range b.messages {
		errs_ := []error{}
		if err := b.parseFields(msg); err != nil {
			errs_ = append(errs_, fmt.Errorf(`parse fields: %w`, err))
		}
		if err := b.parseService(msg); err != nil {
			errs_ = append(errs_, fmt.Errorf(`parse service: %w`, err))
		}
		if len(errs_) > 0 {
			errs = append(errs, fmt.Errorf(`schema "%s": %w`, msg.Schema.Name, errors.Join(errs_...)))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("parse message definitions: %w", errors.Join(errs...))
	}

	return nil
}

func (b *Build) parseMessage(r *load.Schema) error {
	d, ok := decodeAnnotation(&messageAnnotation{}, r.Annotations)
	if !ok {
		return nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}
	if a, ok := decodeAnnotation(&schema.CommentAnnotation{}, r.Annotations); ok {
		d.Comment = a.Text
	}

	f, ok := b.files[d.Filepath]
	if !ok {
		return fmt.Errorf(`message "%s" references non-exists proto file "%s"`, d.Ident, d.Filepath)
	}

	if _, ok := f.messages[d.Ident]; ok {
		return fmt.Errorf(`message name "%s" duplicated with proto file "%s"`, d.Ident, d.Filepath)
	}

	d.File = f
	d.Schema = r
	f.messages[d.Ident] = d
	b.messages[r.Name] = d
	return nil
}

func (p *Build) normalizeEnum(enum *enum) error {
	prefix := ""
	has_zero := false
	if enum.Prefix == nil {
		// no prefix
	} else if *enum.Prefix == "" {
		prefix = fmt.Sprintf("%s_", enum.Ident)
	} else {
		prefix = fmt.Sprintf("%s_", *enum.Prefix)
	}

	for _, v := range enum.Fields {
		if v.Number == 0 {
			has_zero = true
			break
		}
	}
	if !enum.IsClosed && !has_zero {
		enum.Fields = append(enum.Fields, &EnumField{
			Number: 0,
			Value:  "Unspecified",
		})
	}
	for _, v := range enum.Fields {
		name := fmt.Sprintf("%s%s", prefix, v.Value)
		name = strcase.ToSnake(name)
		name = strings.ToUpper(name)
		v.Name = name
	}

	return nil
}

func (b *Build) parseFields(m *messageAnnotation) error {
	errs := []error{}
	for _, field := range m.Schema.Fields {
		d, err := b.parseEntField(field)
		if err != nil {
			errs = append(errs, fmt.Errorf(`field "%s": %w`, field.Name, err))
			continue
		}
		if d == nil {
			continue
		}

		m.Fields = append(m.Fields, d)
	}

	edges := slices.Clone(m.Schema.Edges)
	for _, edge := range m.Schema.Edges {
		if edge.Ref == nil {
			continue
		}
		if edge.Ref.Type == edge.Type {
			edges = append(edges, edge.Ref)
		}
	}
	for _, edge := range edges {
		d, err := b.parseEntEdge(edge)
		if err != nil {
			errs = append(errs, fmt.Errorf(`edge "%s": %w`, edge.Name, err))
			continue
		}
		if d == nil {
			continue
		}

		m.Fields = append(m.Fields, d)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (p *Build) parseEntField(r *load.Field) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}

	if r.Info.Type == field.TypeEnum {
		name := globalTypeNameFromEntTypeInfo(r.Info)
		enum, ok := p.enums[name]
		if !ok {
			return nil, fmt.Errorf(`unregistered enum type: "%s"`, name)
		}

		d.PbType = PbType{
			Ident:   enum.Ident,
			Package: enum.File.pbPackage,
			Import:  enum.File.path,
		}
	} else if t := pb_types[int(r.Info.Type)]; t.Ident == "" {
		return nil, fmt.Errorf("unsupported type: %s", r.Info.Type.String())
	} else {
		d.PbType = t
	}

	d.Comment = r.Comment
	d.EntName = r.Name
	d.EntInfo = r.Info
	d.HasDefault = r.Default
	d.IsOptional = r.Nillable

	return d, nil
}

func (p *Build) parseEntEdge(r *load.Edge) (*fieldAnnotation, error) {
	d, ok := decodeAnnotation(&fieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}

	message, ok := p.messages[r.Type]
	if !ok {
		return nil, fmt.Errorf(`edge "%s" references a schema "%s" that is not a proto message`, r.Name, r.Type)
	}

	d.Comment = r.Comment
	d.EntName = r.Name
	d.EntRef = r.Type
	d.PbType = message.pbType()
	d.IsOptional = !r.Required
	d.IsRepeated = !r.Unique

	return d, nil
}

func (p *Build) parseService(d *messageAnnotation) error {
	s := d.Service
	if s == nil || len(s.Rpcs) == 0 {
		return nil
	}

	s.message = d
	if s.Filepath == "" {
		s.Filepath = d.Filepath
	}

	f, ok := p.files[s.Filepath]
	if !ok {
		return fmt.Errorf(`service "%s" references non-exists proto file "%s"`, d.Ident, s.Filepath)
	}
	if s.Name == "" {
		s.Name = ident.Ident(fmt.Sprintf("%sService", d.Ident))
	}
	if _, ok := f.services[s.Name]; ok {
		return fmt.Errorf(`duplicated service "%s"`, s.Name)
	} else {
		f.services[s.Name] = s
	}

	for _, rpc := range s.Rpcs {
		if rpc.Req.Equal(&PbThis) {
			rpc.Req = d.pbType()
		}
		if rpc.Res.Equal(&PbThis) {
			rpc.Res = d.pbType()
		}

		if rpc.Req.Import == "" {
			return fmt.Errorf(`RPC "%s": parameter type must be message`, rpc.Name)
		}
		if rpc.Res.Import == "" {
			return fmt.Errorf(`RPC "%s": return type must be message`, rpc.Name)
		}
	}

	return nil
}
