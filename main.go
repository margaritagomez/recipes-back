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
	r.HandleFunc("/recipes", c.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes", c.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", c.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes", c.DeleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", c.GetRecipe).Methods("GET")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}
