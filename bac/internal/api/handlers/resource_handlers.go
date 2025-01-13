// internal/api/handlers/resource_handlers.go
package handlers

import (
	"net/http"

	"bac/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ResourceHandler struct {
	db *gorm.DB
}

func NewResourceHandler(db *gorm.DB) *ResourceHandler {
	return &ResourceHandler{db: db}
}

func (h *ResourceHandler) GetResources(c *gin.Context) {
	var resources []models.Resource
	if err := h.db.Find(&resources).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resources"})
		return
	}
	c.JSON(http.StatusOK, resources)
}

func (h *ResourceHandler) SearchNearby(c *gin.Context) {
	var params models.SearchParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search parameters"})
		return
	}

	if params.Radius == 0 {
		params.Radius = 5.0 // Default radius in miles
	}

	var results []models.NearbyResource
	query := `
        SELECT 
            id, name, description, address, 
            distance_miles, diagnoses, contact_info
        FROM find_nearby_resources($1, $2, $3, $4)
    `

	if err := h.db.Raw(query,
		params.Latitude,
		params.Longitude,
		params.Radius,
		pq.Array(params.Diagnoses),
	).Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var resource models.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := resource.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&resource).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create resource"})
		return
	}

	c.JSON(http.StatusCreated, resource)
}

func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	id := c.Param("id")
	var resource models.Resource

	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := resource.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Model(&models.Resource{}).Where("id = ?", id).Updates(resource)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update resource"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, resource)
}
