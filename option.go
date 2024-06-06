package entpb

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type nameOption struct{ v ident.Ident }

func WithName(v ident.Ident) *nameOption {
	return &nameOption{v}
}

func (o *nameOption) messageOpt(t *messageAnnotation) { t.Ident = o.v }
func (o *nameOption) fieldOpt(t *fieldAnnotation)     { t.Ident = o.v }
func (o *nameOption) enumOpt(t *enum)                 { t.ident = o.v }

type commentOption struct{ v string }

func WithComment(v string) *commentOption {
	return &commentOption{v}
}

func (o *commentOption) enumOpt(t *enum) { t.comment = o.v }
