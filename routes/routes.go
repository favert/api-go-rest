package routes

import (
	"log"
	"net/http"

	"github.com/favert/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
