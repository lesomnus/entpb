package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

type Membership struct {
	ent.Schema
}

func (Membership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Membership) Fields() []ent.Field {
	return []ent.Field{}
}

func (Membership) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Membership) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
