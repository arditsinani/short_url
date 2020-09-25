package server

import (
	"github.com/gin-gonic/gin"
)

// Creates server routes
func (s *Server) initRoutes(server *gin.Engine) {
	// NOTE: gin framework doesnt support multiple paths which one is static and one is dynamic
	// this is a workaround to have /api and :redirect as parameter under the same path
	// /api path is checked later in the handler method
	server.GET("/:redirect", s.Controllers.Redir.Redirect)
	server.GET("/:redirect/read/:url", s.Controllers.Redir.ReadByUrl)
	server.POST("/:redirect/create", s.Controllers.Redir.Create)
	server.DELETE("/:redirect/delete/:url", s.Controllers.Redir.Delete)
}
