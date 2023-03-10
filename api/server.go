package api

import (
	"iScore-api/api/controllers"
	"iScore-api/api/models"
	"os"
)

var server = controllers.Server{}

func Run() error {

	var err error

	handler := server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	models.Load(server.DB)

	server.Run(":"+os.Getenv("PORT"), handler)
	return err

}
