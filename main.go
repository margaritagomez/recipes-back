package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	c "github.com/margaritagomez/recipes-back/controllers"
	d "github.com/margaritagomez/recipes-back/dao"
)

var dao = d.RecipesDAO{}

// Port to listen
var Port = "5000"

func init() {
	dao.Server = os.Getenv("SERVER")
	dao.Database = os.Getenv("DATABASE")
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", c.GetRecipes).Methods("GET")
	r.HandleFunc("/recipes", c.CreateRecipe).Methods("POST")
	r.HandleFunc("/recipes", c.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/recipes", c.DeleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", c.GetRecipe).Methods("GET")
	if port := os.Getenv("PORT"); port != "" {
		Port = port
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", Port), r); err != nil {
		log.Fatal(err)
	}
}
