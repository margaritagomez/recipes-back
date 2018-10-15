package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	d "github.com/margaritagomez/recipes-back/dao"
	m "github.com/margaritagomez/recipes-back/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", GetRecipes).Methods("GET")
	r.HandleFunc("/recipes/{id}", GetRecipe).Methods("GET")
	r.HandleFunc("/recipes", CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes", DeleteRecipe).Methods("DELETE")
	return r
}

const testDBSize = 3
const existingID = "5bc2eb4723848db97f68e94a"
const updateID = "5bc2ef0423848dcc853408fa"

var existingRecipe = &m.Recipe{
	ID:    bson.ObjectIdHex(existingID),
	Title: "TestRecipe",
	Ingredients: []m.Ingredient{
		{
			Name:     "TestIngredient1",
			Quantity: 3,
			Unit:     "TestUnit1",
		},
		{
			Name:     "TestIngredient2",
			Quantity: 3,
			Unit:     "TestUnit2",
		},
		{
			Name:     "TestIngredient3",
			Quantity: 3,
			Unit:     "TestUnit3",
		},
	},
	Instructions: "Test Instructions test instructions test instructions",
	Image:        "https://upload.wikimedia.org/wikipedia/commons/thumb/6/68/Pandebono.jpg/250px-Pandebono.jpg",
}

func TestGetRecipes(t *testing.T) {
	d.COLLECTION = "test"
	// Creates HTTP request
	req, err := http.NewRequest("GET", "/recipes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	// Checks status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: Got %v, want %v",
			status, http.StatusOK)
	}
	testResults := []*m.Recipe{}
	errJSON := json.Unmarshal(rr.Body.Bytes(), &testResults)
	if errJSON != nil {
		t.Fatalf("Error unmarshalling JSON %s",
			errJSON.Error())
	}
	// The amount of returned recipes should be the same as the ones stores in the test db (3)
	if len(testResults) != testDBSize {
		t.Errorf(`Expected "%d" results, got "%d"`, testDBSize, len(testResults))
	}
}

func TestGetRecipe(t *testing.T) {
	d.COLLECTION = "test"

	vars := map[string]string{
		"id": existingID,
	}
	requestURI := url.URL{
		Path: "/recipes/" + existingID,
	}
	req, err := http.NewRequest("GET", requestURI.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	Router().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	testResults := make(map[string]interface{})
	errJSON := json.Unmarshal(rr.Body.Bytes(), &testResults)
	if errJSON != nil {
		t.Errorf("Error unmarshalling JSON %s",
			errJSON.Error())
	}
	// The title should be set
	if _, ok := testResults["title"]; !ok {
		t.Fatalf(`Expected element "title", didn't get it: "%s"`, rr.Body.String())
	}
	// Title should be the same
	if testResults["title"] != existingRecipe.Title {
		t.Errorf(`Expected title "%s", got "%s"`, existingRecipe.Title, testResults["title"])
	}
}

func TestCreateRecipe(t *testing.T) {
	d.COLLECTION = "test"
	recipe := &m.Recipe{
		Title: "test title",
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
		Instructions: "This is an example of instructions",
		Image:        "test url to an image",
	}
	// Recipe is created in database
	jsonRecipe, _ := json.Marshal(recipe)
	request, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(jsonRecipe))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")
	// Response body is parsed
	newRecipe := m.Recipe{}
	errJSON := json.NewDecoder(response.Body).Decode(&newRecipe)
	if errJSON != nil {
		t.Fatalf("Error decoding body response %s",
			errJSON.Error())
	}
	// Recipe gets deleted to leave DB in same state as before
	recipe2 := &m.Recipe{
		ID:           newRecipe.ID,
		Title:        recipe.Title,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		Image:        recipe.Image,
	}
	jsonRecipe2, _ := json.Marshal(recipe2)
	request2, _ := http.NewRequest("DELETE", "/recipes", bytes.NewBuffer(jsonRecipe2))
	response2 := httptest.NewRecorder()
	Router().ServeHTTP(response2, request2)
}

func TestUpdateRecipe(t *testing.T) {
	d.COLLECTION = "test"
	recipe := &m.Recipe{
		ID:    bson.ObjectIdHex(updateID),
		Title: "NewName",
		Ingredients: []m.Ingredient{
			{
				Name:     "Pandebonos",
				Quantity: 34324823,
				Unit:     "testMeasureUnit",
			},
			{
				Name:     "testName2",
				Quantity: 4.3,
				Unit:     "testMeasureUnit 2",
			},
		},
		Instructions: "This is an example of instructions",
		Image:        "test url to an image",
	}
	// Recipe is updated
	jsonRecipe, _ := json.Marshal(recipe)
	request, _ := http.NewRequest("PUT", "/recipes", bytes.NewBuffer(jsonRecipe))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	// Recipe is updated again to be exactly as it was before
	jsonRecipe2, _ := json.Marshal(existingRecipe)
	request2, _ := http.NewRequest("PUT", "/recipes", bytes.NewBuffer(jsonRecipe2))
	response2 := httptest.NewRecorder()
	Router().ServeHTTP(response2, request2)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestDeleteRecipe(t *testing.T) {
	d.COLLECTION = "test"
	recipe := &m.Recipe{
		Title: "test title",
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
		Instructions: "This is an example of instructions",
		Image:        "test url to an image",
	}
	// Recipe is first created to be deleted later on
	jsonRecipe, _ := json.Marshal(recipe)
	request, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(jsonRecipe))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	// Response is parsed to obtain ID
	newRecipe := m.Recipe{}
	err := json.NewDecoder(response.Body).Decode(&newRecipe)
	if err != nil {
		t.Fatal(err)
	}
	// Recipe is successfully deleted
	recipe2 := &m.Recipe{
		ID:           newRecipe.ID,
		Title:        recipe.Title,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		Image:        recipe.Image,
	}
	jsonRecipe2, _ := json.Marshal(recipe2)
	request2, _ := http.NewRequest("DELETE", "/recipes", bytes.NewBuffer(jsonRecipe2))
	response2 := httptest.NewRecorder()
	Router().ServeHTTP(response2, request2)
	assert.Equal(t, 201, response.Code, "OK response is expected")
}
