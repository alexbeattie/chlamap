// internal/models/aba_center.go
package models

import (
	"time"
	"github.com/google/uuid"
)

// ABACenter represents an ABA therapy center in the database
type ABACenter struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name                 string    `gorm:"not null" json:"name"`
	Street               string    `gorm:"not null" json:"street"`
	City                 string    `gorm:"not null" json:"city"`
	Zip                  string    `gorm:"not null" json:"zip"`
	Phone                string    `gorm:"not null" json:"phone"`
	ServiceType          string    `gorm:"not null" json:"serviceType"`
	WaitlistAvailability string    `json:"waitlistAvailability"`
	WaitlistNotes        string    `json:"waitlistNotes"`
	DxVerification       string    `json:"dxVerification"`
	InsuranceAccepted    string    `json:"insuranceAccepted"`
	MediCalPlans         string    `json:"mediCalPlans"`
	Notes                string    `json:"notes"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name for the ABACenter model
func (ABACenter) TableName() string {
	return "aba_centers"
}

// ABACenterRequest is used for request/response binding
type ABACenterRequest struct {
	Name                 string `json:"name" binding:"required"`
	Street               string `json:"street" binding:"required"`
	City                 string `json:"city" binding:"required"`
	Zip                  string `json:"zip" binding:"required"`
	Phone                string `json:"phone" binding:"required"`
	ServiceType          string `json:"serviceType" binding:"required"`
	WaitlistAvailability string `json:"waitlistAvailability"`
	WaitlistNotes        string `json:"waitlistNotes"`
	DxVerification       string `json:"dxVerification"`
	InsuranceAccepted    string `json:"insuranceAccepted"`
	MediCalPlans         string `json:"mediCalPlans"`
	Notes                string `json:"notes"`
}