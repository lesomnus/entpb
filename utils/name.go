package utils

import (
	"fmt"
	"reflect"

	"entgo.io/ent/schema/field"
)

func FullIdent(path string, ident string) string {
	return fmt.Sprintf("%s:%s", path, ident)
}

func FullIdentFromReflect(t reflect.Type) string {
	return FullIdent(t.PkgPath(), t.String())
}

func FullIdentFromEntTypeInfo(t *field.TypeInfo) string {
	return FullIdent(t.PkgPath, t.Ident)
}
