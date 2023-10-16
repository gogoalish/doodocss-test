package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoalish/doodocs-test/internal/service"
)

type Server struct {
	router  *gin.Engine
	service *service.Service
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()
	apiRoutes := router.Group("/api")
	{
		archiveRoutes := apiRoutes.Group("/archive")
		{
			archiveRoutes.POST("/information", server.archiveHandler)
		}
	}

	server.router = router
	server.service = service.NewService()
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
