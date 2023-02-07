package controllers

import (
	"iScore-api/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/verify", middlewares.SetMiddlewareJSON(s.Verify)).Methods("GET")
	//s.Router.HandleFunc("/cards/{id:[0-9]}", s.GetCard).Methods("GET", "OPTIONS")
	//s.Router.HandleFunc("/cards/all", s.GetAllCard).Methods("GET", "OPTIONS")
	//s.Router.HandleFunc("/accounts/{id}", middlewares.SetMiddlewareAuthentication(s.GetAccount)).Methods("GET")
	s.Router.HandleFunc("/accounts-create", middlewares.SetMiddlewareJSON(s.CreateAccount)).Methods("POST", "OPTIONS")
	////s.Router.HandleFunc("/accounts/generate/{id:[0-9]+}/{details:[\\W\\S_]+:[\\W\\S_]+}", s.GenerateAPIKey).Methods("POST")
	s.Router.HandleFunc("/accounts-login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")
	//s.Router.HandleFunc("/packs/generate/{pack:[A-Z]{3}}", middlewares.SetMiddlewareAuthentication(s.PackGenerate)).Methods("POST")
	//s.Router.HandleFunc("/cards/{packCode:[A-Z0-9]{24}}", middlewares.SetMiddlewareAuthentication(s.CardGenerateInit)).Methods("POST", "OPTIONS", "GET")
	//s.Router.HandleFunc("/cards/me", middlewares.SetMiddlewareAuthentication(s.CardList)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/accounts/me", middlewares.SetMiddlewareAuthentication(s.GetAccountDisplay)).Methods("GET", "OPTIONS")
}
