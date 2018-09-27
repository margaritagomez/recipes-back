package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/margaritagomez/recipes-back/config"
	. "github.com/margaritagomez/recipes-back/dao"
	. "github.com/margaritagomez/recipes-back/model"
)

var config = Config{}
var dao = RecipesDAO{}

// AllRecipeEndPoint GET list of recipes
func AllRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	recipes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, recipes)
}

// FindRecipeEndpoint GET a recipe by its ID
func FindRecipeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Recipe ID")
		return
	}
	respondWithJson(w, http.StatusOK, recipe)
}

// CreateRecipeEndPoint POST a new recipe
func CreateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	recipe.ID = bson.NewObjectId()
	if err := dao.Insert(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, recipe)
}

// UpdateRecipeEndPoint PUT update an existing recipe
func UpdateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteRecipeEndPoint DELETE an existing recipe
func DeleteRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", AllRecipesEndPoint).Methods("GET")
	r.HandleFunc("/recipes", CreateRecipeEndPoint).Methods("POST")
	r.HandleFunc("/recipes", UpdateRecipeEndPoint).Methods("PUT")
	r.HandleFunc("/recipes", DeleteRecipeEndPoint).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", FindRecipeEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
