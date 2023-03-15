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

	countryName := vars["countryName"]

	cities, err := city.GetCities(server.DB, countryName)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("City pre: ", cities)

	for i, v := range cities {
		fmt.Println("In Loop", v.ImageLocation)
		imageURL, _ := getImagesS3General(v.ImageLocation)
		cities[i].ImageLocation = imageURL[0]
	}

	fmt.Println("City Final: ", cities)

	responses.JSON(w, http.StatusOK, cities)
}
