// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"github.com/andrearampin/menyoo/internal/graph/generated"
)

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
