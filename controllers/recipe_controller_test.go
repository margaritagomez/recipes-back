package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	m "github.com/margaritagomez/recipes-back/models"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", DeleteRecipe).Methods("DELETE")
	return r
}

func TestCreateRecipe(t *testing.T) {
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
	jsonRecipe, _ := json.Marshal(recipe)
	request, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(jsonRecipe))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "OK response is expected")

	newRecipe := new(m.Recipe)
	error := json.NewDecoder(response.Body).Decode(newRecipe)
	fmt.Printf(error.Error())
	if error != nil {
		fmt.Printf("holaholaholahola")
		recipe2 := &m.Recipe{
			ID:    newRecipe.ID,
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
		jsonRecipe2, _ := json.Marshal(recipe2)
		request2, _ := http.NewRequest("DELETE", "/recipes", bytes.NewBuffer(jsonRecipe2))
		response2 := httptest.NewRecorder()
		Router().ServeHTTP(response2, request2)
	}
}
