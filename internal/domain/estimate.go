package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Estimate struct {
    ProductID         primitive.ObjectID  `bson:"_id" omitempty`
    ProductName       string              `bson:"product_name" omitempty`
    Price              float64            `bson:"price" omitempty`
    GeoLocation        *Location          `bson:"geolocation,omitempty"`
}

type Location struct {
	Longitude float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
}