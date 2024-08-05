package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb"
)

type Invitation struct {
	ent.Schema
}

func (Invitation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin{},
	}
}

func (Invitation) Fields() []ent.Field {
	return []ent.Field{
		field.String("invitee").
			Annotations(entpb.Field(5)).
			Immutable().
			NotEmpty(),

		field.String("type").
			Annotations(entpb.Field(6)).
			Immutable().
			NotEmpty(),

		field.Time("date_expired").
			Annotations(entpb.Field(14)),
		field.Time("date_accepted").
			Annotations(entpb.Field(13)).
			Optional().
			Nillable(),
		field.Time("date_declined").
			Annotations(entpb.Field(12)).
			Optional().
			Nillable(),
		field.Time("date_canceled").
			Annotations(entpb.Field(11)).
			Optional().
			Nillable(),
	}
}

func (Invitation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inviter", Account.Type).
			Annotations(entpb.Field(3)).
			Ref("invitations").
			Immutable().
			Unique().
			Required(),
		edge.From("silo", Silo.Type).
			Annotations(entpb.Field(4)).
			Ref("invitations").
			Immutable().
			Unique().
			Required(),
	}
}
