package entpb

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type nameOption struct{ v ident.Ident }

func WithName(v ident.Ident) *nameOption {
	return &nameOption{v}
}

func (o *nameOption) messageOpt(t *MessageAnnotation) { t.Ident = o.v }
func (o *nameOption) fieldOpt(t *FieldAnnotation)     { t.Ident = o.v }
func (o *nameOption) enumOpt(t *Enum)                 { t.Ident = o.v }
func (o *nameOption) serviceOpt(t *Service)           { t.Ident = o.v }

type commentOption struct{ v string }

func WithComment(v string) *commentOption {
	return &commentOption{v}
}

func (o *commentOption) enumOpt(t *Enum) { t.Comment = o.v }
