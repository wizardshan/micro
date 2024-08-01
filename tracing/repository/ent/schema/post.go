package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash_id").Default(""),
		field.Int("user_id").Optional().Default(0),
		field.String("title").Default(""),
		field.String("content").Default(""),
		field.Int("times_of_read").Default(0),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("user", User.Type).
			Ref("posts").Field("user_id").
			Unique(),
	}
}
