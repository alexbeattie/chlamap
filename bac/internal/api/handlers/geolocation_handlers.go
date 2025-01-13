// internal/api/handlers/geolocation_handlers.go
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"bac/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GeolocationHandler struct {
	db *gorm.DB
}

func NewGeolocationHandler(db *gorm.DB) *GeolocationHandler {
	return &GeolocationHandler{db: db}
}

func (h *GeolocationHandler) SearchNearby(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude parameter"})
		return
	}

	lng, err := strconv.ParseFloat(c.Query("lng"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude parameter"})
		return
	}

	radius, err := strconv.ParseFloat(c.DefaultQuery("radius", "5"), 64)
	if err != nil {
		radius = 5.0
	}

	var diagnoses []string
	if diagnosesParam := c.Query("diagnoses"); diagnosesParam != "" {
		diagnoses = pq.StringArray(strings.Split(diagnosesParam, ","))
	}

	var results []models.NearbyResource
	query := `
        SELECT 
            id, name, description, address, 
            distance_miles, diagnoses, contact_info
        FROM find_nearby_resources($1, $2, $3, $4)
    `

	if err := h.db.Raw(query, lat, lng, radius, pq.Array(diagnoses)).Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search nearby resources"})
		return
	}

	c.JSON(http.StatusOK, results)
}
