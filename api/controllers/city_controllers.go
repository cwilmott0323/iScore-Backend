package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"net/http"
)

func (server *Server) GetCities(w http.ResponseWriter, r *http.Request) {
	city := models.City{}

	vars := mux.Vars(r)
	fmt.Println("Vars:", vars)
	countryName := vars["countryName"]

	cities, err := city.GetCities(server.DB, countryName)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}

func (server *Server) GetCitiesInfo(w http.ResponseWriter, r *http.Request) {
	city := models.CityInfo{}

	vars := mux.Vars(r)
	fmt.Println("Vars:", vars)
	countryName := vars["countryName"]
	cityName := vars["cityName"]

	cities, err := city.GetCitiesInfo(server.DB, countryName, cityName)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}
