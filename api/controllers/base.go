package controllers

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) http.Handler {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	//server.DB.Debug().AutoMigrate(&models.Card{}, &models.Account{}, &models.AccountCardMap{}, &models.Packs{}, &models.PackCode{}, &models.APIKey{}) //database migration

	server.Router = mux.NewRouter()

	//c := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:3000", "https://master.d1mumm5t7mhs02.amplifyapp.com/", "https://master.d1mumm5t7mhs02.amplifyapp.com", "http://master.d1mumm5t7mhs02.amplifyapp.com/", "http://master.d1mumm5t7mhs02.amplifyapp.com", "*"},
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"Authorization", "ACCEPT"},
	//})
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://main.dgmwlgzg611l2.amplifyapp.com"},
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
	fmt.Println("Listening to port 5000")
	log.Fatal(gateway.ListenAndServe(addr, handler))
}
