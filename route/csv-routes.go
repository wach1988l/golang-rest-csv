package routes

import (
	"Verve_Test_project_rest/controller"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/csv/{id}", controller.GetById).Methods("GET")

}
