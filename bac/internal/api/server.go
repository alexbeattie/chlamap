package api

import (
	"bac/internal/api/handlers"
	authMiddleware "bac/internal/api/middleware/auth" // Import with alias
	"bac/internal/config"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	
	"gorm.io/gorm"
)
type Server struct {
	router *gin.Engine
	db     *gorm.DB
	config *config.Config
	server *http.Server
	middleware struct {
		AuthMiddleware    gin.HandlerFunc
		RequirePermission func(string) gin.HandlerFunc
  }
}

type Logger interface {
    Error(message string, args ...interface{})
    // Other methods as needed
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8081", "http://localhost:8080", "http://192.168.1.158:8080", "http://alex-macbookpro.local:8080"},
		// AllowAllOrigins: true, // TEMPORARY: Allow all origins (use carefully in production)

		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server := &Server{
		router: router,
		db:     db,
		config: cfg,
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
	}
	
	// Initialize middleware
	server.middleware.AuthMiddleware = authMiddleware.AuthMiddleware([]byte(cfg.JWTSecret))
	server.middleware.RequirePermission = authMiddleware.RequirePermission
		
	// Register routes
	server.RegisterAuthRoutes()
	server.setupRoutes()
	return server
}

// Rest of the file stays the same
func (s *Server) setupRoutes() {
	resourceHandler := handlers.NewResourceHandler(s.db)
	geoHandler := handlers.NewGeolocationHandler(s.db)
	regionalCenterHandler := handlers.NewRegionalCenterHandler(s.db)
	abaCentersHandler := handlers.NewABACenterHandler(s.db)
	providersHandler := handlers.NewProvidersHandler(s.db)
	api := s.router.Group("/api")
	{
		api.HEAD("/regional-centers", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// Existing routes remain the same
		api.GET("/resources/nearby", geoHandler.SearchNearby)
		api.GET("/resources", resourceHandler.GetResources)
		api.POST("/resources", resourceHandler.CreateResource)
		api.GET("/resources/:id", resourceHandler.GetResource)
		api.PUT("/resources/:id", resourceHandler.UpdateResource)
		api.DELETE("/resources/:id", resourceHandler.DeleteResource)
		api.GET("/resource-center", resourceHandler.GetResourceCenters)
		api.POST("/resource-center", resourceHandler.CreateResourceCenter)
		api.GET("/resource-center/:id", resourceHandler.GetResourceCenterByID)

		// Regional Centers routes - simplified and corrected
		api.GET("/regional-centers", regionalCenterHandler.GetAllRegionalCenters)
		api.GET("/regional-centers/search", regionalCenterHandler.SearchRegionalCenters)
		api.GET("/regional-centers/nearest", regionalCenterHandler.FindNearestCenters)
		api.GET("/regional-centers/:id", regionalCenterHandler.GetRegionalCenterByID)

		api.GET("/aba-centers", abaCentersHandler.GetABACenters)
		api.POST("/aba-centers", abaCentersHandler.CreateABACenter)
		api.GET("/aba-centers/search", abaCentersHandler.SearchABACenters)
		api.GET("/aba-centers/:id", abaCentersHandler.GetABACenterByID)
		api.PUT("/aba-centers/:id", abaCentersHandler.UpdateABACenter)
		api.DELETE("/aba-centers/:id", abaCentersHandler.DeleteABACenter)

		api.GET("/providers", providersHandler.GetProviders)

		// Debug route
		api.GET("/routes", func(c *gin.Context) {
			routes := []string{}
			for _, r := range s.router.Routes() {
				routes = append(routes, fmt.Sprintf("%s %s", r.Method, r.Path))
			}
			c.JSON(http.StatusOK, routes)
		})
	}
}

func (s *Server) Start() error {
		fmt.Println("Server running on port:", s.server.Addr)

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}