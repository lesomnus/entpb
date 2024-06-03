package entpb

import (
	"entgo.io/ent/schema"
	"github.com/lesomnus/entpb/pbgen/ident"
)

const FieldAnnotation = "ProtoField"

func Field(num int) schema.Annotation {
	a := &fieldAnnotation{Number: num}
	// for _, apply := range opts {
	// 	apply(&a)
	// }
	return a
}

type fieldAnnotation struct {
	Number int

	pb_type PbType

	name    ident.Ident
	comment string

	isOptional bool
	isRepeated bool
}

func (fieldAnnotation) Name() string {
	return FieldAnnotation
}
