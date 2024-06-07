package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/lesomnus/entpb"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Immutable().
			Default(uuid.New).
			Annotations(entpb.Field(1, entpb.WithReadOnly())),

		field.Time("date_created").
			Immutable().
			Default(time.Now).
			Annotations(entpb.Field(15, entpb.WithReadOnly())),
	}
}

func (BaseMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entpb.Message("entpb/directory/common.proto",
			entpb.WithService("entpb/directory/service.proto",
				entpb.RpcEntCreate(),
			),
		),
	}
}
