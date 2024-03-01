// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/vulnequal"
)

// VulnEqualDelete is the builder for deleting a VulnEqual entity.
type VulnEqualDelete struct {
	config
	hooks    []Hook
	mutation *VulnEqualMutation
}

// Where appends a list predicates to the VulnEqualDelete builder.
func (ved *VulnEqualDelete) Where(ps ...predicate.VulnEqual) *VulnEqualDelete {
	ved.mutation.Where(ps...)
	return ved
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ved *VulnEqualDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ved.sqlExec, ved.mutation, ved.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ved *VulnEqualDelete) ExecX(ctx context.Context) int {
	n, err := ved.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ved *VulnEqualDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vulnequal.Table, sqlgraph.NewFieldSpec(vulnequal.FieldID, field.TypeUUID))
	if ps := ved.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ved.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ved.mutation.done = true
	return affected, err
}

// VulnEqualDeleteOne is the builder for deleting a single VulnEqual entity.
type VulnEqualDeleteOne struct {
	ved *VulnEqualDelete
}

// Where appends a list predicates to the VulnEqualDelete builder.
func (vedo *VulnEqualDeleteOne) Where(ps ...predicate.VulnEqual) *VulnEqualDeleteOne {
	vedo.ved.mutation.Where(ps...)
	return vedo
}

// Exec executes the deletion query.
func (vedo *VulnEqualDeleteOne) Exec(ctx context.Context) error {
	n, err := vedo.ved.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vulnequal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vedo *VulnEqualDeleteOne) ExecX(ctx context.Context) {
	if err := vedo.Exec(ctx); err != nil {
		panic(err)
	}
}
