package datastore

import (
	"time"

	"github.com/google/uuid"
)

const (
	_restaurantTableName = "restaurants"
)

// Restaurant represents the restaurant entity.
type Restaurant struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

// TableName return the name of the table for the restaurant records.
func (r Restaurant) TableName() string {
	return _restaurantTableName
}

// CreateRestaurant creates a new Restaurant record.
func (d DB) CreateRestaurant(name string) (Restaurant, error) {
	r := Restaurant{
		ID:        uuid.New().String(),
		Name:      name,
		UpdatedAt: time.Now().UTC().Unix(),
	}

	return r, d.putItem(r)
}

// FetchRestaurant return a restaurant given the restaurant UUID.
func (d DB) FetchRestaurant(id string) (Restaurant, error) {
	var r Restaurant
	out, err := d.getItem(Restaurant{
		ID: id,
	})

	if err != nil {
		return r, err
	}

	return r, itemOutputStruct(*out, &r)
}

// UpdateRestaurant simply replace the current restaurant record with what passed.
func (d DB) UpdateRestaurant(r Restaurant) (Restaurant, error) {
	return r, d.putItem(r)
}
