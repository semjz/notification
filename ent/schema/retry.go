package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Message holds the schema definition for the Message entity.
type Retry struct {
	ent.Schema
}

// Fields of the Message.
func (Retry) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("message_uuid", uuid.UUID{}).Unique(),
		field.Enum("status").
			Values("pending", "failed").
			Default("pending"),
		field.Int("attempts").Default(1),
		field.Time("next_retry_at").Nillable().Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Retry.
func (Retry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).
			Ref("retry").
			Unique().
			Field("message_uuid").Required(),
	}
}
