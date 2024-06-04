package entpb

import (
	"reflect"

	"github.com/lesomnus/entpb/pbgen/ident"
)

type EnumOption interface {
	enumOpt(*enum)
}

type enum struct {
	t  reflect.Type
	vs []EnumField

	ident   ident.Ident
	comment string

	prefix    *string
	is_closed bool
}

type EnumField struct {
	Name    string
	Number  int
	Comment string
}

// Field names are prefixed by enum name; e.g. "owner" of `enum Role` will become "ROLE_OWNER".
// Change the prefix using `WithPrefix`, or disable the prefix using `WithNoPrefix`.
// Zero value will be generated with postfix "UNSPECIFIED" if the zero value not given; e.g. zero value of `enum Role` will be "ROLE_UNSPECIFIED".
// Zero value will not be generated if enum is closed by `WithClose` option.
func (f ProtoFile) AddEnum(t any, fields []EnumField, opts ...EnumOption) ProtoFile {
	z := ""

	r := reflect.TypeOf(t)
	v := &enum{
		t:  r,
		vs: fields,

		ident: ident.Ident(r.Name()),

		prefix: &z,
	}
	for _, opt := range opts {
		opt.enumOpt(v)
	}

	f.enums[globalTypeNameFromReflect(r)] = v
	return f
}

type enumOptPrefix struct{ v *string }

func (o *enumOptPrefix) enumOpt(t *enum) {
	t.prefix = o.v
}

func WithPrefix(v string) EnumOption { return &enumOptPrefix{&v} }
func WithNoPrefix() EnumOption       { return &enumOptPrefix{nil} }

type enumOptBound struct{ v bool }

func (o *enumOptBound) enumOpt(t *enum) {
	t.is_closed = o.v
}

func WithOpen() EnumOption  { return &enumOptBound{false} }
func WithClose() EnumOption { return &enumOptBound{true} }
