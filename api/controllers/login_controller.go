package controllers

import (
	"encoding/json"
	"github.com/glaubergoncalves/go-api-jwt-token/api/auth"
	"io/ioutil"
	"net/http"

	formaterror "github.com/glaubergoncalves/go-api-jwt-token/api/util"

	"github.com/glaubergoncalves/go-api-jwt-token/api/models"
	"github.com/glaubergoncalves/go-api-jwt-token/api/responses"
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
