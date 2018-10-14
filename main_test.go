package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	d "github.com/margaritagomez/recipes-back/dao"
	m "github.com/margaritagomez/recipes-back/models"
	"gopkg.in/mgo.v2/dbtest"
)

// Server holds the dbtest DBServer
var Server dbtest.DBServer

// Recipes fixtures are intentionally setup as map[string]Recipe so I can easily select them from within the tests
var Recipes = map[string]m.Recipe{
	"number1": m.Recipe{
		ID:    "testid",
		Title: "Ding!",
		Ingredients: []m.Ingredient{
			{
				Name:     "testName",
				Quantity: 34324823,
				Unit:     "testMeasureUnit",
			},
			{
				Name:     "testName2",
				Quantity: 4.3,
				Unit:     "testMeasureUnit 2",
			},
		},
		Instructions: "testInstructions",
		Image:        "Test image",
	},
	"number2": m.Recipe{
		ID:    "testid",
		Title: "Ding!",
		Ingredients: []m.Ingredient{
			{
				Name:     "testName",
				Quantity: 34324823,
				Unit:     "testMeasureUnit",
			},
			{
				Name:     "testName2",
				Quantity: 4.3,
				Unit:     "testMeasureUnit 2",
			},
		},
		Instructions: "testInstructions",
		Image:        "Test image",
	},
}

// insertFixtures just inserts all recipes (and other types) I've defined above.
func insertFixtures() {
	for _, recipe := range Recipes {
		if err := d.Session.DB(dao.Database).C("recipes").Insert(recipe); err != nil {
			log.Println(err)
		}
	}
}

// reInsertFixtures drops database and re-inserts all fixtures so we can
// make sure every test can start fresh.
func reInsertFixtures() {
	d.Session.DB(dao.Database).DropDatabase()
	insertFixtures()
}

// TestMain wraps all tests with the needed initialized mock DB and fixtures
func TestMain(m *testing.M) {
	// The tempdir is created so MongoDB has a location to store its files.
	// Contents are wiped once the server stops
	tempDir, _ := ioutil.TempDir("", "testing")
	Server.SetPath(tempDir)

	// My main session var is now set to the temporary MongoDB instance
	d.Session = Server.Session()

	// Make sure to insert my fixtures
	insertFixtures()

	// Run the test suite
	retCode := m.Run()

	// Make sure we DropDatabase so we make absolutely sure nothing is left or locked while wiping the data and
	// close session
	d.Session.DB(dao.Database).DropDatabase()
	d.Session.Close()

	// Stop shuts down the temporary server and removes data on disk.
	Server.Stop()

	// call with result of m.Run()
	os.Exit(retCode)
}
