package entpb

import (
	"entgo.io/ent/schema"
	"github.com/lesomnus/entpb/pbgen/ident"
)

const NameAnnotation = "ProtoName"

func Name(v ident.Ident) schema.Annotation {
	return &nameAnnotation{Value: v}
}

type nameAnnotation struct {
	Value ident.Ident
}

func (nameAnnotation) Name() string {
	return NameAnnotation
}

func (nameAnnotation) Merge(other schema.Annotation) schema.Annotation {
	return other
}
