package entpb

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen/ident"
)

const FieldAnnotation = "ProtoField"

type FieldOption interface {
	fieldOpt(*fieldAnnotation)
}

func Field(num int, opts ...FieldOption) schema.Annotation {
	a := &fieldAnnotation{Number: num}
	for _, opt := range opts {
		opt.fieldOpt(a)
	}
	return a
}

type fieldAnnotation struct {
	Ident   ident.Ident
	Number  int
	Comment string `mapstructure:"-"`

	EntName string          `mapstructure:"-"`
	EntInfo *field.TypeInfo `mapstructure:"-"`
	EntRef  string          `mapstructure:"-"` // Name of the schema that this edge references.
	PbType  PbType          `mapstructure:"-"`

	HasDefault bool `mapstructure:"-"`
	IsOptional bool `mapstructure:"-"`
	IsRepeated bool `mapstructure:"-"`
	IsReadOnly bool // Make this field cannot be set manually.
}

func (a *fieldAnnotation) IsEnum() bool {
	if a.EntInfo == nil {
		return false
	}
	return a.EntInfo.Type == field.TypeEnum
}

func (a *fieldAnnotation) IsEdge() bool {
	return a.EntInfo == nil
}

func (fieldAnnotation) Name() string {
	return FieldAnnotation
}

type fieldOptReadonly struct{}

func (o *fieldOptReadonly) fieldOpt(t *fieldAnnotation) {
	t.IsReadOnly = true
}

func WithReadOnly() FieldOption { return &fieldOptReadonly{} }
