// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mall/ent/psstrategy"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PsStrategy is the model entity for the PsStrategy schema.
type PsStrategy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner int `json:"owner,omitempty"`
	// ScriptContent holds the value of the "script_content" field.
	ScriptContent string `json:"script_content,omitempty"`
	// IsDelete holds the value of the "is_delete" field.
	IsDelete     int `json:"is_delete,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PsStrategy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case psstrategy.FieldID, psstrategy.FieldOwner, psstrategy.FieldIsDelete:
			values[i] = new(sql.NullInt64)
		case psstrategy.FieldScriptContent:
			values[i] = new(sql.NullString)
		case psstrategy.FieldCreateTime, psstrategy.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PsStrategy fields.
func (ps *PsStrategy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case psstrategy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ps.ID = int(value.Int64)
		case psstrategy.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ps.CreateTime = value.Time
			}
		case psstrategy.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ps.UpdateTime = value.Time
			}
		case psstrategy.FieldOwner:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				ps.Owner = int(value.Int64)
			}
		case psstrategy.FieldScriptContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field script_content", values[i])
			} else if value.Valid {
				ps.ScriptContent = value.String
			}
		case psstrategy.FieldIsDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_delete", values[i])
			} else if value.Valid {
				ps.IsDelete = int(value.Int64)
			}
		default:
			ps.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PsStrategy.
// This includes values selected through modifiers, order, etc.
func (ps *PsStrategy) Value(name string) (ent.Value, error) {
	return ps.selectValues.Get(name)
}

// Update returns a builder for updating this PsStrategy.
// Note that you need to call PsStrategy.Unwrap() before calling this method if this PsStrategy
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *PsStrategy) Update() *PsStrategyUpdateOne {
	return NewPsStrategyClient(ps.config).UpdateOne(ps)
}

// Unwrap unwraps the PsStrategy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *PsStrategy) Unwrap() *PsStrategy {
	_tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("ent: PsStrategy is not a transactional entity")
	}
	ps.config.driver = _tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *PsStrategy) String() string {
	var builder strings.Builder
	builder.WriteString("PsStrategy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ps.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ps.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ps.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner=")
	builder.WriteString(fmt.Sprintf("%v", ps.Owner))
	builder.WriteString(", ")
	builder.WriteString("script_content=")
	builder.WriteString(ps.ScriptContent)
	builder.WriteString(", ")
	builder.WriteString("is_delete=")
	builder.WriteString(fmt.Sprintf("%v", ps.IsDelete))
	builder.WriteByte(')')
	return builder.String()
}

// PsStrategies is a parsable slice of PsStrategy.
type PsStrategies []*PsStrategy
