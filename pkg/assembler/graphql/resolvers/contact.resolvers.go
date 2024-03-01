package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestPointOfContact is the resolver for the ingestPointOfContact field.
func (r *mutationResolver) IngestPointOfContact(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType model.MatchFlags, pointOfContact model.PointOfContactInputSpec) (string, error) {
	funcName := "IngestPointOfContact"
	if err := validatePackageSourceOrArtifactInput(&subject, funcName); err != nil {
		return "", gqlerror.Errorf("%v :: %s", funcName, err)
	}
	return r.Backend.IngestPointOfContact(ctx, subject, &pkgMatchType, pointOfContact)
}

// IngestPointOfContacts is the resolver for the ingestPointOfContacts field.
func (r *mutationResolver) IngestPointOfContacts(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType model.MatchFlags, pointOfContacts []*model.PointOfContactInputSpec) ([]string, error) {
	funcName := "IngestPointOfContacts"
	valuesDefined := 0
	if len(subjects.Packages) > 0 {
		if len(subjects.Packages) != len(pointOfContacts) {
			return []string{}, gqlerror.Errorf("%v :: uneven packages and pointOfContacts for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Artifacts) > 0 {
		if len(subjects.Artifacts) != len(pointOfContacts) {
			return []string{}, gqlerror.Errorf("%v :: uneven artifacts and pointOfContacts for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Sources) > 0 {
		if len(subjects.Sources) != len(pointOfContacts) {
			return []string{}, gqlerror.Errorf("%v :: uneven sources and pointOfContacts for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if valuesDefined != 1 {
		return []string{}, gqlerror.Errorf("%v :: must specify at most packages, artifacts or sources", funcName)
	}
	return r.Backend.IngestPointOfContacts(ctx, subjects, &pkgMatchType, pointOfContacts)
}

// PointOfContact is the resolver for the PointOfContact field.
func (r *queryResolver) PointOfContact(ctx context.Context, pointOfContactSpec model.PointOfContactSpec) ([]*model.PointOfContact, error) {
	if err := validatePackageSourceOrArtifactQueryFilter(pointOfContactSpec.Subject); err != nil {
		return nil, gqlerror.Errorf("PointOfContact :: %s", err)
	}
	return r.Backend.PointOfContact(ctx, &pointOfContactSpec)
}
