package api

import (
	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// create new Server instance.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts:id", server.GetAccount)
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
