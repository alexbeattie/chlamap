package handlers

import (
	"net/http"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Provider struct
type Provider struct {
	ID                  int     `json:"id" gorm:"primaryKey"`
	Name                string  `json:"name"`
	Phone               string  `json:"phone"`
	CoverageAreas       string  `json:"coverage_areas"`
	CenterBasedServices string  `json:"center_based_services"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Areas               string  `json:"areas"` // ✅ Treat as a plain string
}
// ProvidersHandler struct
type ProvidersHandler struct {
	DB *gorm.DB
}

// ✅ Function to initialize the handler
func NewProvidersHandler(db *gorm.DB) *ProvidersHandler {
	if db == nil {
		log.Fatal("Database connection is nil in ProvidersHandler")
	}
	return &ProvidersHandler{DB: db}
}


// ✅ Gin-compatible function to get providers
func (h *ProvidersHandler) GetProviders(c *gin.Context) {
	var providers []Provider

	// ✅ Fetch providers using GORM
	result := h.DB.Find(&providers)
	if result.Error != nil {
		log.Println("Database Query Error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// ✅ Convert `areas` (comma-separated string) into a JSON array
	type ProviderResponse struct {
		ID                  int      `json:"id"`
		Name                string   `json:"name"`
		Phone               string   `json:"phone"`
		CoverageAreas       string   `json:"coverage_areas"`
		CenterBasedServices string   `json:"center_based_services"`
		Latitude            float64  `json:"latitude"`
		Longitude           float64  `json:"longitude"`
		Areas               []string `json:"areas"`
	}

	var formattedProviders []ProviderResponse
	for _, p := range providers {
		areas := []string{}
		if p.Areas != "" {
			areas = splitAreasString(p.Areas) // Helper function
		}

		formattedProviders = append(formattedProviders, ProviderResponse{
			ID:                  p.ID,
			Name:                p.Name,
			Phone:               p.Phone,
			CoverageAreas:       p.CoverageAreas,
			CenterBasedServices: p.CenterBasedServices,
			Latitude:            p.Latitude,
			Longitude:           p.Longitude,
			Areas:               areas,
		})
	}

	// ✅ Return transformed JSON
	c.JSON(http.StatusOK, formattedProviders)
}

// ✅ Helper function to split comma-separated areas
func splitAreasString(s string) []string {
	trimmedAreas := strings.TrimSpace(s)
	if trimmedAreas == "" {
		return []string{}
	}
	return strings.Split(trimmedAreas, ",")
}
