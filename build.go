package entpb

import (
	"cmp"
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
	"github.com/lesomnus/entpb/utils"
)

type Build struct {
	// Filepath or alias to proto file to be output.
	Files map[string]*ProtoFile

	// Key is global name of the Go type that bound to enum.
	// e.g. for enum "Role" that bound to Go type "Role" in package "github.com/lesomnus/entpb/example",
	// the key would be its global name, "github.com/lesomnus/entpb/example:example.Role".
	// Global name can be built using following functions:
	//   - utils.FullIdent
	//   - utils.FullIdentFromReflect
	//   - utils.FullIdentFromEntTypeInfo
	Enums map[string]*Enum

	// Holds message definitions.
	// Key is name of message.
	// e.g. User, GetUserRequest, Identity, ...
	Messages map[ident.Ident]*MessageAnnotation

	// Messages defined by Ent schema.
	Schemas map[string]*MessageAnnotation
}

func NewBuild(graph *gen.Graph) (*Build, error) {
	b := &Build{
		Files:    map[string]*ProtoFile{},
		Enums:    map[string]*Enum{},
		Messages: map[ident.Ident]*MessageAnnotation{},
		Schemas:  map[string]*MessageAnnotation{},
	}
	if err := b.parse(graph); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Build) parse(graph *gen.Graph) error {
	if d, ok := DecodeAnnotation(&ProtoFiles{}, graph.Annotations); !ok {
		return nil
	} else {
		for p, f := range *d {
			for name, enum := range f.Enums {
				if _, ok := b.Enums[name]; ok {
					return fmt.Errorf(`multiple definition of enum for same Go type "%s"`, name)
				}

				enum.File = f
				b.Enums[name] = enum
			}

			f.Filepath = p
			b.Files[p] = f
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
	for _, enum := range b.Enums {
		if err := b.normalizeEnum(enum); err != nil {
			errs = append(errs, fmt.Errorf(`normalize enum "%s": %w`, enum.Ident, err))
		}
	}
	for _, msg := range b.Schemas {
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
	d, ok := DecodeAnnotation(&MessageAnnotation{}, r.Annotations)
	if !ok {
		return nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}
	if a, ok := DecodeAnnotation(&schema.CommentAnnotation{}, r.Annotations); ok {
		d.Comment = a.Text
	}

	f, ok := b.Files[d.Filepath]
	if !ok {
		return fmt.Errorf(`message "%s" references non-exists proto file "%s"`, d.Ident, d.Filepath)
	}

	if _, ok := f.Messages[d.Ident]; ok {
		return fmt.Errorf(`message name "%s" duplicated with proto file "%s"`, d.Ident, d.Filepath)
	}

	d.File = f
	d.Schema = r
	f.Messages[d.Ident] = d
	b.Messages[d.Ident] = d
	b.Schemas[r.Name] = d
	return nil
}

func (p *Build) normalizeEnum(enum *Enum) error {
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

func (b *Build) parseFields(m *MessageAnnotation) error {
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

func (p *Build) parseEntField(r *load.Field) (*FieldAnnotation, error) {
	d, ok := DecodeAnnotation(&FieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}

	if r.Info.Type == field.TypeEnum {
		name := utils.FullIdentFromEntTypeInfo(r.Info)
		enum, ok := p.Enums[name]
		if !ok {
			return nil, fmt.Errorf(`unregistered enum type: "%s"`, name)
		}

		d.PbType = PbType{
			Ident:   enum.Ident,
			Package: enum.File.PbPackage,
			Import:  enum.File.Filepath,
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
	d.IsKey = r.Unique
	d.IsImmutable = r.Immutable
	d.IsOptional = r.Nillable

	return d, nil
}

func (p *Build) parseEntEdge(r *load.Edge) (*FieldAnnotation, error) {
	d, ok := DecodeAnnotation(&FieldAnnotation{}, r.Annotations)
	if !ok {
		return nil, nil
	}
	if d.Ident == "" {
		d.Ident = ident.Ident(r.Name)
	}

	message, ok := p.Schemas[r.Type]
	if !ok {
		return nil, fmt.Errorf(`references a schema "%s" that is not a proto message`, r.Type)
	}

	d.Comment = r.Comment
	d.EntName = r.Name
	d.EntRef = r.Type
	d.PbType = message.pbType()
	d.IsOptional = !r.Required
	d.IsRepeated = !r.Unique
	d.IsImmutable = r.Immutable

	return d, nil
}

func (p *Build) parseService(d *MessageAnnotation) error {
	s := d.Service
	if s == nil || len(s.Rpcs) == 0 {
		return nil
	}

	s.Message = d
	if s.Filepath == "" {
		s.Filepath = d.Filepath
	}

	f, ok := p.Files[s.Filepath]
	if !ok {
		return fmt.Errorf(`service "%s" references non-exists proto file "%s"`, d.Ident, s.Filepath)
	} else {
		s.File = f
	}
	if s.Ident == "" {
		s.Ident = ident.Ident(fmt.Sprintf("%sService", d.Ident))
	}
	if _, ok := f.Services[s.Ident]; ok {
		return fmt.Errorf(`duplicated service "%s"`, s.Ident)
	} else {
		f.Services[s.Ident] = s
	}

	for _, rpc := range s.Rpcs {
		switch rpc.Ident {
		case "Create":
			rpc.Req = d.pbType()
			rpc.Res = d.pbType()
			rpc.EntReq = d
			rpc.EntRes = d

		case "Get":
			req_name := ident.Ident(fmt.Sprintf("Get%sRequest", d.Ident))
			rpc.Req = d.pbType()
			rpc.Req.Ident = req_name
			rpc.Res = d.pbType()

			msg := &MessageAnnotation{
				Filepath: s.Filepath,
				Ident:    req_name,
				File:     s.File,
			}
			key_fields := []*FieldAnnotation{}
			for _, field := range d.Fields {
				if !field.IsKey {
					continue
				}

				key_fields = append(key_fields, field)
			}
			if len(key_fields) == 1 {
				msg.Fields = append(msg.Fields, key_fields[0])
			} else {
				field := &FieldAnnotation{Ident: "key"}
				field.Oneof = append(field.Oneof, key_fields...)
				field.Number = slices.MinFunc(key_fields, func(a, b *FieldAnnotation) int {
					return cmp.Compare(a.Number, b.Number)
				}).Number
				msg.Fields = append(msg.Fields, field)
			}
			// TODO: add index as a key? How?
			// To make an index as a oneof field, new message need to be created.

			s.File.Messages[req_name] = msg
			p.Messages[req_name] = msg

			rpc.EntReq = msg
			rpc.EntRes = d

		case "Update":
			req_name := ident.Ident(fmt.Sprintf("Update%sRequest", d.Ident))
			rpc.Req = d.pbType()
			rpc.Req.Ident = req_name
			rpc.Res = d.pbType()

			msg := &MessageAnnotation{
				Filepath: s.Filepath,
				Ident:    req_name,
				File:     s.File,
			}
			for _, field := range d.Fields {
				if field.EntName == "id" {
					msg.Fields = append(msg.Fields, field)
					continue
				}
				if field.IsReadOnly {
					continue
				}

				v := *field
				v.IsOptional = true
				msg.Fields = append(msg.Fields, &v)
			}

			s.File.Messages[req_name] = msg
			p.Messages[req_name] = msg

			rpc.EntReq = msg
			rpc.EntRes = d

		case "Delete":
			req_name := ident.Ident(fmt.Sprintf("Delete%sRequest", d.Ident))
			rpc.Req = d.pbType()
			rpc.Req.Ident = req_name
			rpc.Res = PbEmpty

			msg := &MessageAnnotation{
				Filepath: s.Filepath,
				Ident:    req_name,
				File:     s.File,
			}
			key_fields := []*FieldAnnotation{}
			for _, field := range d.Fields {
				if !field.IsKey {
					continue
				}

				key_fields = append(key_fields, field)
			}
			if len(key_fields) == 1 {
				msg.Fields = append(msg.Fields, key_fields[0])
			} else {
				field := &FieldAnnotation{Ident: "key"}
				field.Oneof = append(field.Oneof, key_fields...)
				field.Number = slices.MinFunc(key_fields, func(a, b *FieldAnnotation) int {
					return cmp.Compare(a.Number, b.Number)
				}).Number
				msg.Fields = append(msg.Fields, field)
			}
			// TODO: add index as a key? How?
			// To make an index as a oneof field, new message need to be created.

			s.File.Messages[req_name] = msg
			p.Messages[req_name] = msg

			rpc.EntReq = msg
			rpc.EntRes = d

		default:
			if rpc.Req.Equal(&PbThis) {
				rpc.Req = d.pbType()
			}
			if rpc.Res.Equal(&PbThis) {
				rpc.Res = d.pbType()
			}

			if rpc.Req.Import == "" {
				return fmt.Errorf(`RPC "%s": parameter type must be message`, rpc.Ident)
			}
			if rpc.Res.Import == "" {
				return fmt.Errorf(`RPC "%s": return type must be message`, rpc.Ident)
			}
		}
	}

	return nil
}
