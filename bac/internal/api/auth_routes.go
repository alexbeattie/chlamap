package api

import (
	"bac/internal/auth"
	
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

// RegisterAuthRoutes registers the authentication routes
func (s *Server) RegisterAuthRoutes() {
	// Create auth service
	authService := auth.NewAuthService(s.db, []byte(s.config.JWTSecret))

	s.router.POST("/api/register", func(c *gin.Context) {
		var req auth.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println("JSON binding error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := authService.Register(req); err != nil {
			log.Println("Registration error:", err)

			if err.Error() == "email already exists" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	s.router.POST("/api/login", func(c *gin.Context) {
		var req auth.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := authService.Login(req)
		if err != nil {
			if err.Error() == "invalid credentials" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
			return
		}

		c.JSON(http.StatusOK, response)
	})

	// Protected routes group
	authGroup := s.router.Group("/api")
	authGroup.Use(s.middleware.AuthMiddleware)
	{
		authGroup.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")

			// Ensure userID is of correct type
			if userIDInt, ok := userID.(int); ok {
				profile, err := authService.GetUserProfile(userIDInt)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
					return
				}
				c.JSON(http.StatusOK, profile)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
			}
		})

		authGroup.GET("/users", s.middleware.RequirePermission("read:users"), func(c *gin.Context) {
			users, err := authService.GetAllUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"users": users})
		})
	}
}

