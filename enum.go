package entpb

import (
	"reflect"

	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen/ident"
)

type EnumOption interface {
	enumOpt(*enum)
}

type enum struct {
	// Go type that is mapped with this enum.
	// e.g. `example.Role`.
	GoType field.RType

	// Name in proto.
	// It may be different from the name of Go type.
	// e.g. `GroupRole`.
	Ident   ident.Ident
	Comment string

	File   *ProtoFile // Proto file that defines this enum.
	Fields []*EnumField

	Prefix   *string
	IsClosed bool
}

type EnumField struct {
	// Name in proto.
	// e.g. `GroupRole_GROUP_ROLE_OWNER`.
	Name    string
	Number  int
	Comment string

	// Mapped value in Go.
	// e.g. `OWNER` which is `example.RoleOwner`.
	Value string
}

type EnumDesc struct {
	Number  int
	Comment string
}

type protoFileOptAddEnum[T ~string] struct {
	name string
	enum *enum
}

func (o protoFileOptAddEnum[T]) protoFileOpt(t *ProtoFile) {
	t.enums[o.name] = o.enum
}

// Field names are prefixed by enum name; e.g. "owner" of `enum Role` will become "ROLE_OWNER".
// Change the prefix using `WithPrefix`, or disable the prefix using `WithNoPrefix`.
// Zero value will be generated with postfix "UNSPECIFIED" if the zero value not given; e.g. zero value of `enum Role` will be "ROLE_UNSPECIFIED".
// Zero value will not be generated if enum is closed by `WithClose` option.
func WithEnum[T ~string](fields map[T]EnumDesc, opts ...EnumOption) ProtoFileOption {
	var t T
	r := reflect.TypeOf(t)
	z := ""

	v := &enum{
		GoType: field.RType{
			Name:    r.Name(),
			Ident:   r.String(),
			PkgPath: r.PkgPath(),
		},
		Ident: ident.Ident(r.Name()),

		Prefix: &z,
	}
	for name, d := range fields {
		v.Fields = append(v.Fields, &EnumField{
			Name:    string(name),
			Number:  d.Number,
			Comment: d.Comment,

			Value: string(name),
		})
	}
	for _, opt := range opts {
		opt.enumOpt(v)
	}

	name := globalTypeNameFromReflect(r)
	return &protoFileOptAddEnum[T]{name, v}
}

type enumOptPrefix struct{ v *string }

func (o *enumOptPrefix) enumOpt(t *enum) {
	t.Prefix = o.v
}

func WithPrefix(v string) EnumOption { return &enumOptPrefix{&v} }
func WithNoPrefix() EnumOption       { return &enumOptPrefix{nil} }

type enumOptBound struct{ v bool }

func (o *enumOptBound) enumOpt(t *enum) {
	t.IsClosed = o.v
}

func WithOpen() EnumOption  { return &enumOptBound{false} }
func WithClose() EnumOption { return &enumOptBound{true} }
