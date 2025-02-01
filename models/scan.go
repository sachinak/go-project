package models

import "time"

type ScanReq struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

type ScanResult struct {
	ScanID          string    `json:"scanId"`
	ScanTime        time.Time `json:"scanTime"`
	Status          string    `json:"status"`
	ResourceDetails struct {
		Type         string `json:"type"`
		Name         string `json:"name"`
		Registry     string `json:"registry"`
		Architecture string `json:"architecture"`
	} `json:"resourceDetails"`
	Findings []struct {
		ID       string  `json:"id"`
		Type     string  `json:"type"`
		CveID    string  `json:"cveId"`
		Severity string  `json:"severity"`
		Score    float64 `json:"score"`
		Package  struct {
			Name         string `json:"name"`
			Version      string `json:"version"`
			Path         string `json:"path"`
			FixedVersion string `json:"fixedVersion"`
		} `json:"package"`
		Description    string    `json:"description"`
		Remediation    string    `json:"remediation"`
		FirstDetected  time.Time `json:"firstDetected"`
		Status         string    `json:"status"`
		Exploitability string    `json:"exploitability"`
		ThreatContext  struct {
			InTheWild       bool   `json:"inTheWild"`
			HasExploit      bool   `json:"hasExploit"`
			ExploitMaturity string `json:"exploitMaturity"`
		} `json:"threatContext"`
	} `json:"findings"`
	Summary struct {
		TotalIssues       int `json:"totalIssues"`
		SeverityBreakdown struct {
			Critical int `json:"CRITICAL"`
			High     int `json:"HIGH"`
			Medium   int `json:"MEDIUM"`
			Low      int `json:"LOW"`
		} `json:"severityBreakdown"`
		FixableIssues int `json:"fixableIssues"`
		SecurityScore int `json:"securityScore"`
	} `json:"summary"`
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
type ScanResult2 struct {
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

type ScanResp2 struct {
	ScanRes ScanResult2 `json:"scanResults"`
}

type ScanResp struct {
	ScanRes ScanResult `json:"scanResults"`
}

type QueryReq struct {
	Filters map[string]string `json:"filters"`
}
