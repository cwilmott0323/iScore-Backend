package controllers

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"iScore-api/api/auth"
	"iScore-api/api/models"
	"iScore-api/api/responses"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	err = createS3Bucket(accountCreated.AccountId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	token, err := auth.CreateToken(uint32(account.AccountId))
	if err != nil {
		responses.JSON(w, http.StatusNotAcceptable, err)
	}

	responses.JSON(w, http.StatusOK, token)

}

func createS3Bucket(accountID int64) error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {

		return errors.New("error: cant create item")
	}

	key := strconv.FormatInt(accountID, 10) + "/"
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_USER_BUCKET")),
		Key:    aws.String(key),
	})

	if err != nil {

		return errors.New("error: cant create item")
	}

	return nil
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

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
