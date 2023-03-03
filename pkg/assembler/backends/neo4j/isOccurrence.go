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

package neo4jBackend

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (c *neo4jClient) IsOccurrence(ctx context.Context, isOccurrenceSpec *model.IsOccurrenceSpec) ([]*model.IsOccurrence, error) {
	queryAll := false
	if isOccurrenceSpec.Package != nil && isOccurrenceSpec.Source != nil {
		return nil, gqlerror.Errorf("cannot specify both package and source for IsOccurrence")
	}

	if isOccurrenceSpec.Package == nil && isOccurrenceSpec.Source == nil {
		queryAll = true
	}

	session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	aggregateIsOccurrence := []*model.IsOccurrence{}

	if queryAll || isOccurrenceSpec.Package != nil {
		var sb strings.Builder
		var firstMatch bool = true
		queryValues := map[string]any{}

		returnValue := " RETURN type.type, namespace.namespace, name.name, version.version, version.subpath, " +
			"version.qualifier_list, isOccurrence, objArt.algorithm, objArt.digest"

		// query with pkgVersion
		query := "MATCH (root:Pkg)-[:PkgHasType]->(type:PkgType)-[:PkgHasNamespace]->(namespace:PkgNamespace)" +
			"-[:PkgHasName]->(name:PkgName)-[:PkgHasVersion]->(version:PkgVersion)" +
			"-[:subject]-(isOccurrence:IsOccurrence)-[:has_occurrence]-(objArt:Artifact)"
		sb.WriteString(query)

		setPkgMatchValues(&sb, isOccurrenceSpec.Package, false, &firstMatch, queryValues)
		setArtifactMatchValues(&sb, isOccurrenceSpec.Artifact, true, &firstMatch, queryValues)
		setIsOccurrenceValues(&sb, isOccurrenceSpec, &firstMatch, queryValues)
		sb.WriteString(returnValue)

		if isOccurrenceSpec.Package == nil || isOccurrenceSpec.Package != nil && isOccurrenceSpec.Package.Version == nil && isOccurrenceSpec.Package.Subpath == nil &&
			len(isOccurrenceSpec.Package.Qualifiers) == 0 && !*isOccurrenceSpec.Package.MatchOnlyEmptyQualifiers {

			sb.WriteString("\nUNION")
			// query without pkgVersion
			query = "\nMATCH (root:Pkg)-[:PkgHasType]->(type:PkgType)-[:PkgHasNamespace]->(namespace:PkgNamespace)" +
				"-[:PkgHasName]->(name:PkgName)" +
				"-[:subject]-(isOccurrence:IsOccurrence)-[:has_occurrence]-(objArt:Artifact)" +
				"\nWITH *, null AS version"
			sb.WriteString(query)

			firstMatch = true
			setPkgMatchValues(&sb, isOccurrenceSpec.Package, false, &firstMatch, queryValues)
			setArtifactMatchValues(&sb, isOccurrenceSpec.Artifact, true, &firstMatch, queryValues)
			setIsOccurrenceValues(&sb, isOccurrenceSpec, &firstMatch, queryValues)
			sb.WriteString(returnValue)
		}

		result, err := session.ReadTransaction(
			func(tx neo4j.Transaction) (interface{}, error) {

				result, err := tx.Run(sb.String(), queryValues)
				if err != nil {
					return nil, err
				}

				collectedIsOccurrence := []*model.IsOccurrence{}

				for result.Next() {
					pkgQualifiers := result.Record().Values[5]
					subPath := result.Record().Values[4]
					version := result.Record().Values[3]
					nameString := result.Record().Values[2].(string)
					namespaceString := result.Record().Values[1].(string)
					typeString := result.Record().Values[0].(string)

					pkg := generateModelPackage(typeString, namespaceString, nameString, version, subPath, pkgQualifiers)

					algorithm := result.Record().Values[7].(string)
					digest := result.Record().Values[8].(string)
					artifact := generateModelArtifact(algorithm, digest)

					isOccurrenceNode := dbtype.Node{}
					if result.Record().Values[6] != nil {
						isOccurrenceNode = result.Record().Values[6].(dbtype.Node)
					} else {
						return nil, gqlerror.Errorf("isOccurrence Node not found in neo4j")
					}

					isOccurrence := generateModelIsOccurrence(pkg, artifact, isOccurrenceNode.Props[justification].(string),
						isOccurrenceNode.Props[origin].(string), isOccurrenceNode.Props[collector].(string))

					collectedIsOccurrence = append(collectedIsOccurrence, isOccurrence)
				}
				if err = result.Err(); err != nil {
					return nil, err
				}

				return collectedIsOccurrence, nil
			})
		if err != nil {
			return nil, err
		}
		aggregateIsOccurrence = append(aggregateIsOccurrence, result.([]*model.IsOccurrence)...)
	}

	if queryAll || isOccurrenceSpec.Source != nil {
		var sb strings.Builder
		var firstMatch bool = true
		queryValues := map[string]any{}

		query := "MATCH (root:Src)-[:SrcHasType]->(type:SrcType)-[:SrcHasNamespace]->(namespace:SrcNamespace)" +
			"-[:SrcHasName]->(name:SrcName)-[:subject]-(isOccurrence:IsOccurrence)-[:has_occurrence]-(objArt:Artifact)"
		sb.WriteString(query)

		setSrcMatchValues(&sb, isOccurrenceSpec.Source, false, &firstMatch, queryValues)
		setArtifactMatchValues(&sb, isOccurrenceSpec.Artifact, true, &firstMatch, queryValues)
		setIsOccurrenceValues(&sb, isOccurrenceSpec, &firstMatch, queryValues)
		sb.WriteString(" RETURN type.type, namespace.namespace, name.name, name.tag, name.commit, isOccurrence, objArt.algorithm, objArt.digest")

		result, err := session.ReadTransaction(
			func(tx neo4j.Transaction) (interface{}, error) {

				result, err := tx.Run(sb.String(), queryValues)
				if err != nil {
					return nil, err
				}

				collectedIsOccurrence := []*model.IsOccurrence{}

				for result.Next() {
					tag := result.Record().Values[3]
					commit := result.Record().Values[4]
					nameStr := result.Record().Values[2].(string)
					namespaceStr := result.Record().Values[1].(string)
					srcType := result.Record().Values[0].(string)
					src := generateModelSource(srcType, namespaceStr, nameStr, commit, tag)

					algorithm := result.Record().Values[6].(string)
					digest := result.Record().Values[7].(string)
					artifact := generateModelArtifact(algorithm, digest)

					isOccurrenceNode := dbtype.Node{}
					if result.Record().Values[5] != nil {
						isOccurrenceNode = result.Record().Values[5].(dbtype.Node)
					} else {
						return nil, gqlerror.Errorf("isOccurrence Node not found in neo4j")
					}

					isOccurrence := generateModelIsOccurrence(src, artifact, isOccurrenceNode.Props[justification].(string),
						isOccurrenceNode.Props[origin].(string), isOccurrenceNode.Props[collector].(string))

					collectedIsOccurrence = append(collectedIsOccurrence, isOccurrence)
				}
				if err = result.Err(); err != nil {
					return nil, err
				}

				return collectedIsOccurrence, nil
			})
		if err != nil {
			return nil, err
		}
		aggregateIsOccurrence = append(aggregateIsOccurrence, result.([]*model.IsOccurrence)...)
	}
	return aggregateIsOccurrence, nil
}

func setIsOccurrenceValues(sb *strings.Builder, isOccurrenceSpec *model.IsOccurrenceSpec, firstMatch *bool, queryValues map[string]any) {
	if isOccurrenceSpec.Justification != nil {
		matchProperties(sb, *firstMatch, "isOccurrence", "justification", "$justification")
		*firstMatch = false
		queryValues["justification"] = isOccurrenceSpec.Justification
	}
	if isOccurrenceSpec.Origin != nil {
		matchProperties(sb, *firstMatch, "isOccurrence", "origin", "$origin")
		*firstMatch = false
		queryValues["origin"] = isOccurrenceSpec.Origin
	}
	if isOccurrenceSpec.Collector != nil {
		matchProperties(sb, *firstMatch, "isOccurrence", "collector", "$collector")
		*firstMatch = false
		queryValues["collector"] = isOccurrenceSpec.Collector
	}
}

func generateModelIsOccurrence(subject model.PkgSrcObject, artifact *model.Artifact, justification, origin, collector string) *model.IsOccurrence {
	isOccurrence := model.IsOccurrence{
		Subject:            subject,
		OccurrenceArtifact: artifact,
		Justification:      justification,
		Origin:             origin,
		Collector:          collector,
	}
	return &isOccurrence
}