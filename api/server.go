package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type Server struct{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("arquivo .env não encontrado")
	}
}

func (s *Server) Run() {

	fmt.Println("Server rodando em localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}
}
