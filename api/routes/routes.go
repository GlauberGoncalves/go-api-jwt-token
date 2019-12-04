package routes

import (
	"net/http"

	"github.com/glaubergoncalves/api-estrutura/api/controllers"
	"github.com/glaubergoncalves/api-estrutura/api/middlewares"
)

// CarregaRotas Carrega todas as notas do sistema
func CarregaRotas() {
	http.HandleFunc("/home", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareJSON(controllers.Home)))
	http.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.Login))

	// usuario
	http.HandleFunc("/usuarios", middlewares.SetMiddlewareAuthentication(controllers.TodosUsuarios))
}
