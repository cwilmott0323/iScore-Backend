package controllers

import (
	"github.com/gorilla/mux"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"net/http"
)

func (server *Server) GetCities(w http.ResponseWriter, r *http.Request) {
	city := models.City{}

	vars := mux.Vars(r)

	countryName := vars["countryName"]

	cities, err := city.GetCities(server.DB, countryName)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, cities)
}
