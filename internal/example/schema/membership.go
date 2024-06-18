package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
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
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}).
			Immutable(),
		field.String("name").
			Annotations(entpb.Field(3)).
			DefaultFunc(uuid.NewString),
	}
}

func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Annotations(entpb.Field(2)).
			Ref("memberships").
			Field("account_id").
			Immutable().
			Unique().
			Required(),
	}
}

func (Membership) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (Membership) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "name").Unique(),
	}
}
