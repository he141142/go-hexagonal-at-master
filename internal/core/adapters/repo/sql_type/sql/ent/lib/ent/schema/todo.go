package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	schema_helper "hex-base/internal/core/utils"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").SchemaType(
			map[string]string{
				dialect.Postgres: string(schema_helper.INTEGER),
			}),
		field.String("name").SchemaType(
			map[string]string{
				dialect.Postgres: schema_helper.Varchar(50),
			}).Optional(),
		field.String("task").SchemaType(
			map[string]string{
				dialect.Postgres: schema_helper.Varchar(200),
			}).Optional(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("form", Form.Type).Annotations(
			entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "todo"},
	}
}
