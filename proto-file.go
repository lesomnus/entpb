package entpb

import (
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/iancoleman/strcase"
	"github.com/lesomnus/entpb/pbgen/ident"
	"github.com/lesomnus/entpb/utils"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const ProtoFilesAnnotation = "ProtoFiles"

type ProtoFileOption interface {
	protoFileOpt(*ProtoFile)
}

type ProtoFile struct {
	Filepath string

	PbPackage ident.Full
	GoPackage string

	Enums    map[string]*Enum // key is global type name of the type bound to mapping enum.
	Messages map[ident.Ident]*MessageAnnotation
	Services map[ident.Ident]*Service
}

type ProtoFileInit struct {
	PbPackage ident.Full
	GoPackage string
}

func NewProtoFile(init ProtoFileInit, opts ...ProtoFileOption) *ProtoFile {
	v := &ProtoFile{
		PbPackage: init.PbPackage,
		GoPackage: init.GoPackage,

		Enums:    map[string]*Enum{},
		Messages: map[ident.Ident]*MessageAnnotation{},
		Services: map[ident.Ident]*Service{},
	}
	for _, opt := range opts {
		opt.protoFileOpt(v)
	}

	return v
}

type ProtoFiles map[string]*ProtoFile

func (ProtoFiles) Name() string {
	return ProtoFilesAnnotation
}

func (f *ProtoFile) ImportPaths() []string {
	ps := map[string]struct{}{}
	for _, message := range f.Messages {
		for _, field := range message.Fields {
			ps[field.PbType.Import] = struct{}{}
		}
	}
	for _, service := range f.Services {
		for _, rpc := range service.Rpcs {
			ps[rpc.Req.Import] = struct{}{}
			ps[rpc.Res.Import] = struct{}{}
		}
	}
	delete(ps, "")
	delete(ps, f.Filepath)

	return maps.Keys(ps)
}

func ForwardDeclarations(files map[string]*protogen.File, graph *gen.Graph) {
	d := ProtoFiles{}
	DecodeAnnotation(&d, graph.Annotations)
	for p, f := range d {
		f.Filepath = p
	}

	for _, s := range graph.Schemas {
		d_m, ok := DecodeAnnotation(&MessageAnnotation{}, s.Annotations)
		if !ok {
			continue
		}

		_, ok = d[d_m.Filepath]
		if ok {
			continue
		}

		// FIXME: it maybe overwrites existing file.
		// I think above continue must be fixed first.
		if s := d_m.Service; s != nil {
			d[s.Filepath] = NewProtoFile(ProtoFileInit{})
		}

		file := NewProtoFile(ProtoFileInit{})
		d[d_m.Filepath] = file

		for _, ent_field := range s.Fields {
			if ent_field.Info.Type != field.TypeEnum {
				continue
			}
			if ent_field.Info.RType == nil {
				// Not an external type such as `example.Role`
				continue
			}
			if _, ok := DecodeAnnotation(&FieldAnnotation{}, ent_field.Annotations); !ok {
				continue
			}

			// Resolve name of the enum from the proto file.
			// We cannot use `f.Info.RType.Name` since it can be renamed.
			pb_file, ok := files[d_m.Filepath]
			if !ok {
				// Maybe the file is not given as a protoc input?
				continue
			}

			var pb_enum *protogen.Enum
			for _, message := range pb_file.Messages {
				name := string(d_m.Ident)
				if name == "" {
					name = s.Name
				}
				if name != string(message.Desc.Name()) {
					continue
				}

				for _, field := range message.Fields {
					if field.Desc.Name() != protoreflect.Name(ent_field.Name) {
						continue
					}
					if field.Desc.Kind() != protoreflect.EnumKind {
						panic("looking for the enum but type in the proto file is not an enum")
					}

					pb_enum = field.Enum
					break
				}
				if pb_enum != nil {
					break
				}
			}
			if pb_enum == nil {
				panic("enum not found")
			}

			z := ""
			r := ent_field.Info.RType
			enum := &Enum{
				GoType: *r,

				Ident: ident.Ident(pb_enum.Desc.Name()),

				File: file,

				// Values come from proto file so the keys are already prefixed.
				Prefix: &z,
			}
			if len(pb_enum.Values) == 0 {
				panic("proto doest not have enum values")
			}
			if len(ent_field.Enums) == 0 {
				panic("ent does not have enum values")
			}

			// TODO: this is naive implementation.
			prefix := utils.GuessPrefix(
				utils.Map(pb_enum.Values, func(v *protogen.EnumValue) string { return string(v.Desc.Name()) }),
				utils.Map(ent_field.Enums, func(v struct{ N, V string }) string { return strings.ToUpper(strcase.ToSnake(v.V)) }),
			)
			for _, v := range pb_enum.Values {
				name := string(v.Desc.Name())
				remain, ok := strings.CutPrefix(name, prefix)
				if !ok {
					panic("guessing prefix failed?")
				}

				enum.Fields = append(enum.Fields, &EnumField{
					Name:   name,
					Number: int(v.Desc.Number()),
					Value:  remain,
				})
			}

			name := utils.FullIdent(r.PkgPath, r.Ident)
			file.Enums[name] = enum
		}
	}
	for p, f := range d {
		f.Filepath = p
	}

	for p, proto_file := range files {
		f, ok := d[p]
		if !ok {
			continue
		}

		if pkg := proto_file.Proto.GetPackage(); pkg != "" {
			segs := strings.Split(pkg, ".")
			f.PbPackage = ident.Must(segs[0], segs...)
		}
		f.GoPackage = string(proto_file.GoPackageName)
	}

	graph.Annotations.Set(d.Name(), d)
}
