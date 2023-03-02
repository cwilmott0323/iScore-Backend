package controllers

import (
	"github.com/gorilla/mux"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"log"
	"net/http"
)

func (server *Server) GetActivity(w http.ResponseWriter, r *http.Request) {
	activity := models.Activity{}

	vars := mux.Vars(r)
	log.Println("Vars:", vars)
	countryName := vars["countryName"]
	cityName := vars["cityName"]

	cities, err := activity.GetActivities(server.DB, countryName, cityName)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}

func (server *Server) GetActivityInfo(w http.ResponseWriter, r *http.Request) {
	activity := models.Activity{}

	vars := mux.Vars(r)
	log.Println("Vars:", vars)
	countryName := vars["countryName"]
	cityName := vars["cityName"]
	activityName := vars["activityName"]

	cities, err := activity.GetActivityInfo(server.DB, countryName, cityName, activityName)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}
