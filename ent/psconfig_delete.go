// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"mall/ent/predicate"
	"mall/ent/psconfig"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PsConfigDelete is the builder for deleting a PsConfig entity.
type PsConfigDelete struct {
	config
	hooks    []Hook
	mutation *PsConfigMutation
}

// Where appends a list predicates to the PsConfigDelete builder.
func (pcd *PsConfigDelete) Where(ps ...predicate.PsConfig) *PsConfigDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *PsConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pcd.sqlExec, pcd.mutation, pcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *PsConfigDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *PsConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(psconfig.Table, sqlgraph.NewFieldSpec(psconfig.FieldID, field.TypeInt))
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pcd.mutation.done = true
	return affected, err
}

// PsConfigDeleteOne is the builder for deleting a single PsConfig entity.
type PsConfigDeleteOne struct {
	pcd *PsConfigDelete
}

// Where appends a list predicates to the PsConfigDelete builder.
func (pcdo *PsConfigDeleteOne) Where(ps ...predicate.PsConfig) *PsConfigDeleteOne {
	pcdo.pcd.mutation.Where(ps...)
	return pcdo
}

// Exec executes the deletion query.
func (pcdo *PsConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{psconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *PsConfigDeleteOne) ExecX(ctx context.Context) {
	if err := pcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
