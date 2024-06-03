package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb"
)

type Identity struct {
	ent.Schema
}

func (Identity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Identity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entpb.Field(3)).
			Comment("Name of the user").
			Default(""),
		field.String("email").
			Annotations(entpb.Field(4)).
			Optional().
			Nillable(),
	}
}

func (Identity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Annotations(entpb.Field(2)).
			Ref("identities").
			Immutable().
			Unique().
			Required(),
	}
}

func (Identity) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
