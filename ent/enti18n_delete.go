// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hiifong/i18n/ent/enti18n"
	"github.com/hiifong/i18n/ent/predicate"
)

// EntI18nDelete is the builder for deleting a EntI18n entity.
type EntI18nDelete struct {
	config
	hooks    []Hook
	mutation *EntI18nMutation
}

// Where appends a list predicates to the EntI18nDelete builder.
func (eid *EntI18nDelete) Where(ps ...predicate.EntI18n) *EntI18nDelete {
	eid.mutation.Where(ps...)
	return eid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eid *EntI18nDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eid.sqlExec, eid.mutation, eid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eid *EntI18nDelete) ExecX(ctx context.Context) int {
	n, err := eid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eid *EntI18nDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enti18n.Table, sqlgraph.NewFieldSpec(enti18n.FieldID, field.TypeInt64))
	if ps := eid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eid.mutation.done = true
	return affected, err
}

// EntI18nDeleteOne is the builder for deleting a single EntI18n entity.
type EntI18nDeleteOne struct {
	eid *EntI18nDelete
}

// Where appends a list predicates to the EntI18nDelete builder.
func (eido *EntI18nDeleteOne) Where(ps ...predicate.EntI18n) *EntI18nDeleteOne {
	eido.eid.mutation.Where(ps...)
	return eido
}

// Exec executes the deletion query.
func (eido *EntI18nDeleteOne) Exec(ctx context.Context) error {
	n, err := eido.eid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enti18n.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eido *EntI18nDeleteOne) ExecX(ctx context.Context) {
	if err := eido.Exec(ctx); err != nil {
		panic(err)
	}
}
