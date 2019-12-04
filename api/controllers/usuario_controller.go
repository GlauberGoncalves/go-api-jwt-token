package controllers

import (
	"net/http"

	"github.com/glaubergoncalves/api-estrutura/api/models"
	"github.com/glaubergoncalves/api-estrutura/api/responses"
)

func TodosUsuarios(w http.ResponseWriter, r *http.Request) {

	usuario := models.Usuario{}
	lista, _ := usuario.FindAllUsuarios()
	responses.JSON(w, http.StatusOK, lista)
}
