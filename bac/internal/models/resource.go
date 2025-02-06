package models

import (
    "errors"
    "time"
    "github.com/lib/pq"
)

type Resource struct {
    ID          string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Address     string    `json:"address"`
    Latitude    float64   `json:"latitude"`
    Longitude   float64   `json:"longitude"`
	Diagnoses   pq.StringArray `gorm:"type:text[]"` // Correct type for PostgreSQL array
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type ResourceResponse struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Address     string    `json:"address"`
    Latitude    float64   `json:"latitude"`
    Longitude   float64   `json:"longitude"`
	Diagnoses   []string  `json:"diagnoses"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// ResourceCenter defines the spatial table for resource centers
type ResourceCenter struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // Use default for UUID generation
	Name      string    `gorm:"not null"`
	Location  string    `gorm:"type:geometry(Point,4326)"` // PostGIS Point
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
// NearbyResource extends Resource with distance information
type NearbyResource struct {
    Resource
    Distance float64 `json:"distance"` // Distance in meters
}

func (r *Resource) Validate() error {
    if r.Name == "" {
        return errors.New("name is required")
    }
    if r.Address == "" {
        return errors.New("address is required")
    }
    if r.Description == "" {
        return errors.New("description is required")
    }
    return nil
}