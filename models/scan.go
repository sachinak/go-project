package models

import "time"

type ScanReq struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

type Vulnerability struct {
	ID             string    `json:"id"`
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
}
type ScanResult struct {
	ScanID          string          `json:"scan_id"`
	Timestamp       time.Time       `json:"timestamp"`
	ScanStatus      string          `json:"scan_status"`
	ResourceType    string          `json:"resource_type"`
	ResourceName    string          `json:"resource_name"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Summary         struct {
		TotalVulnerabilities int `json:"total_vulnerabilities"`
		SeverityCounts       struct {
			Critical int `json:"CRITICAL"`
			High     int `json:"HIGH"`
			Medium   int `json:"MEDIUM"`
			Low      int `json:"LOW"`
		} `json:"severity_counts"`
		FixableCount int  `json:"fixable_count"`
		Compliant    bool `json:"compliant"`
	} `json:"summary"`
	ScanMetadata struct {
		ScannerVersion  string   `json:"scanner_version"`
		PoliciesVersion string   `json:"policies_version"`
		ScanningRules   []string `json:"scanning_rules"`
		ExcludedPaths   []string `json:"excluded_paths"`
	} `json:"scan_metadata"`
}

type ScanResp struct {
	ScanRes ScanResult `json:"scanResults"`
}

type QueryReq struct {
	Filters map[string]string `json:"filters"`
}
