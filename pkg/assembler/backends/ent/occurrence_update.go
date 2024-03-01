// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// OccurrenceUpdate is the builder for updating Occurrence entities.
type OccurrenceUpdate struct {
	config
	hooks    []Hook
	mutation *OccurrenceMutation
}

// Where appends a list predicates to the OccurrenceUpdate builder.
func (ou *OccurrenceUpdate) Where(ps ...predicate.Occurrence) *OccurrenceUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetArtifactID sets the "artifact_id" field.
func (ou *OccurrenceUpdate) SetArtifactID(u uuid.UUID) *OccurrenceUpdate {
	ou.mutation.SetArtifactID(u)
	return ou
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillableArtifactID(u *uuid.UUID) *OccurrenceUpdate {
	if u != nil {
		ou.SetArtifactID(*u)
	}
	return ou
}

// SetJustification sets the "justification" field.
func (ou *OccurrenceUpdate) SetJustification(s string) *OccurrenceUpdate {
	ou.mutation.SetJustification(s)
	return ou
}

// SetNillableJustification sets the "justification" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillableJustification(s *string) *OccurrenceUpdate {
	if s != nil {
		ou.SetJustification(*s)
	}
	return ou
}

// SetOrigin sets the "origin" field.
func (ou *OccurrenceUpdate) SetOrigin(s string) *OccurrenceUpdate {
	ou.mutation.SetOrigin(s)
	return ou
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillableOrigin(s *string) *OccurrenceUpdate {
	if s != nil {
		ou.SetOrigin(*s)
	}
	return ou
}

// SetCollector sets the "collector" field.
func (ou *OccurrenceUpdate) SetCollector(s string) *OccurrenceUpdate {
	ou.mutation.SetCollector(s)
	return ou
}

// SetNillableCollector sets the "collector" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillableCollector(s *string) *OccurrenceUpdate {
	if s != nil {
		ou.SetCollector(*s)
	}
	return ou
}

// SetSourceID sets the "source_id" field.
func (ou *OccurrenceUpdate) SetSourceID(u uuid.UUID) *OccurrenceUpdate {
	ou.mutation.SetSourceID(u)
	return ou
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillableSourceID(u *uuid.UUID) *OccurrenceUpdate {
	if u != nil {
		ou.SetSourceID(*u)
	}
	return ou
}

// ClearSourceID clears the value of the "source_id" field.
func (ou *OccurrenceUpdate) ClearSourceID() *OccurrenceUpdate {
	ou.mutation.ClearSourceID()
	return ou
}

// SetPackageID sets the "package_id" field.
func (ou *OccurrenceUpdate) SetPackageID(u uuid.UUID) *OccurrenceUpdate {
	ou.mutation.SetPackageID(u)
	return ou
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (ou *OccurrenceUpdate) SetNillablePackageID(u *uuid.UUID) *OccurrenceUpdate {
	if u != nil {
		ou.SetPackageID(*u)
	}
	return ou
}

// ClearPackageID clears the value of the "package_id" field.
func (ou *OccurrenceUpdate) ClearPackageID() *OccurrenceUpdate {
	ou.mutation.ClearPackageID()
	return ou
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (ou *OccurrenceUpdate) SetArtifact(a *Artifact) *OccurrenceUpdate {
	return ou.SetArtifactID(a.ID)
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (ou *OccurrenceUpdate) SetPackage(p *PackageVersion) *OccurrenceUpdate {
	return ou.SetPackageID(p.ID)
}

// SetSource sets the "source" edge to the SourceName entity.
func (ou *OccurrenceUpdate) SetSource(s *SourceName) *OccurrenceUpdate {
	return ou.SetSourceID(s.ID)
}

// AddIncludedInSbomIDs adds the "included_in_sboms" edge to the BillOfMaterials entity by IDs.
func (ou *OccurrenceUpdate) AddIncludedInSbomIDs(ids ...uuid.UUID) *OccurrenceUpdate {
	ou.mutation.AddIncludedInSbomIDs(ids...)
	return ou
}

// AddIncludedInSboms adds the "included_in_sboms" edges to the BillOfMaterials entity.
func (ou *OccurrenceUpdate) AddIncludedInSboms(b ...*BillOfMaterials) *OccurrenceUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ou.AddIncludedInSbomIDs(ids...)
}

// Mutation returns the OccurrenceMutation object of the builder.
func (ou *OccurrenceUpdate) Mutation() *OccurrenceMutation {
	return ou.mutation
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (ou *OccurrenceUpdate) ClearArtifact() *OccurrenceUpdate {
	ou.mutation.ClearArtifact()
	return ou
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (ou *OccurrenceUpdate) ClearPackage() *OccurrenceUpdate {
	ou.mutation.ClearPackage()
	return ou
}

// ClearSource clears the "source" edge to the SourceName entity.
func (ou *OccurrenceUpdate) ClearSource() *OccurrenceUpdate {
	ou.mutation.ClearSource()
	return ou
}

// ClearIncludedInSboms clears all "included_in_sboms" edges to the BillOfMaterials entity.
func (ou *OccurrenceUpdate) ClearIncludedInSboms() *OccurrenceUpdate {
	ou.mutation.ClearIncludedInSboms()
	return ou
}

// RemoveIncludedInSbomIDs removes the "included_in_sboms" edge to BillOfMaterials entities by IDs.
func (ou *OccurrenceUpdate) RemoveIncludedInSbomIDs(ids ...uuid.UUID) *OccurrenceUpdate {
	ou.mutation.RemoveIncludedInSbomIDs(ids...)
	return ou
}

// RemoveIncludedInSboms removes "included_in_sboms" edges to BillOfMaterials entities.
func (ou *OccurrenceUpdate) RemoveIncludedInSboms(b ...*BillOfMaterials) *OccurrenceUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ou.RemoveIncludedInSbomIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OccurrenceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OccurrenceUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OccurrenceUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OccurrenceUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OccurrenceUpdate) check() error {
	if _, ok := ou.mutation.ArtifactID(); ou.mutation.ArtifactCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Occurrence.artifact"`)
	}
	return nil
}

func (ou *OccurrenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(occurrence.Table, occurrence.Columns, sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Justification(); ok {
		_spec.SetField(occurrence.FieldJustification, field.TypeString, value)
	}
	if value, ok := ou.mutation.Origin(); ok {
		_spec.SetField(occurrence.FieldOrigin, field.TypeString, value)
	}
	if value, ok := ou.mutation.Collector(); ok {
		_spec.SetField(occurrence.FieldCollector, field.TypeString, value)
	}
	if ou.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.ArtifactTable,
			Columns: []string{occurrence.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.ArtifactTable,
			Columns: []string{occurrence.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.PackageTable,
			Columns: []string{occurrence.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.PackageTable,
			Columns: []string{occurrence.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.SourceTable,
			Columns: []string{occurrence.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.SourceTable,
			Columns: []string{occurrence.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.IncludedInSbomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedIncludedInSbomsIDs(); len(nodes) > 0 && !ou.mutation.IncludedInSbomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.IncludedInSbomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occurrence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OccurrenceUpdateOne is the builder for updating a single Occurrence entity.
type OccurrenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OccurrenceMutation
}

// SetArtifactID sets the "artifact_id" field.
func (ouo *OccurrenceUpdateOne) SetArtifactID(u uuid.UUID) *OccurrenceUpdateOne {
	ouo.mutation.SetArtifactID(u)
	return ouo
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillableArtifactID(u *uuid.UUID) *OccurrenceUpdateOne {
	if u != nil {
		ouo.SetArtifactID(*u)
	}
	return ouo
}

// SetJustification sets the "justification" field.
func (ouo *OccurrenceUpdateOne) SetJustification(s string) *OccurrenceUpdateOne {
	ouo.mutation.SetJustification(s)
	return ouo
}

// SetNillableJustification sets the "justification" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillableJustification(s *string) *OccurrenceUpdateOne {
	if s != nil {
		ouo.SetJustification(*s)
	}
	return ouo
}

// SetOrigin sets the "origin" field.
func (ouo *OccurrenceUpdateOne) SetOrigin(s string) *OccurrenceUpdateOne {
	ouo.mutation.SetOrigin(s)
	return ouo
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillableOrigin(s *string) *OccurrenceUpdateOne {
	if s != nil {
		ouo.SetOrigin(*s)
	}
	return ouo
}

// SetCollector sets the "collector" field.
func (ouo *OccurrenceUpdateOne) SetCollector(s string) *OccurrenceUpdateOne {
	ouo.mutation.SetCollector(s)
	return ouo
}

// SetNillableCollector sets the "collector" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillableCollector(s *string) *OccurrenceUpdateOne {
	if s != nil {
		ouo.SetCollector(*s)
	}
	return ouo
}

// SetSourceID sets the "source_id" field.
func (ouo *OccurrenceUpdateOne) SetSourceID(u uuid.UUID) *OccurrenceUpdateOne {
	ouo.mutation.SetSourceID(u)
	return ouo
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillableSourceID(u *uuid.UUID) *OccurrenceUpdateOne {
	if u != nil {
		ouo.SetSourceID(*u)
	}
	return ouo
}

// ClearSourceID clears the value of the "source_id" field.
func (ouo *OccurrenceUpdateOne) ClearSourceID() *OccurrenceUpdateOne {
	ouo.mutation.ClearSourceID()
	return ouo
}

// SetPackageID sets the "package_id" field.
func (ouo *OccurrenceUpdateOne) SetPackageID(u uuid.UUID) *OccurrenceUpdateOne {
	ouo.mutation.SetPackageID(u)
	return ouo
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (ouo *OccurrenceUpdateOne) SetNillablePackageID(u *uuid.UUID) *OccurrenceUpdateOne {
	if u != nil {
		ouo.SetPackageID(*u)
	}
	return ouo
}

// ClearPackageID clears the value of the "package_id" field.
func (ouo *OccurrenceUpdateOne) ClearPackageID() *OccurrenceUpdateOne {
	ouo.mutation.ClearPackageID()
	return ouo
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (ouo *OccurrenceUpdateOne) SetArtifact(a *Artifact) *OccurrenceUpdateOne {
	return ouo.SetArtifactID(a.ID)
}

// SetPackage sets the "package" edge to the PackageVersion entity.
func (ouo *OccurrenceUpdateOne) SetPackage(p *PackageVersion) *OccurrenceUpdateOne {
	return ouo.SetPackageID(p.ID)
}

// SetSource sets the "source" edge to the SourceName entity.
func (ouo *OccurrenceUpdateOne) SetSource(s *SourceName) *OccurrenceUpdateOne {
	return ouo.SetSourceID(s.ID)
}

// AddIncludedInSbomIDs adds the "included_in_sboms" edge to the BillOfMaterials entity by IDs.
func (ouo *OccurrenceUpdateOne) AddIncludedInSbomIDs(ids ...uuid.UUID) *OccurrenceUpdateOne {
	ouo.mutation.AddIncludedInSbomIDs(ids...)
	return ouo
}

// AddIncludedInSboms adds the "included_in_sboms" edges to the BillOfMaterials entity.
func (ouo *OccurrenceUpdateOne) AddIncludedInSboms(b ...*BillOfMaterials) *OccurrenceUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ouo.AddIncludedInSbomIDs(ids...)
}

// Mutation returns the OccurrenceMutation object of the builder.
func (ouo *OccurrenceUpdateOne) Mutation() *OccurrenceMutation {
	return ouo.mutation
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (ouo *OccurrenceUpdateOne) ClearArtifact() *OccurrenceUpdateOne {
	ouo.mutation.ClearArtifact()
	return ouo
}

// ClearPackage clears the "package" edge to the PackageVersion entity.
func (ouo *OccurrenceUpdateOne) ClearPackage() *OccurrenceUpdateOne {
	ouo.mutation.ClearPackage()
	return ouo
}

// ClearSource clears the "source" edge to the SourceName entity.
func (ouo *OccurrenceUpdateOne) ClearSource() *OccurrenceUpdateOne {
	ouo.mutation.ClearSource()
	return ouo
}

// ClearIncludedInSboms clears all "included_in_sboms" edges to the BillOfMaterials entity.
func (ouo *OccurrenceUpdateOne) ClearIncludedInSboms() *OccurrenceUpdateOne {
	ouo.mutation.ClearIncludedInSboms()
	return ouo
}

// RemoveIncludedInSbomIDs removes the "included_in_sboms" edge to BillOfMaterials entities by IDs.
func (ouo *OccurrenceUpdateOne) RemoveIncludedInSbomIDs(ids ...uuid.UUID) *OccurrenceUpdateOne {
	ouo.mutation.RemoveIncludedInSbomIDs(ids...)
	return ouo
}

// RemoveIncludedInSboms removes "included_in_sboms" edges to BillOfMaterials entities.
func (ouo *OccurrenceUpdateOne) RemoveIncludedInSboms(b ...*BillOfMaterials) *OccurrenceUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ouo.RemoveIncludedInSbomIDs(ids...)
}

// Where appends a list predicates to the OccurrenceUpdate builder.
func (ouo *OccurrenceUpdateOne) Where(ps ...predicate.Occurrence) *OccurrenceUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OccurrenceUpdateOne) Select(field string, fields ...string) *OccurrenceUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Occurrence entity.
func (ouo *OccurrenceUpdateOne) Save(ctx context.Context) (*Occurrence, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OccurrenceUpdateOne) SaveX(ctx context.Context) *Occurrence {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OccurrenceUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OccurrenceUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OccurrenceUpdateOne) check() error {
	if _, ok := ouo.mutation.ArtifactID(); ouo.mutation.ArtifactCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Occurrence.artifact"`)
	}
	return nil
}

func (ouo *OccurrenceUpdateOne) sqlSave(ctx context.Context) (_node *Occurrence, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(occurrence.Table, occurrence.Columns, sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Occurrence.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, occurrence.FieldID)
		for _, f := range fields {
			if !occurrence.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != occurrence.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Justification(); ok {
		_spec.SetField(occurrence.FieldJustification, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Origin(); ok {
		_spec.SetField(occurrence.FieldOrigin, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Collector(); ok {
		_spec.SetField(occurrence.FieldCollector, field.TypeString, value)
	}
	if ouo.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.ArtifactTable,
			Columns: []string{occurrence.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.ArtifactTable,
			Columns: []string{occurrence.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.PackageTable,
			Columns: []string{occurrence.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.PackageTable,
			Columns: []string{occurrence.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.SourceTable,
			Columns: []string{occurrence.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   occurrence.SourceTable,
			Columns: []string{occurrence.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.IncludedInSbomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedIncludedInSbomsIDs(); len(nodes) > 0 && !ouo.mutation.IncludedInSbomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.IncludedInSbomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   occurrence.IncludedInSbomsTable,
			Columns: occurrence.IncludedInSbomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Occurrence{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occurrence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
