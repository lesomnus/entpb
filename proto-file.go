package entpb

import (
	"fmt"
	"reflect"

	"entgo.io/ent/schema/field"
)

const ProtoFilesAnnotation = "ProtoFiles"

type ProtoFile struct {
	path  string
	alias string

	pbPackage string
	goPackage string

	enums    map[string]*enum // key is global type name of the type bound to mapping enum.
	messages map[string]*messageAnnotation
}

type ProtoFileInit struct {
	PbPackage string
	GoPackage string
}

func NewProtoFile(init ProtoFileInit) ProtoFile {
	return ProtoFile{
		pbPackage: init.PbPackage,
		goPackage: init.GoPackage,

		enums:    map[string]*enum{},
		messages: map[string]*messageAnnotation{},
	}
}

type enum struct {
	t  reflect.Type
	vs map[string]int
}

type ProtoFiles map[string]ProtoFile

func (ProtoFiles) Name() string {
	return ProtoFilesAnnotation
}

func (f ProtoFile) AliasAs(v string) ProtoFile {
	f.alias = v
	return f
}

func (f ProtoFile) AddEnum(t any, values map[string]int) ProtoFile {
	r := reflect.TypeOf(t)
	f.enums[globalTypeNameFromReflect(r)] = &enum{
		t:  r,
		vs: values,
	}
	return f
}

func globalTypeName(path string, ident string) string {
	return fmt.Sprintf("%s:%s", path, ident)
}

func globalTypeNameFromReflect(t reflect.Type) string {
	return globalTypeName(t.PkgPath(), t.String())
}

func globalTypeNameFromEntTypeInfo(t *field.TypeInfo) string {
	return globalTypeName(t.PkgPath, t.Ident)
}
