package controllers

import (
	"iScore-api/api/responses"
	"net/http"
)

func (server *Server) Verify(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Verify")
}
