package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	d "github.com/margaritagomez/recipes-back/dao"
	m "github.com/margaritagomez/recipes-back/models"
)

var dao = d.RecipesDAO{}

func init() {
	dao.Server = "localhost:27017"
	if dbHostname := os.Getenv("SERVER"); dbHostname != "" {
		dao.Server = dbHostname
	}
	dao.Database = "local"
	if dbName := os.Getenv("DATABASE"); dbName != "" {
		dao.Database = dbName
	}
	dao.Connect()
}

// GetRecipes gets all recipes
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, recipes)
}

// GetRecipe gets a recipe by ID
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := dao.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Recipe ID")
		return
	}
	respondWithJSON(w, http.StatusOK, recipe)
}

// CreateRecipe posts a new recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe m.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	recipe.ID = bson.NewObjectId()
	if err := dao.Insert(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, recipe)
}

// UpdateRecipe puts existing recipe
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe m.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteRecipe deletes existing recipe
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe m.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJSON responds with json
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
