package entpb

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/lesomnus/entpb/pbgen"
	"golang.org/x/exp/maps"
)

type Fs interface {
	Create(name string) (io.WriteCloser, error)
	MkdirAll(path string, perm os.FileMode) error
}

type Printer interface {
	// Prints out build result into `fs`.
	// `fs` must be subtree by desired output root directory.
	Print(b *Build) error
}

func NewProtoPrinter(fs Fs, edition pbgen.Edition) (Printer, error) {
	switch edition {
	case pbgen.Edition2023:
		return &edition2023Printer{printerUtils{fs}}, nil
	case pbgen.SyntaxProto3:
		return &proto3Printer{printerUtils{fs}}, nil

	case pbgen.SyntaxProto2:
		fallthrough
	default:
		return nil, errors.New("not implemented")
	}
}

type edition2023Printer struct {
	printerUtils
}

func (p *edition2023Printer) Print(b *Build) error {
	return p.print(b, p.printFile)
}

func (p *edition2023Printer) printFile(f *ProtoFile) pbgen.ProtoFile {
	v := pbgen.ProtoFile{
		Edition: pbgen.Edition2023,
		Package: f.PbPackage,
		Imports: p.importPaths(f),
		Options: []pbgen.Option{pbgen.FeatureFieldPresenceImplicit},
	}
	if f.GoPackage != "" {
		v.Options = append(v.Options, pbgen.OptionGoPackage(f.GoPackage))
	}

	a := func(ds ...pbgen.TopLevelDef) {
		v.TopLevelDefinitions = append(v.TopLevelDefinitions, ds...)
	}

	a(p.printServices(f)...)
	a(p.printEnums(f, func(enum *Enum) pbgen.Enum {
		d := pbgen.Enum{}
		if enum.IsClosed {
			d.Options = append(d.Options, pbgen.FeatureEnumTypeClosed)
		}

		return d
	})...)
	a(p.printMessages(f, func(a *FieldAnnotation) pbgen.MessageField {
		d := pbgen.MessageField{
			Type:   a.PbType.ReferencedIn(f.PbPackage),
			Name:   a.Ident,
			Number: a.Number,
		}
		if a.IsRepeated {
			d.Labels = append(d.Labels, pbgen.LabelRepeated)
		} else if a.IsOptional {
			// Presence of "repeated" fields are not tracked.
			d.Options = append(d.Options, pbgen.FeatureFieldPresenceExplicit)
		}

		return d
	})...)

	return v
}

type proto3Printer struct {
	printerUtils
}

func (p *proto3Printer) Print(b *Build) error {
	return p.print(b, p.printFile)
}

func (p *proto3Printer) printFile(f *ProtoFile) pbgen.ProtoFile {
	v := pbgen.ProtoFile{
		Edition: pbgen.SyntaxProto3,
		Package: f.PbPackage,
		Imports: p.importPaths(f),
	}
	if f.GoPackage != "" {
		v.Options = append(v.Options, pbgen.OptionGoPackage(f.GoPackage))
	}

	a := func(ds ...pbgen.TopLevelDef) {
		v.TopLevelDefinitions = append(v.TopLevelDefinitions, ds...)
	}

	a(p.printServices(f)...)
	a(p.printEnums(f, func(enum *Enum) pbgen.Enum {
		return pbgen.Enum{}
	})...)
	a(p.printMessages(f, func(a *FieldAnnotation) pbgen.MessageField {
		d := pbgen.MessageField{
			Type:   a.PbType.ReferencedIn(f.PbPackage),
			Name:   a.Ident,
			Number: a.Number,
		}
		if a.IsRepeated {
			d.Labels = append(d.Labels, pbgen.LabelRepeated)
		} else if a.IsOptional {
			d.Labels = append(d.Labels, pbgen.LabelOptional)
		}

		return d
	})...)

	return v
}

type printerUtils struct {
	fs Fs
}

func (u *printerUtils) importPaths(f *ProtoFile) []pbgen.Import {
	paths := f.ImportPaths()
	slices.Sort(paths)

	v := make([]pbgen.Import, len(paths))
	for i, p := range paths {
		v[i] = pbgen.Import{Name: p}
	}

	return v
}

func (p *printerUtils) print(b *Build, print_file func(*ProtoFile) pbgen.ProtoFile) error {
	errs := []error{}
	for path, f := range b.Files {
		if err := p.fs.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return fmt.Errorf(`create directory for proto files: %w`, err)
		}

		w, err := p.fs.Create(path)
		if err != nil {
			return fmt.Errorf(`create proto file: %w`, err)
		}
		defer w.Close()

		v := print_file(f)
		if err := pbgen.Execute(w, v); err != nil {
			return fmt.Errorf(`generate proto file for "%s": %w`, path, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (u *printerUtils) printServices(f *ProtoFile) []pbgen.TopLevelDef {
	ds := []pbgen.TopLevelDef{}

	services := maps.Values(f.Services)
	slices.SortFunc(services, func(a, b *Service) int {
		return cmp.Compare(a.Ident, b.Ident)
	})
	for _, service := range services {
		d := pbgen.Service{Name: service.Ident}

		rpcs := maps.Values(service.Rpcs)
		slices.SortFunc(rpcs, func(a, b *Rpc) int {
			return cmp.Compare(a.Ident, b.Ident)
		})
		for _, rpc := range rpcs {
			if rpc.Comment != "" {
				d.Body = append(d.Body, pbgen.Comment{Value: rpc.Comment})
			}

			v := pbgen.Rpc{
				Name: rpc.Ident,
				Request: pbgen.RpcType{
					Type:   rpc.Req.ReferencedIn(f.PbPackage),
					Stream: (rpc.Stream & StreamClient) > 0,
				},
				Response: pbgen.RpcType{
					Type:   rpc.Res.ReferencedIn(f.PbPackage),
					Stream: (rpc.Stream & StreamServer) > 0,
				},
			}

			d.Body = append(d.Body, v)
		}

		ds = append(ds, d)
	}

	return ds
}

func (u *printerUtils) printEnums(f *ProtoFile, new_enum_def func(*Enum) pbgen.Enum) []pbgen.TopLevelDef {
	ds := []pbgen.TopLevelDef{}

	enums := maps.Values(f.Enums)
	slices.SortFunc(enums, func(a, b *Enum) int {
		return cmp.Compare(a.Ident, b.Ident)
	})
	for _, enum := range enums {
		if enum.Comment != "" {
			ds = append(ds, pbgen.Comment{Value: enum.Comment})
		}

		//
		// Field definitions
		//
		d := new_enum_def(enum)
		d.Name = enum.Ident
		fields := slices.Clone(enum.Fields)
		slices.SortFunc(fields, func(a, b *EnumField) int {
			return cmp.Compare(a.Number, b.Number)
		})
		for _, v := range fields {
			if v.Comment != "" {
				d.Body = append(d.Body, pbgen.Comment{Value: v.Comment})
			}
			d.Body = append(d.Body, pbgen.EnumField{Name: v.Name, Number: v.Number})
		}

		ds = append(ds, d)
	}

	return ds
}

func (u *printerUtils) printMessages(f *ProtoFile, print_field func(*FieldAnnotation) pbgen.MessageField) []pbgen.TopLevelDef {
	ds := []pbgen.TopLevelDef{}

	messages := maps.Values(f.Messages)
	slices.SortFunc(messages, func(a, b *MessageAnnotation) int {
		return cmp.Compare(a.Ident, b.Ident)
	})
	for _, message := range messages {
		if message.Comment != "" {
			ds = append(ds, pbgen.Comment{Value: message.Comment})
		}

		//
		// Field definitions
		//
		d := pbgen.Message{Name: message.Ident}
		fields := slices.Clone(message.Fields)
		slices.SortFunc(fields, func(a, b *FieldAnnotation) int {
			return cmp.Compare(a.Number, b.Number)
		})
		for _, field := range fields {
			if field.Comment != "" {
				d.Body = append(d.Body, pbgen.Comment{Value: field.Comment})
			}
			if field.IsOneof() {
				d_ := pbgen.MessageOneof{Name: field.Ident}
				for _, oneof_field := range field.Oneof {
					d_.Body = append(d_.Body, pbgen.MessageOneofField{
						Type:   oneof_field.PbType.ReferencedIn(f.PbPackage),
						Name:   oneof_field.Ident,
						Number: oneof_field.Number,
					})
				}
				d.Body = append(d.Body, d_)
			} else {
				d.Body = append(d.Body, print_field(field))
			}
		}

		ds = append(ds, d)
	}

	return ds
}
