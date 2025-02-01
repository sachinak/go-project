package api

import (
	"go-project/db"
	"go-project/models"
	"go-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Scan(c *gin.Context) {
	var ReqBody models.ScanReq

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.GetFilesFromRemote(ReqBody.Repo, ReqBody.Files)
}

func Query(c *gin.Context) {
	var ReqBody models.QueryReq

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filters := ReqBody.Filters
	resp := db.GetVulnerabilitesFromDBWithFilters(filters)
	c.JSON(200, resp)
}
