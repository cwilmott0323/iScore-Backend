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

	for i, v := range countries {
		imageURL, _ := getImagesS3General(v.ImageLocation)
		countries[i].ImageLocation = imageURL[0]
	}

	responses.JSON(w, http.StatusOK, countries)
}
