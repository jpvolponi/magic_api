package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Server
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigureRoutes(s.server)
	log.Printf("server is running at port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
