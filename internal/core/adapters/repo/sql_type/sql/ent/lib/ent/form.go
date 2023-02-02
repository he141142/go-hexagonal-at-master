// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/todo"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Form is the model entity for the Form schema.
type Form struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// TodoID holds the value of the "todo_id" field.
	TodoID int32 `json:"todo_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FormQuery when eager-loading is set.
	Edges FormEdges `json:"edges"`
}

// FormEdges holds the relations/edges for other nodes in the graph.
type FormEdges struct {
	// Todo holds the value of the todo edge.
	Todo *Todo `json:"todo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodoOrErr returns the Todo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FormEdges) TodoOrErr() (*Todo, error) {
	if e.loadedTypes[0] {
		if e.Todo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Todo, nil
	}
	return nil, &NotLoadedError{edge: "todo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Form) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case form.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case form.FieldID, form.FieldTodoID:
			values[i] = new(sql.NullInt64)
		case form.FieldCategory, form.FieldTitle, form.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Form", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Form fields.
func (f *Form) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case form.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int64(value.Int64)
		case form.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				f.Category = value.String
			}
		case form.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				f.Title = value.String
			}
		case form.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = value.String
			}
		case form.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				f.IsDeleted = value.Bool
			}
		case form.FieldTodoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field todo_id", values[i])
			} else if value.Valid {
				f.TodoID = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryTodo queries the "todo" edge of the Form entity.
func (f *Form) QueryTodo() *TodoQuery {
	return NewFormClient(f.config).QueryTodo(f)
}

// Update returns a builder for updating this Form.
// Note that you need to call Form.Unwrap() before calling this method if this Form
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Form) Update() *FormUpdateOne {
	return NewFormClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Form entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Form) Unwrap() *Form {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Form is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Form) String() string {
	var builder strings.Builder
	builder.WriteString("Form(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("category=")
	builder.WriteString(f.Category)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(f.Title)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(f.Status)
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", f.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("todo_id=")
	builder.WriteString(fmt.Sprintf("%v", f.TodoID))
	builder.WriteByte(')')
	return builder.String()
}

// Forms is a parsable slice of Form.
type Forms []*Form

func (f Forms) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}