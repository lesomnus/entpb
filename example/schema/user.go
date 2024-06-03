package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"github.com/lesomnus/entpb"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", User.Type).
			Annotations(entpb.Field(2)).
			From("parent").
			Annotations(entpb.Field(3)).
			Unique(),
		edge.To("identities", Identity.Type),
		edge.To("accounts", Account.Type),
		edge.To("memberships", Membership.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("User is an actor"),
	}
}
