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

type (
	gormOpener func(dialect string, args ...interface{}) (db *gorm.DB, err error)
)

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
	log.Println("Listening to port:", addr)
	if os.Getenv("DEV") == "true" {
		log.Println("DEV")
		log.Fatal(http.ListenAndServe(addr, handler))
	}
	log.Println("PROD")
	log.Fatal(gateway.ListenAndServe(addr, handler))
}

func OpenDB(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	return gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	//return open(Dbdriver, DBURL)
}
