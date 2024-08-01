package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash_id").Default(""),
		field.String("mobile").Default(""),
		field.String("password").Default(""),
		field.Int("level").Default(0),
		field.String("nickname").Default(""),
		field.String("avatar").Default(""),
		field.String("bio").Default(""),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.To("posts", Post.Type),
	}
}
