package dao

import (
	"log"

	mod "github.com/margaritagomez/recipes-back/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RecipesDAO data acces obj
type RecipesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION of recipes
	COLLECTION = "recipes"
)

// Connect to the db
func (m *RecipesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll finds all
func (m *RecipesDAO) FindAll() ([]mod.Recipe, error) {
	var recipes []mod.Recipe
	err := db.C(COLLECTION).Find(bson.M{}).All(&recipes)
	return recipes, err
}

// FindByID finds by id
func (m *RecipesDAO) FindByID(id string) (mod.Recipe, error) {
	var recipe mod.Recipe
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&recipe)
	return recipe, err
}

// Insert inserts
func (m *RecipesDAO) Insert(recipe mod.Recipe) error {
	err := db.C(COLLECTION).Insert(&recipe)
	return err
}

// Delete deletes
func (m *RecipesDAO) Delete(recipe mod.Recipe) error {
	err := db.C(COLLECTION).Remove(&recipe)
	return err
}

// Update updates
func (m *RecipesDAO) Update(recipe mod.Recipe) error {
	err := db.C(COLLECTION).UpdateId(recipe.ID, &recipe)
	return err
}
