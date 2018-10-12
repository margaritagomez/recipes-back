package models

// Ingredient model
type Ingredient struct {
	Name     string  `bson:"name" json:"name"`
	Quantity float32 `bson:"quantity" json:"quantity"`
	Unit     string  `bson:"unit" json:"unit"`
}
