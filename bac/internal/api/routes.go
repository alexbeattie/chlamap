// internal/api/routes.go
package api

import (
	"bac/internal/api/handlers"
	"bac/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
	config *config.Config
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		db:     db,
		config: cfg,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	resourceHandler := handlers.NewResourceHandler(s.db)
	geoHandler := handlers.NewGeolocationHandler(s.db)

	api := s.router.Group("/api")
	{
		// Resource routes
		api.GET("/resources", resourceHandler.GetResources)
		api.GET("/resources/nearby", geoHandler.SearchNearby)
		api.POST("/resources", resourceHandler.CreateResource)
		api.PUT("/resources/:id", resourceHandler.UpdateResource)
	}
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
