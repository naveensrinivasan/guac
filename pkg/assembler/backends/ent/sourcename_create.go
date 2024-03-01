// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// SourceNameCreate is the builder for creating a SourceName entity.
type SourceNameCreate struct {
	config
	mutation *SourceNameMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (snc *SourceNameCreate) SetType(s string) *SourceNameCreate {
	snc.mutation.SetType(s)
	return snc
}

// SetNamespace sets the "namespace" field.
func (snc *SourceNameCreate) SetNamespace(s string) *SourceNameCreate {
	snc.mutation.SetNamespace(s)
	return snc
}

// SetName sets the "name" field.
func (snc *SourceNameCreate) SetName(s string) *SourceNameCreate {
	snc.mutation.SetName(s)
	return snc
}

// SetCommit sets the "commit" field.
func (snc *SourceNameCreate) SetCommit(s string) *SourceNameCreate {
	snc.mutation.SetCommit(s)
	return snc
}

// SetNillableCommit sets the "commit" field if the given value is not nil.
func (snc *SourceNameCreate) SetNillableCommit(s *string) *SourceNameCreate {
	if s != nil {
		snc.SetCommit(*s)
	}
	return snc
}

// SetTag sets the "tag" field.
func (snc *SourceNameCreate) SetTag(s string) *SourceNameCreate {
	snc.mutation.SetTag(s)
	return snc
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (snc *SourceNameCreate) SetNillableTag(s *string) *SourceNameCreate {
	if s != nil {
		snc.SetTag(*s)
	}
	return snc
}

// SetID sets the "id" field.
func (snc *SourceNameCreate) SetID(u uuid.UUID) *SourceNameCreate {
	snc.mutation.SetID(u)
	return snc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (snc *SourceNameCreate) SetNillableID(u *uuid.UUID) *SourceNameCreate {
	if u != nil {
		snc.SetID(*u)
	}
	return snc
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (snc *SourceNameCreate) AddOccurrenceIDs(ids ...uuid.UUID) *SourceNameCreate {
	snc.mutation.AddOccurrenceIDs(ids...)
	return snc
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (snc *SourceNameCreate) AddOccurrences(o ...*Occurrence) *SourceNameCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return snc.AddOccurrenceIDs(ids...)
}

// Mutation returns the SourceNameMutation object of the builder.
func (snc *SourceNameCreate) Mutation() *SourceNameMutation {
	return snc.mutation
}

// Save creates the SourceName in the database.
func (snc *SourceNameCreate) Save(ctx context.Context) (*SourceName, error) {
	snc.defaults()
	return withHooks(ctx, snc.sqlSave, snc.mutation, snc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (snc *SourceNameCreate) SaveX(ctx context.Context) *SourceName {
	v, err := snc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (snc *SourceNameCreate) Exec(ctx context.Context) error {
	_, err := snc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snc *SourceNameCreate) ExecX(ctx context.Context) {
	if err := snc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snc *SourceNameCreate) defaults() {
	if _, ok := snc.mutation.ID(); !ok {
		v := sourcename.DefaultID()
		snc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snc *SourceNameCreate) check() error {
	if _, ok := snc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "SourceName.type"`)}
	}
	if _, ok := snc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "SourceName.namespace"`)}
	}
	if _, ok := snc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SourceName.name"`)}
	}
	return nil
}

func (snc *SourceNameCreate) sqlSave(ctx context.Context) (*SourceName, error) {
	if err := snc.check(); err != nil {
		return nil, err
	}
	_node, _spec := snc.createSpec()
	if err := sqlgraph.CreateNode(ctx, snc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	snc.mutation.id = &_node.ID
	snc.mutation.done = true
	return _node, nil
}

func (snc *SourceNameCreate) createSpec() (*SourceName, *sqlgraph.CreateSpec) {
	var (
		_node = &SourceName{config: snc.config}
		_spec = sqlgraph.NewCreateSpec(sourcename.Table, sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = snc.conflict
	if id, ok := snc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := snc.mutation.GetType(); ok {
		_spec.SetField(sourcename.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := snc.mutation.Namespace(); ok {
		_spec.SetField(sourcename.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := snc.mutation.Name(); ok {
		_spec.SetField(sourcename.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := snc.mutation.Commit(); ok {
		_spec.SetField(sourcename.FieldCommit, field.TypeString, value)
		_node.Commit = value
	}
	if value, ok := snc.mutation.Tag(); ok {
		_spec.SetField(sourcename.FieldTag, field.TypeString, value)
		_node.Tag = value
	}
	if nodes := snc.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcename.OccurrencesTable,
			Columns: []string{sourcename.OccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SourceName.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SourceNameUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (snc *SourceNameCreate) OnConflict(opts ...sql.ConflictOption) *SourceNameUpsertOne {
	snc.conflict = opts
	return &SourceNameUpsertOne{
		create: snc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SourceName.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (snc *SourceNameCreate) OnConflictColumns(columns ...string) *SourceNameUpsertOne {
	snc.conflict = append(snc.conflict, sql.ConflictColumns(columns...))
	return &SourceNameUpsertOne{
		create: snc,
	}
}

type (
	// SourceNameUpsertOne is the builder for "upsert"-ing
	//  one SourceName node.
	SourceNameUpsertOne struct {
		create *SourceNameCreate
	}

	// SourceNameUpsert is the "OnConflict" setter.
	SourceNameUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *SourceNameUpsert) SetType(v string) *SourceNameUpsert {
	u.Set(sourcename.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SourceNameUpsert) UpdateType() *SourceNameUpsert {
	u.SetExcluded(sourcename.FieldType)
	return u
}

// SetNamespace sets the "namespace" field.
func (u *SourceNameUpsert) SetNamespace(v string) *SourceNameUpsert {
	u.Set(sourcename.FieldNamespace, v)
	return u
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *SourceNameUpsert) UpdateNamespace() *SourceNameUpsert {
	u.SetExcluded(sourcename.FieldNamespace)
	return u
}

// SetName sets the "name" field.
func (u *SourceNameUpsert) SetName(v string) *SourceNameUpsert {
	u.Set(sourcename.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SourceNameUpsert) UpdateName() *SourceNameUpsert {
	u.SetExcluded(sourcename.FieldName)
	return u
}

// SetCommit sets the "commit" field.
func (u *SourceNameUpsert) SetCommit(v string) *SourceNameUpsert {
	u.Set(sourcename.FieldCommit, v)
	return u
}

// UpdateCommit sets the "commit" field to the value that was provided on create.
func (u *SourceNameUpsert) UpdateCommit() *SourceNameUpsert {
	u.SetExcluded(sourcename.FieldCommit)
	return u
}

// ClearCommit clears the value of the "commit" field.
func (u *SourceNameUpsert) ClearCommit() *SourceNameUpsert {
	u.SetNull(sourcename.FieldCommit)
	return u
}

// SetTag sets the "tag" field.
func (u *SourceNameUpsert) SetTag(v string) *SourceNameUpsert {
	u.Set(sourcename.FieldTag, v)
	return u
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *SourceNameUpsert) UpdateTag() *SourceNameUpsert {
	u.SetExcluded(sourcename.FieldTag)
	return u
}

// ClearTag clears the value of the "tag" field.
func (u *SourceNameUpsert) ClearTag() *SourceNameUpsert {
	u.SetNull(sourcename.FieldTag)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SourceName.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sourcename.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SourceNameUpsertOne) UpdateNewValues() *SourceNameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sourcename.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SourceName.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SourceNameUpsertOne) Ignore() *SourceNameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SourceNameUpsertOne) DoNothing() *SourceNameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SourceNameCreate.OnConflict
// documentation for more info.
func (u *SourceNameUpsertOne) Update(set func(*SourceNameUpsert)) *SourceNameUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SourceNameUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *SourceNameUpsertOne) SetType(v string) *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SourceNameUpsertOne) UpdateType() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateType()
	})
}

// SetNamespace sets the "namespace" field.
func (u *SourceNameUpsertOne) SetNamespace(v string) *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *SourceNameUpsertOne) UpdateNamespace() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateNamespace()
	})
}

// SetName sets the "name" field.
func (u *SourceNameUpsertOne) SetName(v string) *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SourceNameUpsertOne) UpdateName() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateName()
	})
}

// SetCommit sets the "commit" field.
func (u *SourceNameUpsertOne) SetCommit(v string) *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetCommit(v)
	})
}

// UpdateCommit sets the "commit" field to the value that was provided on create.
func (u *SourceNameUpsertOne) UpdateCommit() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateCommit()
	})
}

// ClearCommit clears the value of the "commit" field.
func (u *SourceNameUpsertOne) ClearCommit() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.ClearCommit()
	})
}

// SetTag sets the "tag" field.
func (u *SourceNameUpsertOne) SetTag(v string) *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *SourceNameUpsertOne) UpdateTag() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateTag()
	})
}

// ClearTag clears the value of the "tag" field.
func (u *SourceNameUpsertOne) ClearTag() *SourceNameUpsertOne {
	return u.Update(func(s *SourceNameUpsert) {
		s.ClearTag()
	})
}

// Exec executes the query.
func (u *SourceNameUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SourceNameCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SourceNameUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SourceNameUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SourceNameUpsertOne.ID is not supported by MySQL driver. Use SourceNameUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SourceNameUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SourceNameCreateBulk is the builder for creating many SourceName entities in bulk.
type SourceNameCreateBulk struct {
	config
	err      error
	builders []*SourceNameCreate
	conflict []sql.ConflictOption
}

// Save creates the SourceName entities in the database.
func (sncb *SourceNameCreateBulk) Save(ctx context.Context) ([]*SourceName, error) {
	if sncb.err != nil {
		return nil, sncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sncb.builders))
	nodes := make([]*SourceName, len(sncb.builders))
	mutators := make([]Mutator, len(sncb.builders))
	for i := range sncb.builders {
		func(i int, root context.Context) {
			builder := sncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SourceNameMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sncb *SourceNameCreateBulk) SaveX(ctx context.Context) []*SourceName {
	v, err := sncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sncb *SourceNameCreateBulk) Exec(ctx context.Context) error {
	_, err := sncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sncb *SourceNameCreateBulk) ExecX(ctx context.Context) {
	if err := sncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SourceName.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SourceNameUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (sncb *SourceNameCreateBulk) OnConflict(opts ...sql.ConflictOption) *SourceNameUpsertBulk {
	sncb.conflict = opts
	return &SourceNameUpsertBulk{
		create: sncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SourceName.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sncb *SourceNameCreateBulk) OnConflictColumns(columns ...string) *SourceNameUpsertBulk {
	sncb.conflict = append(sncb.conflict, sql.ConflictColumns(columns...))
	return &SourceNameUpsertBulk{
		create: sncb,
	}
}

// SourceNameUpsertBulk is the builder for "upsert"-ing
// a bulk of SourceName nodes.
type SourceNameUpsertBulk struct {
	create *SourceNameCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SourceName.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sourcename.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SourceNameUpsertBulk) UpdateNewValues() *SourceNameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sourcename.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SourceName.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SourceNameUpsertBulk) Ignore() *SourceNameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SourceNameUpsertBulk) DoNothing() *SourceNameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SourceNameCreateBulk.OnConflict
// documentation for more info.
func (u *SourceNameUpsertBulk) Update(set func(*SourceNameUpsert)) *SourceNameUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SourceNameUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *SourceNameUpsertBulk) SetType(v string) *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *SourceNameUpsertBulk) UpdateType() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateType()
	})
}

// SetNamespace sets the "namespace" field.
func (u *SourceNameUpsertBulk) SetNamespace(v string) *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *SourceNameUpsertBulk) UpdateNamespace() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateNamespace()
	})
}

// SetName sets the "name" field.
func (u *SourceNameUpsertBulk) SetName(v string) *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SourceNameUpsertBulk) UpdateName() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateName()
	})
}

// SetCommit sets the "commit" field.
func (u *SourceNameUpsertBulk) SetCommit(v string) *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetCommit(v)
	})
}

// UpdateCommit sets the "commit" field to the value that was provided on create.
func (u *SourceNameUpsertBulk) UpdateCommit() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateCommit()
	})
}

// ClearCommit clears the value of the "commit" field.
func (u *SourceNameUpsertBulk) ClearCommit() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.ClearCommit()
	})
}

// SetTag sets the "tag" field.
func (u *SourceNameUpsertBulk) SetTag(v string) *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *SourceNameUpsertBulk) UpdateTag() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.UpdateTag()
	})
}

// ClearTag clears the value of the "tag" field.
func (u *SourceNameUpsertBulk) ClearTag() *SourceNameUpsertBulk {
	return u.Update(func(s *SourceNameUpsert) {
		s.ClearTag()
	})
}

// Exec executes the query.
func (u *SourceNameUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SourceNameCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SourceNameCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SourceNameUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
