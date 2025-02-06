package handlers

import (
	"bac/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
)

type ResourceHandler struct {
	DB *gorm.DB
}

func NewResourceHandler(db *gorm.DB) *ResourceHandler {
	return &ResourceHandler{DB: db}
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var input models.ResourceResponse
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource := models.Resource{
		Name:        input.Name,
		Description: input.Description,
		Address:     input.Address,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Diagnoses:   pq.StringArray(input.Diagnoses),
	}

	if result := h.DB.Create(&resource); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Convert back to response format
	response := models.ResourceResponse{
		ID:          resource.ID,
		Name:        resource.Name,
		Description: resource.Description,
		Address:     resource.Address,
		Latitude:    resource.Latitude,
		Longitude:   resource.Longitude,
		Diagnoses:   input.Diagnoses,
		CreatedAt:   resource.CreatedAt,
		UpdatedAt:   resource.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	id := c.Param("id")
	var resource models.Resource

	// Check if the resource exists
	if result := h.DB.First(&resource, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	// Parse input JSON
	var input models.ResourceResponse
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the resource
	updateData := models.Resource{
		Name:        input.Name,
		Description: input.Description,
		Address:     input.Address,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Diagnoses:   pq.StringArray(input.Diagnoses),
	}

	if result := h.DB.Model(&resource).Updates(updateData); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update resource"})
		return
	}

	// Fetch updated resource
	h.DB.First(&resource, "id = ?", id)
	response := models.ResourceResponse{
		ID:          resource.ID,
		Name:        resource.Name,
		Description: resource.Description,
		Address:     resource.Address,
		Latitude:    resource.Latitude,
		Longitude:   resource.Longitude,
		Diagnoses:   input.Diagnoses,
		CreatedAt:   resource.CreatedAt,
		UpdatedAt:   resource.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ResourceHandler) GetResource(c *gin.Context) {
    id := c.Param("id")

    // Validate if id is a UUID before querying the database
    if _, err := uuid.Parse(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID format"})
        return
    }

    var resource models.Resource
    if result := h.DB.First(&resource, "id = ?", id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
        return
    }

    response := models.ResourceResponse{
        ID:          resource.ID,
        Name:        resource.Name,
        Description: resource.Description,
        Address:     resource.Address,
        Latitude:    resource.Latitude,
        Longitude:   resource.Longitude,
        Diagnoses:   []string(resource.Diagnoses),
        CreatedAt:   resource.CreatedAt,
        UpdatedAt:   resource.UpdatedAt,
    }

    c.JSON(http.StatusOK, response)
}


func (h *ResourceHandler) GetResourceCenterByID(c *gin.Context) {
	id := c.Param("id") // Extract the ID from the URL
	var center models.ResourceCenter

	// Look for the resource center in the database
	if err := h.DB.First(&center, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource center not found"})
		return
	}

	// Return the resource center as JSON
	c.JSON(http.StatusOK, center)
}

func (h *ResourceHandler) DeleteResource(c *gin.Context) {
	id := c.Param("id")
	var resource models.Resource

	// Check if the resource exists
	if result := h.DB.First(&resource, "id = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	// Delete the resource
	if result := h.DB.Delete(&resource); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete resource"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}
func (h *ResourceHandler) GetResources(c *gin.Context) {
	var resources []models.Resource
	if result := h.DB.Find(&resources); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Convert resources to ResourceResponse format
	response := make([]models.ResourceResponse, len(resources))
	for i, r := range resources {
		response[i] = models.ResourceResponse{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
			Address:     r.Address,
			Latitude:    r.Latitude,
			Longitude:   r.Longitude,
			Diagnoses:   []string(r.Diagnoses), // Convert pq.StringArray to []string
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *ResourceHandler) CreateResourceCenter(c *gin.Context) {
	var center models.ResourceCenter
	if err := c.ShouldBindJSON(&center); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the location field based on latitude and longitude
	if center.Latitude != 0 && center.Longitude != 0 {
		center.Location = fmt.Sprintf("SRID=4326;POINT(%f %f)", center.Longitude, center.Latitude)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and Longitude are required"})
		return
	}

	if err := h.DB.Create(&center).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, center)
}

func (h *ResourceHandler) GetResourceCenters(c *gin.Context) {
	var centers []models.ResourceCenter
	query := h.DB

	// Optional filter: diagnosis
	if diagnosis := c.Query("diagnosis"); diagnosis != "" {
		query = query.Joins("JOIN center_diagnoses cd ON resource_centers.id = cd.center_id").
			Joins("JOIN diagnoses d ON cd.diagnosis_id = d.id").
			Where("d.name = ?", diagnosis)
	}

	// Optional filter: location and radius
	if lat := c.Query("lat"); lat != "" {
		if lng := c.Query("lng"); lng != "" {
			radius := c.DefaultQuery("radius", "5000") // Default to 5 km
			query = query.Where(
				"ST_DWithin(location, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)",
				lng, lat, radius,
			)
		}
	}

	// Execute the query
	if err := query.Find(&centers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, centers)
}
