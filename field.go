package entpb

import (
	"entgo.io/ent/schema"
	"github.com/lesomnus/entpb/pbgen/ident"
)

const FieldAnnotation = "ProtoField"

type MessageFieldOption interface {
	messageFieldOpt(*messageFieldAnnotation)
}

func Field(num int, opts ...MessageFieldOption) schema.Annotation {
	a := &messageFieldAnnotation{Number: num}
	for _, opt := range opts {
		opt.messageFieldOpt(a)
	}
	return a
}

type messageFieldAnnotation struct {
	Ident  ident.Ident
	Number int

	pb_type PbType

	comment string

	isOptional bool
	isRepeated bool
}

func (messageFieldAnnotation) Name() string {
	return FieldAnnotation
}
