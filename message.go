package entpb

import (
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"github.com/lesomnus/entpb/pbgen/ident"
	"golang.org/x/exp/maps"
)

const MessageAnnotationLabel = "ProtoMessage"

type MessageOption interface {
	messageOpt(*MessageAnnotation)
}

func Message(filepath string, opts ...MessageOption) schema.Annotation {
	a := &MessageAnnotation{Filepath: filepath}
	for _, opt := range opts {
		opt.messageOpt(a)
	}
	return a
}

type MessageAnnotation struct {
	Filepath string

	Ident   ident.Ident
	Comment string `mapstructure:"-"`

	Service *Service

	File *ProtoFile `mapstructure:"-"`
	// Schema that is referenced to generate this message.
	// For example, `User` and `GetUserRequest` has same `Schema` value.
	Schema *load.Schema       `mapstructure:"-"`
	Fields []*FieldAnnotation `mapstructure:"-"`
}

func (a *MessageAnnotation) pbType() PbType {
	t := PbType{
		Ident:  a.Ident,
		Import: a.Filepath,
	}
	if f := a.File; f != nil {
		t.Package = f.PbPackage
	}

	return t
}

func (MessageAnnotation) Name() string {
	return MessageAnnotationLabel
}

func (a MessageAnnotation) Merge(other schema.Annotation) schema.Annotation {
	a_, ok := other.(*MessageAnnotation)
	if !ok {
		panic("invalid annotation")
	}
	if a_.Filepath == PathInherit {
		a_.Filepath = a.Filepath
	}
	if a.Service != nil {
		if a_.Service == nil {
			a_.Service = a.Service
		} else {
			lhs := a.Service
			rhs := a_.Service

			if rhs.Filepath == PathInherit {
				rhs.Filepath = lhs.Filepath
			}

			rpcs := maps.Clone(lhs.Rpcs)
			maps.Copy(rpcs, rhs.Rpcs)
			rhs.Rpcs = rpcs
		}
	}

	return other
}

const PathInherit = "$inherit"
