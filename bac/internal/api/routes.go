// internal/api/routes.go
package api

import (
	"bac/internal/api/handlers"
	"bac/internal/config"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
	resourceCenterHandler := handlers.NewResourceCenterHandler(s.db)
	regionalCenterHandler := handlers.NewRegionalCenterHandler(s.db)
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, // Allow frontend to access API
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := s.router.Group("/api")
	{
		// Resource routes
		api.GET("/resources", resourceHandler.GetResources)
		api.GET("/resources/nearby", resourceHandler.SearchNearby) // No need for a separate geolocation handler
		api.POST("/resources", resourceHandler.CreateResource)
		api.PUT("/resources/:id", resourceHandler.UpdateResource)

		// Resource Center routes
		api.GET("/resource-center", resourceCenterHandler.GetResourceCenters)
		api.POST("/resource-center", resourceCenterHandler.CreateResourceCenter)
		api.GET("/resource-center/:id", resourceCenterHandler.GetResourceCenterByID)

		// Regional Centers routes
		api.GET("/regional-centers", regionalCenterHandler.GetAllRegionalCenters)
		api.GET("/regional-centers/search", regionalCenterHandler.SearchRegionalCenters)
		api.GET("/regional-centers/nearest", regionalCenterHandler.FindNearestCenters)
		api.GET("/regional-centers/:id", regionalCenterHandler.GetRegionalCenterByID)

		// Debugging route to list all registered routes
		api.GET("/routes", func(c *gin.Context) {
			routes := []string{}
			for _, r := range s.router.Routes() {
				routes = append(routes, fmt.Sprintf("%s %s", r.Method, r.Path))
			}
			c.JSON(http.StatusOK, routes)
		})
	}

	// Debugging: Print all registered routes
	fmt.Println("Registered routes:")
	for _, r := range s.router.Routes() {
		fmt.Println(r.Method, r.Path)
	}
}
func (s *Server) Start() error {
	return s.router.Run(":" + s.config.Port)
}
