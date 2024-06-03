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
	Print(b *Build, fs Fs) error
}

func NewPrinter(edition pbgen.Edition) (Printer, error) {
	switch edition {
	case pbgen.Edition2023:
		return &edition2023Printer{}, nil

	case pbgen.SyntaxProto2:
		fallthrough
	case pbgen.SyntaxProto3:
		fallthrough
	default:
		return nil, errors.New("not implemented")
	}
}

type edition2023Printer struct{}

func (p *edition2023Printer) Print(b *Build, fs Fs) error {
	for p, f := range b.files {
		if f.path != p {
			// Link by alias.
			continue
		}

		proto_file := pbgen.ProtoFile{
			Edition: pbgen.Edition2023,
			Package: f.pbPackage,
			Options: []pbgen.Option{pbgen.FeatureFieldPresenceImplicit},
		}
		if f.goPackage != "" {
			proto_file.Options = append(proto_file.Options, pbgen.OptionGoPackage(f.goPackage))
		}

		// Collect imports.
		{
			imports := map[string]struct{}{}
			for _, message := range f.messages {
				for _, field := range message.fields {
					imports[field.pb_type.Import] = struct{}{}
				}
			}
			for _, service := range f.services {
				for _, rpc := range service.Rpcs {
					imports[rpc.Req.Import] = struct{}{}
					imports[rpc.Res.Import] = struct{}{}
				}
			}
			delete(imports, "")
			delete(imports, f.path)

			paths := maps.Keys(imports)
			slices.Sort(paths)
			for _, p := range paths {
				proto_file.Imports = append(proto_file.Imports, pbgen.Import{Name: p})
			}
		}

		//
		// Service definitions.
		//
		services := maps.Values(f.services)
		slices.SortFunc(services, func(a, b *service) int {
			return cmp.Compare(a.Name, b.Name)
		})
		for _, service := range services {
			d := pbgen.Service{Name: service.Name}

			rpcs := maps.Values(service.Rpcs)
			slices.SortFunc(rpcs, func(a, b *Rpc) int {
				return cmp.Compare(a.Name, b.Name)
			})
			for _, v := range rpcs {
				if v.Comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.Comment})
				}

				elem := pbgen.Rpc{Name: v.Name}
				elem.Request.Type = v.Req.ReferencedIn(f.pbPackage)
				elem.Response.Type = v.Res.ReferencedIn(f.pbPackage)
				switch v.Stream {
				case StreamNone:
				case StreamClient:
					elem.Request.Stream = true
				case StreamServer:
					elem.Response.Stream = true
				case StreamBoth:
					elem.Request.Stream = true
					elem.Response.Stream = true

				default:
					// ignore
				}

				d.Body = append(d.Body, elem)
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		//
		// Enum definitions
		//
		enums := maps.Values(f.enums)
		slices.SortFunc(enums, func(a, b *enum) int {
			return cmp.Compare(a.t.Name(), b.t.Name())
		})
		for _, enum := range enums {
			// TODO: comment for enum

			//
			// Field definitions
			//
			d := pbgen.Enum{Name: enum.t.Name()}
			if enum.is_closed {
				d.Options = append(d.Options, pbgen.FeatureEnumTypeClosed)
			}
			fields := slices.Clone(enum.vs)
			slices.SortFunc(fields, func(a, b EnumField) int {
				return cmp.Compare(a.Number, b.Number)
			})
			for _, v := range fields {
				if v.Comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.Comment})
				}
				d.Body = append(d.Body, pbgen.EnumField{Name: v.Name, Number: v.Number})
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		//
		// Message definitions
		//
		messages := maps.Values(f.messages)
		slices.SortFunc(messages, func(a, b *messageAnnotation) int {
			return cmp.Compare(a.name, b.name)
		})
		for _, message := range messages {
			if message.comment != "" {
				proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, pbgen.Comment{Value: message.comment})
			}

			//
			// Field definitions
			//
			d := pbgen.Message{Name: message.name}
			fields := slices.Clone(message.fields)
			slices.SortFunc(fields, func(a, b *fieldAnnotation) int {
				return cmp.Compare(a.Number, b.Number)
			})
			for _, v := range fields {
				if v.comment != "" {
					d.Body = append(d.Body, pbgen.Comment{Value: v.comment})
				}

				elem := pbgen.MessageField{
					Type:   v.pb_type.ReferencedIn(f.pbPackage),
					Name:   v.name,
					Number: v.Number,
				}
				if v.isRepeated {
					elem.Labels = append(elem.Labels, pbgen.LabelRepeated)
				} else if v.isOptional {
					// Presence of "repeated" fields are not tracked.
					elem.Options = append(elem.Options, pbgen.FeatureFieldPresenceExplicit)
				}

				d.Body = append(d.Body, elem)
			}

			proto_file.TopLevelDefinitions = append(proto_file.TopLevelDefinitions, d)
		}

		if err := fs.MkdirAll(filepath.Dir(p), 0755); err != nil {
			return fmt.Errorf(`create directory for proto files: %w`, err)
		}

		w, err := fs.Create(p)
		if err != nil {
			return fmt.Errorf(`create proto file: %w`, err)
		}

		defer w.Close()
		if err := pbgen.Execute(w, proto_file); err != nil {
			return fmt.Errorf(`generate proto file for "%s": %w`, p, err)
		}
	}

	return nil
}
