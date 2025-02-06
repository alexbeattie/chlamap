// models/regional_center.go
package models

import (
    "time"
)

type RegionalCenter struct {
    ID                      uint      `json:"id" gorm:"primaryKey"`
    RegionalCenter         string    `json:"regional_center"`
    OfficeType            string    `json:"office_type"`
    Address               string    `json:"address"`
    Suite                 string    `json:"suite"`
    City                  string    `json:"city"`
    State                 string    `json:"state"`
    ZipCode               string    `json:"zip_code"`
    Telephone             string    `json:"telephone"`
    Website               string    `json:"website"`
    CountyServed          string    `json:"county_served"`
    LosAngelesHealthDistrict string `json:"los_angeles_health_district"`
    LocationCoordinates    string    `json:"location_coordinates"`
    CreatedAt             time.Time `json:"created_at"`
    UpdatedAt             time.Time `json:"updated_at"`
}

type RegionalCenterResponse struct {
		ID                      uint      `json:"id"`
		RegionalCenter         string    `json:"regional_center"`
		OfficeType            string    `json:"office_type"`
		Address               string    `json:"address"`
		Suite                 string    `json:"suite"`
		City                  string    `json:"city"`
		State                 string    `json:"state"`
		ZipCode               string    `json:"zip_code"`
		Telephone             string    `json:"telephone"`
		Website               string    `json:"website"`
		CountyServed          string    `json:"county_served"`
		LosAngelesHealthDistrict string `json:"los_angeles_health_district"`
		LocationCoordinates    string    `json:"location_coordinates"`
		CreatedAt             time.Time `json:"created_at"`
		UpdatedAt             time.Time `json:"updated_at"`
}