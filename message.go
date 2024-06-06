package entpb

import (
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"github.com/lesomnus/entpb/pbgen/ident"
	"golang.org/x/exp/maps"
)

const MessageAnnotation = "ProtoMessage"

type MessageOption interface {
	messageOpt(*messageAnnotation)
}

func Message(filepath string, opts ...MessageOption) schema.Annotation {
	a := &messageAnnotation{Filepath: filepath}
	for _, opt := range opts {
		opt.messageOpt(a)
	}
	return a
}

type messageAnnotation struct {
	Filepath string

	Ident   ident.Ident
	Comment string `mapstructure:"-"`

	Service *service

	File   *ProtoFile         `mapstructure:"-"`
	Schema *load.Schema       `mapstructure:"-"`
	Fields []*fieldAnnotation `mapstructure:"-"`
}

func (a *messageAnnotation) pbType() PbType {
	t := PbType{
		Name:   a.Ident,
		Import: a.Filepath,
	}
	if f := a.File; f != nil {
		t.Package = f.pbPackage
	}

	return t
}

func (messageAnnotation) Name() string {
	return MessageAnnotation
}

func (a messageAnnotation) Merge(other schema.Annotation) schema.Annotation {
	a_, ok := other.(*messageAnnotation)
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

			builtin := maps.Clone(lhs.BuiltIn)
			maps.Copy(builtin, rhs.BuiltIn)
			rhs.BuiltIn = builtin
		}
	}

	return other
}

const PathInherit = "$inherit"
