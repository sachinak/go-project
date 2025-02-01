package utils

import (
	"encoding/json"
	"fmt"
	"go-project/db"
	"go-project/models"
	"io"
	"net/http"
	"sync"
)

func GetFilesFromRemote(url string, files []string) {
	filenamePrefix := url

	if len(files) > 3 {
		var wg sync.WaitGroup
		for _, file := range files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				GetFilesFromRemoteHelper(filenamePrefix, file)
			}(file)
		}
		wg.Wait()
	} else {
		for _, file := range files {
			GetFilesFromRemoteHelper(filenamePrefix, file)
		}
	}
}

func GetFilesFromRemoteHelper(filenamePrefix string, file string) {
	fileUrl := filenamePrefix + file

	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("error in fetching file from github:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in reading request body:", err)
		return
	}

	var scanResults []models.ScanResp2
	err = json.Unmarshal(body, &scanResults)
	if err != nil {
		fmt.Println("error in unmarshalling:", err)
		return
	}

	for _, data := range scanResults {
		for _, vulnerability := range data.ScanRes.Vulnerabilities {
			db.WritePayloadToDB(vulnerability, data.ScanRes.Timestamp, file)
		}
	}
}
