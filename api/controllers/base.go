package controllers

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) http.Handler {

	var err error

	server.DB, err = OpenDB(DbUser, DbPassword, DbPort, DbHost, DbName)

	if err != nil {
		log.Fatal(err)
	}

	server.Router = mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(server.Router)

	server.initializeRoutes()
	return handler
}

func (server *Server) Run(addr string, handler http.Handler) {

	if os.Getenv("DEV") == "true" {

		log.Fatal(http.ListenAndServe(addr, handler))
	}

	log.Fatal(gateway.ListenAndServe(addr, handler))
}

func OpenDB(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	return gorm.Open(postgres.Open(DBURL), &gorm.Config{})
}
