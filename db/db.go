package db

import (
	"encoding/json"
	"fmt"
	"go-project/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Function to fetch vulnerabilities from DB using provided filters
func GetVulnerabilitesFromDBWithFilters(filters map[string]string) []models.QueryResponse {

	var payloads []models.Payload
	var response []models.QueryResponse
	db, err := gorm.Open(sqlite.Open("projectDB"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return response
	}

	query := db.Table("vulnerabilities")

	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	err = query.Find(&payloads).Error
	if err != nil {
		fmt.Println("error fetching data:", err)
	}

	for _, data := range payloads {
		var queryResp models.QueryResponse
		json.Unmarshal([]byte(data.RiskFactors), &queryResp.RiskFactors)
		queryResp.ID = data.ID
		queryResp.Severity = data.Severity
		queryResp.Cvss = data.Cvss
		queryResp.Status = data.Status
		queryResp.PackageName = data.PackageName
		queryResp.CurrentVersion = data.CurrentVersion
		queryResp.FixedVersion = data.FixedVersion
		queryResp.Description = data.Description
		queryResp.PublishedDate = data.PublishedDate
		queryResp.Link = data.Link
		queryResp.ScanTime = data.ScanTime
		queryResp.FileName = data.FileName
		response = append(response, queryResp)
	}
	return response
}

// Function to write files to DB
func WritePayloadToDB(data models.Vulnerability, scanTime time.Time, fileName string) {
	db, err := gorm.Open(sqlite.Open("projectDB"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
	riskFactorsJSON, _ := json.Marshal(data.RiskFactors)

	err = db.Exec(`INSERT INTO vulnerabilities (
        id, severity, cvss, status, package_name, current_version, fixed_version,
        description, published_date, link, risk_factors, file_name, scan_time
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		data.ID, data.Severity, data.Cvss, data.Status, data.PackageName,
		data.CurrentVersion, data.FixedVersion, data.Description,
		data.PublishedDate, data.Link, string(riskFactorsJSON), fileName, scanTime).Error
	if err != nil {
		fmt.Println("error in inserting in db:", err)
	}

}
