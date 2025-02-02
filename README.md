# GoProject


Running the code:
1. On Local Machine:  
>Run `go run main.go` in the `go-project` directory

2. On Docker:
>    - Building the image:  
>Run `docker buildx build --platform linux/arm64 --no-cache -t go-project .` in the `go-project` directory
>    - Running the container:  
>Run `docker run -p 8080:8080 go-project` in the `go-project` directory

  
Testing the code:
>1. Run `go test . -v` from the main `go-project` directory
>To see code coverage and all the code lines that are covered:
>1. Run `go test -coverprofile=coverage.out ./...`
>2. Run `go tool cover -html=coverage.out -o coverage.html`
>3. Open the file `coverage.html` in any browser

APIs:  
A postman sample collection `go-project-apis.json` is available which can be imported directly to postman to test the APIs.  

1. Scan:
>Endpoint: /scan  
>Request Type: POST  
>Request URL Parameters: None  
>Request Body Type: JSON   
>Request Body Example: 
>``` 
>{  "repo":"https//raw.githubusercontent.com/velancio/vulnerability_scans/main/",
>  "files":["vulnscan18.json", "vulnscan456.json", "vulnscan123.json", "vulnscan789.json"]
>} 
>```
>Response:
>1. HTTP Status Code 200  
>Response Body: None
>2. HTTP Status Code 400 Bad Request  
>Reponse Body: 
>```
>{
>  "error": <error_message>
>}
>```  

2. Query:  
>Endpoint: /query  
>Request Type: POST  
>Request URL Parameters: None  
>Request Body Type: JSON   
>Request Body Example: 
>```
>{
>  "filters":{
>    "severity":"HIGH"
>  }
>}
>```
>Response:
>1. HTTP Status Code 200  
>Response Body: 
>```
>[
>  {
>    "id": "CVE-2024-7890",
>    "severity": "HIGH",
>    "cvss": 8.1,
>    "status": "active",
>    "package_name": "log4j",
>    "current_version": "2.14.0",
>    "fixed_version": "2.14.1",
>    "description": "Deserialization vulnerability in log4j logging framework",
>    "published_date": "2024-01-15T00:00:00Z",
>    "link": "",
>    "risk_factors": [
>      "Deserialization",
>      "Remote Code Execution"
>    ],
>    "scan_time": "2025-01-29T09:30:00Z",
>    "file_name": "vulnscan456.json"
>  },
>  {
>    "id": "CVE-2024-5567",
>    "severity": "HIGH",
>    "cvss": 8.2,
>    "status": "active",
>    "package_name": "grafana",
>    "current_version": "8.3.3",
>    "fixed_version": "8.3.4",
>    "description": "Authentication bypass in Grafana dashboards",
>    "published_date": "2024-01-28T00:00:00Z",
>    "link": "",
>    "risk_factors": [
>      "Authentication Bypass",
>      "High CVSS Score"
>    ],
>    "scan_time": "2025-01-29T12:00:00Z",
>    "file_name": "vulnscan456.json"
>  },
>  {
>    "id": "CVE-2024-1234",
>    "severity": "HIGH",
>    "cvss": 8.5,
>    "status": "fixed",
>    "package_name": "openssl",
>    "current_version": "1.1.1t-r0",
>    "fixed_version": "1.1.1u-r0",
>    "description": "Buffer overflow vulnerability in OpenSSL",
>    "published_date": "2024-01-15T00:00:00Z",
>    "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-1234",
>    "risk_factors": [
>      "Remote Code Execution",
>      "High CVSS Score",
>      "Public Exploit Available"
>    ],
>    "scan_time": "2025-01-28T10:30:00Z",
>    "file_name": "vulnscan18.json"
>  },
>  {
>    "id": "CVE-2024-6602",
>    "severity": "HIGH",
>    "cvss": 8.4,
>    "status": "active",
>    "package_name": "jackson-databind",
>    "current_version": "2.13.0",
>    "fixed_version": "2.13.1",
>    "description": "Remote code execution in Jackson databind",
>    "published_date": "2024-01-23T00:00:00Z",
>    "link": "https://nvd.nist.gov/vuln/detail/CVE-2024-6602",
>    "risk_factors": [
>      "Remote Code Execution",
>      "High CVSS Score",
>      "Commonly Used Component"
>    ],
>    "scan_time": "2025-01-28T14:45:00Z",
>    "file_name": "vulnscan18.json"
>  }
>]
>```
>2. HTTP Status Code 400 Bad Request  
>Reponse Body: 
>```
>{
>  "error": <error_message>
>}
>```