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

// Form holds the schema definition for the Form entity.
type Form struct {
	ent.Schema
}

// Fields of the Form.
func (Form) Fields() []ent.Field {
	 return []ent.Field{
		field.Int64("id").SchemaType(
			map[string]string{
				dialect.Postgres: string(schema_helper.INTEGER),
			}),
		field.String("category").SchemaType(
			map[string]string{
				dialect.Postgres: schema_helper.Varchar(50),
			}).Optional(),
		field.String("title").SchemaType(
			map[string]string{
				dialect.Postgres: schema_helper.Varchar(50),
			}).Optional(),
		 field.String("status").SchemaType(
			 map[string]string{
				 dialect.Postgres: schema_helper.Varchar(50),
			 }).Optional(),
		 field.Bool("is_deleted").Default(false),
		 field.Int32("todo_id").
			 SchemaType(
				 map[string]string{
					 dialect.Postgres: string(schema_helper.INTEGER),
				 }).Optional(),

	}
}

// Edges of the Form.
func (Form) Edges() []ent.Edge {
	return 	[]ent.Edge{
		edge.From("todo", Todo.Type).
			Ref("form").Unique().Field("todo_id"),
	}
}
func (Form) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "form"},
	}
}
