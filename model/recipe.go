package model

import "gopkg.in/mgo.v2/bson"

// Recipe represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Recipe struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Name         string        `bson:"name" json:"name"`
	CoverImage   string        `bson:"cover_image" json:"cover_image"`
	Instructions string        `bson:"description" json:"description"`
}
