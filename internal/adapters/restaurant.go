package adapters

import (
	"github.com/andrearampin/menyoo/internal/datastore"
	"github.com/andrearampin/menyoo/internal/graph/model"
)

func RestaurantModelToGraph(datastore datastore.Restaurant) model.Restaurant {
	return model.Restaurant{
		ID:        datastore.ID,
		Name:      datastore.Name,
		UpdatedAt: string(datastore.UpdatedAt),
	}
}
