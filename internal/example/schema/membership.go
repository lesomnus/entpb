package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/internal/example/role"
)

type Membership struct {
	ent.Schema
}

func (Membership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin{},
	}
}
func (Membership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}).
			Immutable(),
		field.UUID("team_id", uuid.UUID{}).
			Immutable(),

		field.Enum("role").
			Annotations(entpb.Field(6)).
			GoType(role.Role("")).
			Default(string(role.Member)),
	}
}

func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Annotations(entpb.Field(3)).
			Ref("memberships").
			Field("account_id").
			Immutable().
			Unique().
			Required(),
		edge.From("team", Team.Type).
			Annotations(entpb.Field(4)).
			Ref("members").
			Field("team_id").
			Immutable().
			Unique().
			Required(),
	}
}

func (Membership) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "team_id").Unique().
			Annotations(entpb.Key("by_account_in_team", 3)),
	}
}
