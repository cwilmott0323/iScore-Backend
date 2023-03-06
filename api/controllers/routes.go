package controllers

import (
	"iScore-api/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/verify", middlewares.SetMiddlewareJSON(s.Verify)).Methods("GET")
	s.Router.HandleFunc("/accounts-create", middlewares.SetMiddlewareJSON(s.CreateAccount)).Methods("POST", "OPTIONS")
	s.Router.HandleFunc("/accounts-login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")
	s.Router.HandleFunc("/accounts/me", middlewares.SetMiddlewareAuthentication(s.GetAccountDisplay)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/countries/all", middlewares.SetMiddlewareJSON(s.GetCountries)).Methods("GET")
	s.Router.HandleFunc("/countries/{countryName:[A-Za-z _-]+}/cities", middlewares.SetMiddlewareJSON(s.GetCities)).Methods("GET")
	s.Router.HandleFunc("/countries/{countryName:[A-Za-z _-]+}/cities/{cityName:[A-Za-z _-]+}", middlewares.SetMiddlewareJSON(s.GetActivity)).Methods("GET")
	s.Router.HandleFunc("/countries/{countryName:[A-Za-z _-]+}/cities/{cityName:[A-Za-z _-]+}/{activityName:[A-Za-z _-]+}", middlewares.SetMiddlewareJSON(s.GetActivityInfo)).Methods("GET")
	s.Router.HandleFunc("/upload", middlewares.SetMiddlewareAuthentication(s.Upload)).Methods("POST")
	s.Router.HandleFunc("/personalise", middlewares.SetMiddlewareJSON(s.GetImages)).Methods("GET")
	s.Router.HandleFunc("/location", middlewares.SetMiddlewareAuthentication(s.CheckLocationData)).Methods("GET")
	s.Router.HandleFunc("/complete", middlewares.SetMiddlewareAuthentication(s.CheckIsComplete)).Methods("GET")
}
