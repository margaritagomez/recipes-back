package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RecipesDAO is the Data Access Object
type RecipesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION of recipes
	COLLECTION = "recipes"
)

// Connect performs the conn with mongo
func (m *RecipesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll retrieves all recipes
func (m *RecipesDAO) FindAll() ([]Recipe, error) {
	var recipes []Recipe
	err := db.C(COLLECTION).Find(bson.M{}).All(&recipes)
	return recipes, err
}

func (m *RecipesDAO) FindById(id string) (Recipe, error) {
	var recipe Recipe
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&recipe)
	return recipe, err
}

func (m *RecipesDAO) Insert(recipe Recipe) error {
	err := db.C(COLLECTION).Insert(&recipe)
	return err
}

func (m *RecipesDAO) Delete(recipe Recipe) error {
	err := db.C(COLLECTION).Remove(&recipe)
	return err
}

func (m *RecipesDAO) Update(recipe Recipe) error {
	err := db.C(COLLECTION).UpdateId(recipe.ID, &recipe)
	return err
}
