# Golang RESTful API for recipes

This project has been deployed in heroku [here](https://evening-peak-29761.herokuapp.com/recipes).

## Technologies used
 * This project was developed in Go. 
 * Gorilla Mux was used for routing and dispatching HTTP requests.
 * For persistence, a MongoDB is used. 
 * Godeps is used to vendor the app's dependencies into its source repository, so it could be deployed.

## Structure
```
.
├── controllers
|   ├── recipe_controller_test.go
|   └── recipe_controller.go
├── dao
|   └── recipe_dao.go
├── Godeps
|   ├── Godeps.json
|   └── Readme
├── models
|   ├── ingredient.go
|   └── recipe.go
├── vendor
|   └── ...
├── Procfile
└── main.go
```
### Models
A recipe has five properties. The ID is automatically generated once the recipe is posted to the database, and it has an array of ingredients. These can be found in the `/models` dir.
#### Recipe Model
```
{
	"id":           ObjectId 
	"title":        string       
	"ingredients":  []Ingredient  
	"instructions": string        
	"image":        string        
}
```
#### Ingredient Model
```
{
	"name":           string 
	"quantity":       float32       
	"unit":           string   
}
```
### Data Access Object
To access the persistence layer the `mgo` package is used. The DB connection is done by the function `Connect()`, which uses the Server and Database variables, that can be set as env variables. If they are not found, it establishes a connection on default port 27017, which would need a running instance of MongoDB if it were to be ran locally. 

The other implemented functions are `FindAll`, `FindByID`, `Insert`, `Delete`, and `Update`; whose names are pretty self-explanatory. 


### Controllers
This is where the HTTP request handlers are implemented, they all respond with JSON format. 
#### Controllers unit tests
Each handler is tested and each test is self-contained. However it does not use a mock DB, it uses a different test collection, that is only addressed by these functions. 

When running `go test -v` inside the `/controllers` directory, there should be an output like the following:

```
=== RUN   TestGetRecipes
--- PASS: TestGetRecipes (0.08s)
=== RUN   TestGetRecipe
--- PASS: TestGetRecipe (0.08s)
=== RUN   TestCreateRecipe
--- PASS: TestCreateRecipe (0.17s)
=== RUN   TestUpdateRecipe
--- PASS: TestUpdateRecipe (0.19s)
=== RUN   TestDeleteRecipe
--- PASS: TestDeleteRecipe (0.20s)
PASS
ok      github.com/margaritagomez/recipes-back/controllers      1.283s
```
### main.go

Here is where execution begins, the port and routes are set. 

To call the back-end services you must call:

**Base URL:** https://evening-peak-29761.herokuapp.com

| HTTP method  | URL path      | Body       |
|--------------|---------------|------------|
| GET          | /recipes      | --         |
| GET          | /recipes/:id  | --         |
| POST         | /recipes      |` { recipe } `|
| PUT          | /recipes      |` { recipe } `|
| DELETE       | /recipes      |` { recipe } `|
