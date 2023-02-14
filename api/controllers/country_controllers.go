package controllers

import (
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"net/http"
)

func (server *Server) GetCountries(w http.ResponseWriter, r *http.Request) {
	country := models.Country{}

	countries, err := country.GetCountries(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, countries)
}
