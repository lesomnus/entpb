package entpb

import "reflect"

type EnumOption func(*enum)

type enum struct {
	t  reflect.Type
	vs []EnumField

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

		prefix: &z,
	}
	for _, opt := range opts {
		opt(v)
	}

	f.enums[globalTypeNameFromReflect(r)] = v
	return f
}

func WithPrefix(v string) EnumOption {
	return func(e *enum) {
		e.prefix = &v
	}
}

func WithNoPrefix() EnumOption {
	return func(e *enum) {
		e.prefix = nil
	}
}

func WithOpen() EnumOption {
	return func(e *enum) {
		e.is_closed = false
	}
}

func WithClose() EnumOption {
	return func(e *enum) {
		e.is_closed = true
	}
}
