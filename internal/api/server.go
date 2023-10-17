package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoalish/doodocs-test/internal/service"
	"github.com/gogoalish/doodocs-test/utils"
)

type Server struct {
	router  *gin.Engine
	service *service.Service
	config  *utils.Config
}

func NewServer(config *utils.Config) *Server {
	server := &Server{config: config}
	router := gin.Default()
	apiRoutes := router.Group("/api")
	{
		archiveRoutes := apiRoutes.Group("/archive")
		{
			archiveRoutes.POST("/information", server.archiveHandler)
			archiveRoutes.POST("/files", server.filesHandler)
		}
		mailRoutes := apiRoutes.Group("/mail")
		{
			mailRoutes.POST("/file", server.mailHandler)
		}
	}

	server.router = router
	server.service = service.NewService()
	return server
}

func (server *Server) Start() error {
	return server.router.Run(server.config.Port)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
