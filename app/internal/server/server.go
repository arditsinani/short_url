package server

import (
	"short_url/app/config"
	"short_url/app/internal/controllers"
	"short_url/app/internal/db"
	"short_url/app/internal/services"

	"github.com/gin-gonic/gin"
)

// Server type stores db connection and controllers
type Server struct {
	Config      *config.Config
	DB          *db.DB
	Controllers controllers.Controllers
}

func (s *Server) Run() {
	//init services
	srvs := services.Services{
		RedirService: services.RedirService{
			DB: s.DB,
		},
	}
	//init controllers
	s.Controllers = controllers.Controllers{
		Redir: controllers.RedirCtrl{
			Config:       s.Config,
			DB:           s.DB,
			RedirService: srvs.RedirService,
		},
	}
	ginEngine := gin.Default()
	//init routes
	s.initRoutes(ginEngine)
	//run server
	ginEngine.Run(s.Config.Server.Port) // listen and serve on 0.0.0.0:8080
}

// package constructor
func New(conf *config.Config, db *db.DB) {
	web := Server{Config: conf, DB: db}
	web.Run()
}
