package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Default(0),
		field.Int("post_id").Optional().Default(0),
		field.String("content").Default(""),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("comments").Field("post_id").
			Unique(),
		edge.From("user", User.Type).
			Ref("comments").Field("user_id").
			Unique(),
	}
}
