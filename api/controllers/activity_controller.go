package controllers

import (
	"errors"
	"fmt"
	"github.com/paulmach/orb"
	"iScore-api/api/auth"
	"strconv"

	//"fmt"
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

func (server *Server) CheckLocationData(w http.ResponseWriter, r *http.Request) {
	//CheckCompletion
	activity := models.Activity{}
	photoLat := r.Header["Lat"][0]
	photoLon := r.Header["Lon"][0]
	photoLatF, err := strconv.ParseFloat(photoLat, 64)
	photoLonF, err := strconv.ParseFloat(photoLon, 64)

	if err != nil {
		fmt.Println(err) // 3.14159265
	}

	coords, err := activity.GetActivityLocation(server.DB, r.Header["Activity"][0])
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	p1 := orb.Point{coords[0].LatX, coords[0].LonX}
	p2 := orb.Point{coords[0].LatY, coords[0].LonY}

	bound := orb.MultiPoint{p1, p2}.Bound()

	fmt.Printf("bound: %+v\n", bound)

	checkPoint := orb.Point{photoLatF, photoLonF}
	if !bound.Contains(checkPoint) {
		fmt.Println("insice check bound")
		responses.JSON(w, http.StatusOK, false)
		return
	}

	server.AddPoints(coords[0].Points, r, w)
	server.CompleteActivity(1, r, w)

	//server.CompleteActivity(coords[0].Points, r, w)
	responses.JSON(w, http.StatusOK, coords[0].Points)

}

func (server *Server) CompleteActivity(points int64, r *http.Request, w http.ResponseWriter) {
	account := models.Account{}
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	resp, err := account.CompleteActivity(server.DB, userId, r.Header["Activity"][0])

	fmt.Println(resp)

}

func (server *Server) AddPoints(points int64, r *http.Request, w http.ResponseWriter) {
	account := models.Account{}
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	resp, err := account.AddPointsDB(server.DB, userId, points)

	fmt.Println(resp)

}

func (server *Server) FillAccountActivities(a *models.Account) {
	fmt.Println(a)
	activity := models.Activity{}
	activityList, _ := activity.GetAllActivities(server.DB)
	fmt.Println(activityList)

	activity.FillAccountActivities(activityList, server.DB, a.AccountId)
	//account := models.Account{}
	//userId, err := auth.ExtractTokenID(r)
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	//	return
	//}
	//resp, err := account.AddPointsDB(server.DB, userId, points)
	//
	//fmt.Println(resp)
}

func (server *Server) CheckIsComplete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Console log: Is complete")
	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	activity := models.Activity{}
	isComplete, err := activity.IsComplete(server.DB, userId, r.Header["Activity"][0])
	fmt.Println(isComplete[0].Completed)
	responses.JSON(w, http.StatusOK, isComplete[0].Completed)
}
