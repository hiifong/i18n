package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type EntI18n struct {
	ent.Schema
}

func (EntI18n) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "i18n"},
	}
}

func (EntI18n) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("lang"),
		field.String("key"),
		field.String("raw"),
		field.Int64("created").Immutable().DefaultFunc(func() int64 {
			return time.Now().Unix()
		}),
		field.Int64("updated").Immutable().DefaultFunc(func() int64 {
			return time.Now().Unix()
		}).UpdateDefault(func() int64 {
			return time.Now().Unix()
		}),
		field.Int64("deleted").Optional().Nillable(),
	}
}

func (EntI18n) Edges() []ent.Edge {
	return nil
}
