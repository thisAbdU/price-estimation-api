package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Estimate struct {
    ProductID         primitive.ObjectID  `bson:"_id",omitempty`
	ProductName     string              `bson:"product_name",omitempty`
	Price 		    float64            `bson:"price",omitempty`
	Location 	  	string             `bson:"location",omitempty`
}

