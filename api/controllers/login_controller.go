package controllers

import (
	"encoding/json"
	"github.com/glaubergoncalves/api-estrutura/api/auth"
	"io/ioutil"
	"net/http"

	formaterror "github.com/glaubergoncalves/api-estrutura/api/util"

	"github.com/glaubergoncalves/api-estrutura/api/models"
	"github.com/glaubergoncalves/api-estrutura/api/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	usuario := models.Usuario{}
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	usuario.Prepara()
	err = usuario.Valida("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(usuario.Email, usuario.Senha)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func SignIn(email, senha string) (string, error) {

	usuario := models.Usuario{}

	// logica para verificar no banco a autenticação
	// ...

	usuario.ID = 100

	return auth.CriaToken(usuario.ID)
}
