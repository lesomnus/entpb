package entpb

import (
	"github.com/lesomnus/entpb/pbgen/ident"
	"golang.org/x/exp/maps"
)

const ProtoFilesAnnotation = "ProtoFiles"

type ProtoFile struct {
	path  string
	alias string

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

func NewProtoFile(init ProtoFileInit) ProtoFile {
	return ProtoFile{
		pbPackage: init.PbPackage,
		goPackage: init.GoPackage,

		enums:    map[string]*enum{},
		messages: map[ident.Ident]*messageAnnotation{},
		services: map[ident.Ident]*service{},
	}
}

type ProtoFiles map[string]ProtoFile

func (ProtoFiles) Name() string {
	return ProtoFilesAnnotation
}

func (f ProtoFile) AliasAs(v string) ProtoFile {
	f.alias = v
	return f
}

func (f *ProtoFile) ImportPaths() []string {
	ps := map[string]struct{}{}
	for _, message := range f.messages {
		for _, field := range message.fields {
			ps[field.pb_type.Import] = struct{}{}
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
