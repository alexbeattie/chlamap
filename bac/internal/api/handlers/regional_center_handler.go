// internal/api/handlers/regional_center_handler.go

package handlers

import (
	"bac/internal/models" // Change this line to use the local import path
	"fmt"                 // Add this for debug logging
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
)

// RegionalCenterHandler handles regional center-related requests
type RegionalCenterHandler struct {
	DB *gorm.DB
}

// NewRegionalCenterHandler creates a new RegionalCenterHandler instance
func NewRegionalCenterHandler(db *gorm.DB) *RegionalCenterHandler {
	return &RegionalCenterHandler{DB: db}
}
func parseCoordinates(coordStr string) (float64, float64, error) {
	re := regexp.MustCompile(`\((-?\d+\.?\d*),\s*(-?\d+\.?\d*)\)`)
	matches := re.FindStringSubmatch(coordStr)

	if len(matches) < 3 {
		return 0, 0, fmt.Errorf("invalid coordinates: %s", coordStr)
	}

	lat, err1 := strconv.ParseFloat(matches[1], 64)
	lng, err2 := strconv.ParseFloat(matches[2], 64)
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing coordinates: %v, %v", err1, err2)
	}

	return lat, lng, nil
}

// GetAllRegionalCenters retrieves all regional centers
func (h *RegionalCenterHandler) GetAllRegionalCenters(c *gin.Context) {
	var centers []models.RegionalCenter

	result := h.DB.Find(&centers)
	if result.Error != nil {
		fmt.Printf("Error getting centers: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve regional centers"})
		return
	}

	response := make([]map[string]interface{}, 0, len(centers))
	for _, center := range centers {
		// Extract coordinates using regular expressions
		coordStr := center.LocationCoordinates
		fmt.Printf("Processing coordinates: %s\n", coordStr)

		// Look for pattern (number, number) anywhere in the string
		re := regexp.MustCompile(`\((-?\d+\.?\d*),\s*(-?\d+\.?\d*)\)`)
		matches := re.FindStringSubmatch(coordStr)

		if len(matches) < 3 {
			fmt.Printf("No valid coordinates found for center %s\n", center.RegionalCenter)
			continue
		}

		lat, lng, err := parseCoordinates(center.LocationCoordinates)
		if err != nil {
			fmt.Printf("Error parsing coordinates for center %s: %v\n", center.RegionalCenter, err)
			continue
		}
		// if err1 != nil || err2 != nil {
		// 	fmt.Printf("Error parsing coordinates for center %s: %v, %v\n",
		// 		center.RegionalCenter, err1, err2)
		// 	continue
		// }

		// Combine address components
		fullAddress := center.Address
		if center.Suite != "" {
			fullAddress += ", " + center.Suite
		}
		fullAddress += fmt.Sprintf(", %s, %s %s", center.City, center.State, center.ZipCode)

		response = append(response, map[string]interface{}{
			"id":        center.ID,
			"name":      center.RegionalCenter,
			"latitude":  lat,
			"longitude": lng,
			"address":   fullAddress,
			"phone":     center.Telephone,
			"website":   center.Website,
			"type":      center.OfficeType,
		})
	}
	
	fmt.Printf("Sending %d centers\n", len(response))
	c.JSON(http.StatusOK, response)
}

// GetRegionalCenterByID retrieves a specific regional center by ID
func (h *RegionalCenterHandler) GetRegionalCenterByID(c *gin.Context) {
	id := c.Param("id")
	var center models.RegionalCenter

	if result := h.DB.First(&center, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Regional center not found"})
		return
	}

	c.JSON(http.StatusOK, center)
}

// SearchRegionalCenters searches for regional centers based on query parameters
func (h *RegionalCenterHandler) SearchRegionalCenters(c *gin.Context) {
	var centers []models.RegionalCenter
	var totalCount int64

	var regionalCenter models.RegionalCenter
	query := h.DB.Model(&regionalCenter)

	// Apply filters
	if county := c.Query("county"); county != "" {
		query = query.Where("county_served ILIKE ?", "%"+county+"%")
	}
	if district := c.Query("district"); district != "" {
		query = query.Where("los_angeles_health_district ILIKE ?", "%"+district+"%")
	}
	if officeType := c.Query("office_type"); officeType != "" {
		query = query.Where("office_type = ?", officeType)
	}
	if city := c.Query("city"); city != "" {
		query = query.Where("city ILIKE ?", "%"+city+"%")
	}
	query.Count(&totalCount)


	// Add pagination logic
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))       // Default to page 1
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")) // Default to 10 results per page
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// Execute the query
	result := query.Find(&centers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search regional centers"})
		return
	}

	// Return paginated results
	c.JSON(http.StatusOK, gin.H{
		"data":       centers,
		"page":       page,
		"pageSize":   pageSize,
		"totalCount": result.RowsAffected, // Total results matching the query
	})
}
// FindNearestCenters finds regional centers near a given location
func (h *RegionalCenterHandler) FindNearestCenters(c *gin.Context) {
    lat := c.Query("lat")
    lng := c.Query("lng")
    distance := c.DefaultQuery("distance", "16093.4") // Default: 10 miles in meters

    if lat == "" || lng == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
        return
    }

    var centers []struct {
        ID       uint    `json:"id"`
        Name     string  `json:"name"`
        Distance float64 `json:"distance"`
    }

    	query := `
		SELECT id, regional_center AS name, 
		       ST_Y(location) AS latitude, ST_X(location) AS longitude,
		       address || ', ' || COALESCE(suite, '') || ', ' || city || ', ' || state || ' ' || zip_code AS address,
		       telephone AS phone, website, office_type AS type
		FROM regional_centers
		WHERE location IS NOT NULL
	`


if err := h.DB.Raw(query, lat, lng, distance).Scan(&centers).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find nearest centers"})
        return
    }

    c.JSON(http.StatusOK, centers)
}