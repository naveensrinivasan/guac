//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// CertifyScorecard holds the schema definition for the CertifyScorecard entity.
type CertifyScorecard struct {
	ent.Schema
}

// Fields of the CertifyScorecard.
func (CertifyScorecard) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(getUUIDv7).
			Unique().
			Immutable(),
		field.UUID("source_id", getUUIDv7()),
		field.JSON("checks", []*model.ScorecardCheck{}),
		field.Float("aggregate_score").Default(0).Comment("Overall Scorecard score for the source"),
		field.Time("time_scanned").Default(time.Now),
		field.String("scorecard_version"),
		field.String("scorecard_commit"),
		field.String("origin"),
		field.String("collector"),
		field.String("checks_hash").Comment("A SHA1 of the checks fields after sorting keys, used to ensure uniqueness of scorecard records."),
	}
}

// Edges of the CertifyScorecard.
func (CertifyScorecard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("source", SourceName.Type).Unique().Required().Field("source_id"),
	}
}
func (CertifyScorecard) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_id", "origin", "collector", "scorecard_version", "scorecard_commit", "aggregate_score", "time_scanned", "checks_hash").Unique(),
	}
}
