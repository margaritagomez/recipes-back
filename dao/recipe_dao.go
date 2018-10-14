package dao

import (
	"log"

	m "github.com/margaritagomez/recipes-back/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RecipesDAO data acces obj
type RecipesDAO struct {
	Server   string
	Database string
}

// DB is the database
var DB *mgo.Database

// Session is the connection to MongoDB
var Session *mgo.Session

const (
	// COLLECTION of recipes
	COLLECTION = "recipes"
)

// Connect to the DB
func (r *RecipesDAO) Connect() {
	Session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = Session.DB(r.Database)
}

// FindAll finds all
func (r *RecipesDAO) FindAll() ([]m.Recipe, error) {
	var recipes []m.Recipe
	err := DB.C(COLLECTION).Find(bson.M{}).All(&recipes)
	return recipes, err
}

// FindByID finds by id
func (r *RecipesDAO) FindByID(id string) (m.Recipe, error) {
	var recipe m.Recipe
	err := DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&recipe)
	return recipe, err
}

// Insert inserts
func (r *RecipesDAO) Insert(recipe m.Recipe) error {
	err := DB.C(COLLECTION).Insert(&recipe)
	return err
}

// Delete deletes
func (r *RecipesDAO) Delete(recipe m.Recipe) error {
	err := DB.C(COLLECTION).Remove(&recipe)
	return err
}

// Update updates
func (r *RecipesDAO) Update(recipe m.Recipe) error {
	err := DB.C(COLLECTION).UpdateId(recipe.ID, &recipe)
	return err
}
