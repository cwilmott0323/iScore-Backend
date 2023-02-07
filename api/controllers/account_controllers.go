package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"iScore-api/api/auth"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (server *Server) GetAccountAll(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	account := models.Account{}

	//server.CheckAuth(w, r)
	accountGotten, err := account.FindAccountByID(server.DB, userId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, accountGotten)
}

func (server *Server) GetAccountDisplay(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	account := models.Account{}

	accountGotten, err := account.FindAccountByID(server.DB, userId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, accountGotten)
}

func (server *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account := models.Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//vars := mux.Vars(r)
	//createDetails := strings.Split(vars["create"], ":")
	//
	//hashedPassword := saltPassword(createDetails)

	if len(account.Password) == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	hashedPassword := saltPassword(account.Password)

	account = models.Account{
		Name:      account.Name,
		Password:  hashedPassword,
		Email:     account.Email,
		Points:    0,
		CreatedOn: time.Now(),
		LastLogin: time.Now(),
	}

	accountCreated, err := account.CreateAccount(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, accountCreated)

}

func saltPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func (server *Server) verifyAccount(w http.ResponseWriter, r *http.Request) bool {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return false
	}

	account := models.Account{}
	accountGotten, err := account.FindAccountByID(server.DB, uint32(uid))

	userpass := strings.Split(vars["details"], ":")
	approved := checkPasswordHash(userpass[1], accountGotten.Password)
	if approved == false {
		responses.JSON(w, http.StatusBadRequest, "User Details incorrect")
		return false
	}

	if userpass[0] != accountGotten.Email {
		responses.JSON(w, http.StatusBadRequest, "User Details incorrect")
		return false
	}
	return true
}

func checkPasswordHash(password string, hashedPassword string) bool {
	passwordByte := []byte(password)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), passwordByte)
	if err != nil {
		return false
	} else {
		return true
	}
}

//func (server *Server) CheckAuth(w http.ResponseWriter, r *http.Request) {
//	if _, ok := r.Header["Authorization"]; !ok {
//		responses.JSON(w, http.StatusBadRequest, "Authentication header missing")
//		return
//	}
//	api_key := r.Header["Authorization"][0]
//	KeyCheck := models.APIKey{APIKey: api_key}
//	KeyCheck.CheckAuth(server.DB)
//}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
