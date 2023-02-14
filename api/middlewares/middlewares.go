package middlewares

import (
	"errors"
	"iScore-api/api/auth"
	"iScore-api/api/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "https://main.dgmwlgzg611l2.amplifyapp.com")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "https://main.dgmwlgzg611l2.amplifyapp.com")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers",
		//	"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
