#!/bin/sh

# Initialize SQLite database if not exists
if [ ! -f "./projectDB" ]; then
    echo "Creating SQLite database..."
    sqlite3 projectDB "CREATE TABLE vulnerabilities (
    id TEXT PRIMARY KEY,
    severity TEXT,
    cvss REAL,
    status TEXT,
    package_name TEXT,
    current_version TEXT,
    fixed_version TEXT,
    description TEXT,
    published_date DATETIME,
    link TEXT,
    risk_factors TEXT,
    file_name TEXT,
    scan_time DATETIME
);
"
fi

# Run the Go application
./app