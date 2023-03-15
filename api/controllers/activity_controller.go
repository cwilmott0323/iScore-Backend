package controllers

import (
	"errors"
	"fmt"
	"github.com/paulmach/orb"
	"iScore-api/api/auth"
	"strconv"

	"github.com/gorilla/mux"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"net/http"
)

func (server *Server) GetActivity(w http.ResponseWriter, r *http.Request) {
	activity := models.Activity{}

	vars := mux.Vars(r)

	countryName := vars["countryName"]
	cityName := vars["cityName"]

	activities, err := activity.GetActivities(server.DB, countryName, cityName)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	fmt.Println("EArly Activities: ", activities)

	for i, v := range activities {
		imageURL, _ := getImagesS3General(v.ImageLocation)
		activities[i].ImageLocation = imageURL[0]
	}

	fmt.Println("Image Location: ", activities)

	responses.JSON(w, http.StatusOK, activities)
}

func (server *Server) GetActivityInfo(w http.ResponseWriter, r *http.Request) {
	activity := models.Activity{}

	vars := mux.Vars(r)

	countryName := vars["countryName"]
	cityName := vars["cityName"]
	activityName := vars["activityName"]

	cities, err := activity.GetActivityInfo(server.DB, countryName, cityName, activityName)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	generalImage, err := getImagesS3General(cities[0].ImageLocation)
	if err != nil {
		fmt.Println(err)
		return
	}

	cities[0].ImageLocation = generalImage[0]

	responses.JSON(w, http.StatusOK, cities)
}

func (server *Server) CheckLocationData(w http.ResponseWriter, r *http.Request) {
	//CheckCompletion
	activity := models.Activity{}
	photoLat := r.Header["Lat"][0]
	photoLon := r.Header["Lon"][0]
	photoLatF, err := strconv.ParseFloat(photoLat, 64)
	photoLonF, err := strconv.ParseFloat(photoLon, 64)

	if err != nil {

	}

	coords, err := activity.GetActivityLocation(server.DB, r.Header["Activity"][0])
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	p1 := orb.Point{coords[0].LatX, coords[0].LonX}
	p2 := orb.Point{coords[0].LatY, coords[0].LonY}

	bound := orb.MultiPoint{p1, p2}.Bound()

	checkPoint := orb.Point{photoLatF, photoLonF}
	if !bound.Contains(checkPoint) {

		responses.JSON(w, http.StatusOK, false)
		return
	}

	server.AddPoints(coords[0].Points, r, w)
	server.CompleteActivity(1, r, w)

	responses.JSON(w, http.StatusOK, coords[0].Points)

}

func (server *Server) CompleteActivity(points int64, r *http.Request, w http.ResponseWriter) {
	account := models.Account{}
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	headerInt, err := strconv.Atoi(r.Header["Activity"][0])
	_, err = account.CompleteActivity(server.DB, userId, int64(headerInt))

}

func (server *Server) AddPoints(points int64, r *http.Request, w http.ResponseWriter) {
	account := models.Account{}
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = account.AddPointsDB(server.DB, userId, points)

}

func (server *Server) CheckIsComplete(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	activity := models.Activity{}
	isComplete, err := activity.IsComplete(server.DB, userId, r.Header["Activity"][0])
	if isComplete == nil {

		responses.JSON(w, http.StatusOK, false)
	} else {
		responses.JSON(w, http.StatusOK, true)
	}

}
