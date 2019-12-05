package main

import (
	"github.com/glaubergoncalves/go-api-jwt-token/api"
	"github.com/glaubergoncalves/go-api-jwt-token/api/routes"
)

func main() {
	routes.CarregaRotas()

	server := api.Server{}
	server.Run()
}
