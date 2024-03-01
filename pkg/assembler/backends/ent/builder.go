// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/builder"
)

// Builder is the model entity for the Builder schema.
type Builder struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// The URI of the builder, used as a unique identifier in the graph query
	URI string `json:"uri,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BuilderQuery when eager-loading is set.
	Edges        BuilderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BuilderEdges holds the relations/edges for other nodes in the graph.
type BuilderEdges struct {
	// SlsaAttestations holds the value of the slsa_attestations edge.
	SlsaAttestations []*SLSAAttestation `json:"slsa_attestations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedSlsaAttestations map[string][]*SLSAAttestation
}

// SlsaAttestationsOrErr returns the SlsaAttestations value or an error if the edge
// was not loaded in eager-loading.
func (e BuilderEdges) SlsaAttestationsOrErr() ([]*SLSAAttestation, error) {
	if e.loadedTypes[0] {
		return e.SlsaAttestations, nil
	}
	return nil, &NotLoadedError{edge: "slsa_attestations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Builder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case builder.FieldURI:
			values[i] = new(sql.NullString)
		case builder.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Builder fields.
func (b *Builder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case builder.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case builder.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				b.URI = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Builder.
// This includes values selected through modifiers, order, etc.
func (b *Builder) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QuerySlsaAttestations queries the "slsa_attestations" edge of the Builder entity.
func (b *Builder) QuerySlsaAttestations() *SLSAAttestationQuery {
	return NewBuilderClient(b.config).QuerySlsaAttestations(b)
}

// Update returns a builder for updating this Builder.
// Note that you need to call Builder.Unwrap() before calling this method if this Builder
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Builder) Update() *BuilderUpdateOne {
	return NewBuilderClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Builder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Builder) Unwrap() *Builder {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Builder is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Builder) String() string {
	var builder strings.Builder
	builder.WriteString("Builder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("uri=")
	builder.WriteString(b.URI)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSlsaAttestations returns the SlsaAttestations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (b *Builder) NamedSlsaAttestations(name string) ([]*SLSAAttestation, error) {
	if b.Edges.namedSlsaAttestations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := b.Edges.namedSlsaAttestations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (b *Builder) appendNamedSlsaAttestations(name string, edges ...*SLSAAttestation) {
	if b.Edges.namedSlsaAttestations == nil {
		b.Edges.namedSlsaAttestations = make(map[string][]*SLSAAttestation)
	}
	if len(edges) == 0 {
		b.Edges.namedSlsaAttestations[name] = []*SLSAAttestation{}
	} else {
		b.Edges.namedSlsaAttestations[name] = append(b.Edges.namedSlsaAttestations[name], edges...)
	}
}

// Builders is a parsable slice of Builder.
type Builders []*Builder
