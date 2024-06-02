package entpb

import (
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
)

const MessageAnnotation = "ProtoMessage"

func Message(filepath string) schema.Annotation {
	return &messageAnnotation{Filepath: filepath}
}

type messageAnnotation struct {
	Filepath string

	ref    *load.Schema
	fields []*fieldAnnotation

	name    string
	comment string
}

func (messageAnnotation) Name() string {
	return MessageAnnotation
}

func (messageAnnotation) Merge(other schema.Annotation) schema.Annotation {
	return other
}
