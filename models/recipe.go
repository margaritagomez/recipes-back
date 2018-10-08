package models

import "gopkg.in/mgo.v2/bson"

// Recipe model
type Recipe struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Instructions string        `bson:"instructions" json:"instructions"`
	Image        string        `bson:"image" json:"image"`
}
