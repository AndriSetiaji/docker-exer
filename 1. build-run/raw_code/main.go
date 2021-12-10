// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var App_port string
var App_message string

// App export
type App struct {
	Router *mux.Router
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	m := "{ \"message\": \"" + App_message + "\" }"
	json.Unmarshal([]byte(m), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", helloWorldHandler)
}

func (app *App) run() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	App_message = os.Getenv("MESSAGE")
	App_port = os.Getenv("PORT")
	fmt.Println("service running  127.0.0.1:" + App_port)
	log.Fatal(http.ListenAndServe(":"+App_port, app.Router))

	// fmt.Println("service running on port :" + App_port)
}

func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
