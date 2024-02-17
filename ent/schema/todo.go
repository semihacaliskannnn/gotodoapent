package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Comment("nesnemizin ismini gireceğimiz alan"),
		field.Text("description").
			Comment("yapacağımızı gireceğimiz yer"),
		field.Bool("status").
			Default(false).
			Comment("nesnenin durumu"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
