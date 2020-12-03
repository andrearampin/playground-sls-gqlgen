// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package resolvers

import "github.com/andrearampin/menyoo/internal/datastore"

type Resolver struct {
	Datastore datastore.Datastore
}
