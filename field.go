package entpb

import "entgo.io/ent/schema"

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

	name    string
	comment string

	isOptional bool
	isRepeated bool
}

func (fieldAnnotation) Name() string {
	return FieldAnnotation
}
