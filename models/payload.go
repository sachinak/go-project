package models

import "time"

type Payload struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Severity       string    `json:"severity"`
	Cvss           float64   `json:"cvss"`
	Status         string    `json:"status"`
	PackageName    string    `json:"package_name"`
	CurrentVersion string    `json:"current_version"`
	FixedVersion   string    `json:"fixed_version"`
	Description    string    `json:"description"`
	PublishedDate  time.Time `json:"published_date"`
	Link           string    `json:"link"`
	RiskFactors    string    `json:"risk_factors"`
	ScanTime       time.Time `json:"scan_time"`
	FileName       string    `json:"file_name"`
}

type QueryResponse struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	Severity       string    `json:"severity"`
	Cvss           float64   `json:"cvss"`
	Status         string    `json:"status"`
	PackageName    string    `json:"package_name"`
	CurrentVersion string    `json:"current_version"`
	FixedVersion   string    `json:"fixed_version"`
	Description    string    `json:"description"`
	PublishedDate  time.Time `json:"published_date"`
	Link           string    `json:"link"`
	RiskFactors    []string  `json:"risk_factors"`
	ScanTime       time.Time `json:"scan_time"`
	FileName       string    `json:"file_name"`
}
