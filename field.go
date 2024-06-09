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

	IsExplicitReadOnly bool
	IsExplicitWritable bool
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

func (a *FieldAnnotation) IsReadOnly() bool {
	if a.IsEdge() {
		return !a.IsExplicitWritable
	} else {
		return a.IsExplicitReadOnly
	}
}

func (a *FieldAnnotation) IsWritable() bool {
	return !a.IsReadOnly()
}

func (FieldAnnotation) Name() string {
	return FieldAnnotationLabel
}

type fieldOptReadOnly struct{}

func (o *fieldOptReadOnly) fieldOpt(t *FieldAnnotation) {
	t.IsExplicitReadOnly = true
}

func WithReadOnly() FieldOption { return &fieldOptReadOnly{} }

type fieldOptWritable struct{}

func (o *fieldOptWritable) fieldOpt(t *FieldAnnotation) {
	t.IsExplicitWritable = true
}

func WithWritable() FieldOption { return &fieldOptWritable{} }
