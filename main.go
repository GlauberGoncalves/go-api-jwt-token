package main

import (
	"github.com/glaubergoncalves/api-estrutura/api"
	"github.com/glaubergoncalves/api-estrutura/api/routes"
)

func main() {
	routes.CarregaRotas()

	server := api.Server{}
	server.Run()
}
