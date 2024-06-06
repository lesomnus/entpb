package entpb

import (
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen/ident"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const ProtoFilesAnnotation = "ProtoFiles"

type ProtoFile struct {
	path string

	pbPackage ident.Full
	goPackage string

	enums    map[string]*enum // key is global type name of the type bound to mapping enum.
	messages map[ident.Ident]*messageAnnotation
	services map[ident.Ident]*service
}

type ProtoFileInit struct {
	PbPackage ident.Full
	GoPackage string
}

func NewProtoFile(init ProtoFileInit) *ProtoFile {
	return &ProtoFile{
		pbPackage: init.PbPackage,
		goPackage: init.GoPackage,

		enums:    map[string]*enum{},
		messages: map[ident.Ident]*messageAnnotation{},
		services: map[ident.Ident]*service{},
	}
}

type ProtoFiles map[string]*ProtoFile

func (ProtoFiles) Name() string {
	return ProtoFilesAnnotation
}

func (f *ProtoFile) ImportPaths() []string {
	ps := map[string]struct{}{}
	for _, message := range f.messages {
		for _, field := range message.Fields {
			ps[field.PbType.Import] = struct{}{}
		}
	}
	for _, service := range f.services {
		for _, rpc := range service.Rpcs {
			ps[rpc.Req.Import] = struct{}{}
			ps[rpc.Res.Import] = struct{}{}
		}
	}
	delete(ps, "")
	delete(ps, f.path)

	return maps.Keys(ps)
}

func ForwardDeclarations(files map[string]*protogen.File, graph *gen.Graph) {
	d := ProtoFiles{}
	decodeAnnotation(&d, graph.Annotations)
	for p, f := range d {
		f.path = p
	}

	for _, s := range graph.Schemas {
		d_m, ok := decodeAnnotation(&messageAnnotation{}, s.Annotations)
		if !ok {
			continue
		}

		_, ok = d[d_m.Filepath]
		if ok {
			continue
		}

		if s := d_m.Service; s != nil {
			d[s.Filepath] = NewProtoFile(ProtoFileInit{})
		}

		file := NewProtoFile(ProtoFileInit{})
		d[d_m.Filepath] = file

		for _, f := range s.Fields {
			if f.Info.Type != field.TypeEnum {
				continue
			}
			if f.Info.RType == nil {
				// Not an external type such as `example.Role`
				continue
			}
			if _, ok := decodeAnnotation(&fieldAnnotation{}, f.Annotations); !ok {
				continue
			}

			// Resolve name of the enum from the proto file.
			// We cannot use `f.Info.RType.Name` since it can be renamed.
			proto_file, ok := files[d_m.Filepath]
			if !ok {
				// Maybe the file is not given as a protoc input?
				continue
			}

			var proto_enum *protogen.Enum
		L:
			for _, message := range proto_file.Messages {
				name := string(d_m.Ident)
				if name == "" {
					name = s.Name
				}
				if name != string(message.Desc.Name()) {
					continue
				}

				for _, field := range message.Fields {
					if field.Desc.Name() != protoreflect.Name(f.Name) {
						continue
					}
					if field.Desc.Kind() != protoreflect.EnumKind {
						panic("looking for the enum but type in the proto file is not an enum")
					}

					proto_enum = field.Enum
					break L
				}
			}
			if proto_enum == nil {
				panic("enum not found")
			}

			z := ""
			r := f.Info.RType
			enum := &enum{
				ident: ident.Ident(r.Name),

				// Values come from proto file so the keys are already prefixed.
				prefix: &z,
			}
			for _, v := range proto_enum.Values {
				enum.vs = append(enum.vs, EnumField{
					Name:   string(v.Desc.Name()),
					Number: int(v.Desc.Number()),
				})
			}

			name := globalTypeName(r.PkgPath, r.Ident)
			file.enums[name] = enum
		}
	}
	for p, f := range d {
		f.path = p
	}

	for p, proto_file := range files {
		f, ok := d[p]
		if !ok {
			continue
		}

		if pkg := proto_file.Proto.GetPackage(); pkg != "" {
			f.pbPackage = ident.Must(strings.Split(pkg, "."))
		}
		f.goPackage = string(proto_file.GoPackageName)
	}

	graph.Annotations.Set(d.Name(), d)
}
