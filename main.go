package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	c "github.com/margaritagomez/recipes-back/controllers"
)

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", c.getRecipes).Methods("GET")
	r.HandleFunc("/recipes", c.createRecipe).Methods("POST")
	r.HandleFunc("/recipes", c.updateRecipe).Methods("PUT")
	r.HandleFunc("/recipes", c.deleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", c.getRecipe).Methods("GET")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}
