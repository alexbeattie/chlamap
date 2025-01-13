// internal/models/resource.go
package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Resource struct {
	ID          string          `json:"id" gorm:"primarykey"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Diagnoses   []string        `json:"diagnoses" gorm:"type:text[]"`
	Latitude    float64         `json:"latitude"`
	Longitude   float64         `json:"longitude"`
	Address     string          `json:"address"`
	ContactInfo json.RawMessage `json:"contact_info" gorm:"type:jsonb"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type NearbyResource struct {
	Resource
	DistanceMiles float64 `json:"distance_miles"`
}

type SearchParams struct {
	Latitude  float64  `json:"lat"`
	Longitude float64  `json:"lng"`
	Radius    float64  `json:"radius"`
	Diagnoses []string `json:"diagnoses"`
}

func (r *Resource) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Latitude < -90 || r.Latitude > 90 {
		return errors.New("invalid latitude")
	}
	if r.Longitude < -180 || r.Longitude > 180 {
		return errors.New("invalid longitude")
	}
	return nil
}
