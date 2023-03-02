package api

import (
	"github.com/jinzhu/gorm"
	"iScore-api/api/controllers"
	"iScore-api/api/models"
	"os"
)

var server = controllers.Server{}

func Run() error {

	var err error

	handler := server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), gorm.Open)

	models.Load(server.DB)

	server.Run(":"+os.Getenv("PORT"), handler)
	return err

}
