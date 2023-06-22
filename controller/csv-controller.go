package controller

import (
	"Verve_Test_project_rest/model"
	"Verve_Test_project_rest/storage"
	"github.com/gorilla/mux"
	"net/http"
)

var CSVData model.CSVData

func GetById(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id := pathVars["id"]
	redisClient := storage.NewRedisClient()
	currentVersion, _ := redisClient.Get("file_version")
	internalId := currentVersion + "_" + id
	val, error := redisClient.Get(internalId)
	if error != nil {
		http.NotFound(w, r)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		byteVal := []byte(val)
		w.Write(byteVal)
	}

}
