package handlers

import (
	"bac/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ResourceCenterHandler struct {
	db *gorm.DB
}

func NewResourceCenterHandler(db *gorm.DB) *ResourceCenterHandler {
	return &ResourceCenterHandler{db: db}
}

func (h *ResourceCenterHandler) GetResourceCenters(c *gin.Context) {
	var centers []models.ResourceCenter
	if err := h.db.Find(&centers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resource centers"})
		return
	}
	c.JSON(http.StatusOK, centers)
}

func (h *ResourceCenterHandler) CreateResourceCenter(c *gin.Context) {
	var center models.ResourceCenter
	if err := c.ShouldBindJSON(&center); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.db.Create(&center).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create resource center"})
		return
	}

	c.JSON(http.StatusCreated, center)
}

func (h *ResourceCenterHandler) GetResourceCenterByID(c *gin.Context) {
	id := c.Param("id")
	var center models.ResourceCenter
	if err := h.db.First(&center, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource center not found"})
		return
	}
	c.JSON(http.StatusOK, center)
}
