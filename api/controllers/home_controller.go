package controllers

import (
	"net/http"

	"github.com/glaubergoncalves/api-estrutura/api/responses"
)

func Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Home - API em Golang")
}
