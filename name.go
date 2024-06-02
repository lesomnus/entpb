package entpb

import "entgo.io/ent/schema"

const NameAnnotation = "ProtoName"

func Name(v string) schema.Annotation {
	return &nameAnnotation{Value: v}
}

type nameAnnotation struct {
	Value string
}

func (nameAnnotation) Name() string {
	return NameAnnotation
}

func (nameAnnotation) Merge(other schema.Annotation) schema.Annotation {
	return other
}
