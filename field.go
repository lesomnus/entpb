package entpb

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/pbgen/ident"
)

const FieldAnnotationLabel = "ProtoField"

type FieldOption interface {
	fieldOpt(*FieldAnnotation)
}

func Field(num int, opts ...FieldOption) schema.Annotation {
	a := &FieldAnnotation{Number: num}
	for _, opt := range opts {
		opt.fieldOpt(a)
	}
	return a
}

type FieldAnnotation struct {
	Ident   ident.Ident
	Number  int
	Comment string `mapstructure:"-"`

	// `Number` will be set by minimum `Number` in this field.
	// This field is used to generate "oneof" in proto file.
	Oneof []*FieldAnnotation `mapstructure:"-"`

	EntName string          `mapstructure:"-"`
	EntInfo *field.TypeInfo `mapstructure:"-"`
	EntRef  string          `mapstructure:"-"` // Name of the schema that this edge references.
	PbType  PbType          `mapstructure:"-"`

	HasDefault  bool `mapstructure:"-"`
	IsKey       bool `mapstructure:"-"`
	IsOptional  bool `mapstructure:"-"`
	IsRepeated  bool `mapstructure:"-"`
	IsImmutable bool `mapstructure:"-"`
	IsReadOnly  bool // Make this field cannot be set manually.
}

func (a *FieldAnnotation) IsOneof() bool {
	return a.Oneof != nil
}

func (a *FieldAnnotation) IsEnum() bool {
	if a.EntInfo == nil {
		return false
	}
	return a.EntInfo.Type == field.TypeEnum
}

func (a *FieldAnnotation) IsEdge() bool {
	return a.EntInfo == nil
}

func (FieldAnnotation) Name() string {
	return FieldAnnotationLabel
}

type fieldOptReadonly struct{}

func (o *fieldOptReadonly) fieldOpt(t *FieldAnnotation) {
	t.IsReadOnly = true
}

func WithReadOnly() FieldOption { return &fieldOptReadonly{} }
