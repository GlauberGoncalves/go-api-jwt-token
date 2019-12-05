package routes

import (
	"net/http"

	"github.com/glaubergoncalves/go-api-jwt-token/api/controllers"
	"github.com/glaubergoncalves/go-api-jwt-token/api/middlewares"
)

// CarregaRotas Carrega todas as notas do sistema
func CarregaRotas() {
	http.HandleFunc("/home", middlewares.SetMiddlewareJSON(controllers.Home))
	http.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.Login))

	// usuario
	http.HandleFunc("/usuarios", middlewares.SetMiddlewareAuthentication(controllers.TodosUsuarios))
}
