package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb"
	"github.com/lesomnus/entpb/example"
)

type Shape string

const (
	Triangle Shape = "TRIANGLE"
	Circle   Shape = "CIRCLE"
)

// Values provides list valid values for Enum.
func (Shape) Values() (kinds []string) {
	for _, s := range []Shape{Triangle, Circle} {
		kinds = append(kinds, string(s))
	}
	return
}

type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			Annotations(entpb.Field(4)).
			GoType(example.Role("")),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Annotations(entpb.Field(2)).
			Ref("accounts").
			Immutable().
			Unique().
			Required(),
	}
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
