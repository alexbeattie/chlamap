// internal/api/handlers/aba_centers_handler.go

package handlers

import (
	"bac/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// ABACentersHandler handles ABA center-related requests
type ABACentersHandler struct {
	DB *gorm.DB
}

// This should be in internal/api/handlers/aba_centers_handler.go
func NewABACenterHandler(db *gorm.DB) *ABACentersHandler {
    return &ABACentersHandler{DB: db}
}

// CreateABACenter creates a new ABA therapy center
func (h *ABACentersHandler) CreateABACenter(c *gin.Context) {
	var input models.ABACenterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map to database model
	center := models.ABACenter{
		Name:                 input.Name,
		Street:               input.Street,
		City:                 input.City,
		Zip:                  input.Zip,
		Phone:                input.Phone,
		ServiceType:          input.ServiceType,
		WaitlistAvailability: input.WaitlistAvailability,
		WaitlistNotes:        input.WaitlistNotes,
		DxVerification:       input.DxVerification,
		InsuranceAccepted:    input.InsuranceAccepted,
		MediCalPlans:         input.MediCalPlans,
		Notes:                input.Notes,
	}

	// Create record in database
	if result := h.DB.Create(&center); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "ABA Center added successfully",
		"data":    center,
	})
}

// GetABACenters retrieves all ABA centers
func (h *ABACentersHandler) GetABACenters(c *gin.Context) {
	var centers []models.ABACenter

	if result := h.DB.Find(&centers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve ABA centers"})
		return
	}

	c.JSON(http.StatusOK, centers)
}

// GetABACenterByID retrieves a specific ABA center by ID
func (h *ABACentersHandler) GetABACenterByID(c *gin.Context) {
	id := c.Param("id")
	var center models.ABACenter

	if result := h.DB.First(&center, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ABA center not found"})
		return
	}

	c.JSON(http.StatusOK, center)
}

// UpdateABACenter updates an existing ABA center
func (h *ABACentersHandler) UpdateABACenter(c *gin.Context) {
	id := c.Param("id")
	var center models.ABACenter

	// Check if the center exists
	if result := h.DB.First(&center, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ABA center not found"})
		return
	}

	// Parse request body
	var input models.ABACenterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	updates := models.ABACenter{
		Name:                 input.Name,
		Street:               input.Street,
		City:                 input.City,
		Zip:                  input.Zip,
		Phone:                input.Phone,
		ServiceType:          input.ServiceType,
		WaitlistAvailability: input.WaitlistAvailability,
		WaitlistNotes:        input.WaitlistNotes,
		DxVerification:       input.DxVerification,
		InsuranceAccepted:    input.InsuranceAccepted,
		MediCalPlans:         input.MediCalPlans,
		Notes:                input.Notes,
	}

	if result := h.DB.Model(&center).Updates(updates); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ABA center"})
		return
	}

	// Fetch updated center
	h.DB.First(&center, id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ABA center updated successfully",
		"data":    center,
	})
}

// DeleteABACenter deletes an ABA center
func (h *ABACentersHandler) DeleteABACenter(c *gin.Context) {
	id := c.Param("id")
	var center models.ABACenter

	// Check if the center exists
	if result := h.DB.First(&center, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ABA center not found"})
		return
	}

	// Delete the center
	if result := h.DB.Delete(&center); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ABA center"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ABA center deleted successfully",
	})
}

// SearchABACenters searches for ABA centers based on criteria
func (h *ABACentersHandler) SearchABACenters(c *gin.Context) {
	var centers []models.ABACenter
	query := h.DB.Model(&models.ABACenter{})

	// Add filters based on query parameters
	if city := c.Query("city"); city != "" {
		query = query.Where("city ILIKE ?", "%"+city+"%")
	}

	if serviceType := c.Query("service_type"); serviceType != "" {
		query = query.Where("service_type = ?", serviceType)
	}

	if insuranceProvider := c.Query("insurance"); insuranceProvider != "" {
		query = query.Where("insurance_accepted ILIKE ?", "%"+insuranceProvider+"%")
	}

	if mediCal := c.Query("medi_cal"); mediCal == "true" {
		query = query.Where("medi_cal_plans IS NOT NULL AND medi_cal_plans != ''")
	}

	// Execute query
	if result := query.Find(&centers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search ABA centers"})
		return
	}

	c.JSON(http.StatusOK, centers)
}