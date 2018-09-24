package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AllRecipesEndPoint returns all recipes.
func AllRecipesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// FindRecipeEndpoint finds a recipe.
func FindRecipeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// CreateRecipeEndPoint creates a recipe.
func CreateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// UpdateRecipeEndPoint updates a recipe.
func UpdateRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteRecipeEndPoint deletes a recipe.
func DeleteRecipeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

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
