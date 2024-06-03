package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"github.com/lesomnus/entpb"
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
	return []ent.Edge{
		edge.From("owner", User.Type).
			Annotations(entpb.Field(2)).
			Ref("memberships").
			Immutable().
			Unique().
			Required(),
	}
}

func (Membership) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
