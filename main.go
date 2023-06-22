package main

import (
	"Verve_Test_project_rest/route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:1321", r))
}
