// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valasek/timesheet/server/logger"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(c *gin.Context) {
	consultants := api.consultants.ConsultantList()
	c.JSON(http.StatusOK, consultants)
}

// ConsultantDelete deletes the consultant and all associated records
func (api *API) ConsultantDelete(c *gin.Context) {
	ID := c.Param("id")
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logger.Log.Error("ConsultantDelete, param 'id' is missing:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reportedRecords := api.consultants.ConsultantDelete(IDn)
	c.JSON(http.StatusOK, reportedRecords)
}

// ConsultantToggle switch projects visibility
func (api *API) ConsultantToggle(c *gin.Context) {
	ID := c.Param("id")
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logger.Log.Error("ConsultantToggle, param 'id' is missing:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project := api.consultants.ConsultantToggle(IDn)
	c.JSON(http.StatusOK, project)
}
