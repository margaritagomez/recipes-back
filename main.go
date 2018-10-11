package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	pConfig "github.com/margaritagomez/recipes-back/config"
	pDao "github.com/margaritagomez/recipes-back/dao"
	pModels "github.com/margaritagomez/recipes-back/models"
)

var config = pConfig.Config{}
var dao = pDao.RecipesDAO{}

// getRecipes gets all recipes
func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, recipes)
}

// getRecipe gets a recipe by ID
func getRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := dao.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Recipe ID")
		return
	}
	respondWithJSON(w, http.StatusOK, recipe)
}

// createRecipe posts a new recipe
func createRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe pModels.Recipe
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

// updateRecipe puts existing recipe
func updateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe pModels.Recipe
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

// deleteRecipe deletes existing recipe
func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe pModels.Recipe
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

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	//config.Read()

	dao.Server = os.Getenv("SERVER")
	dao.Database = os.Getenv("DATABASE")
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/recipes", createRecipe).Methods("POST")
	r.HandleFunc("/recipes", updateRecipe).Methods("PUT")
	r.HandleFunc("/recipes", deleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")
	if err := http.ListenAndServe(os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}
