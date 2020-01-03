// Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/models"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(c *gin.Context) {
	consultants := api.consultants.ConsultantList()
	c.JSON(http.StatusOK, consultants)
}

// ConsultantAdd add new record
func (api *API) ConsultantAdd(c *gin.Context) {
	var consultant models.Consultant
	if err := c.ShouldBindJSON(&consultant); err != nil {
		logger.Log.Error("unable to decode consultant record, error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedConsultant := api.consultants.ConsultantAdd(consultant)
	c.JSON(http.StatusOK, addedConsultant)
}

// ConsultantUpdate add new record
func (api *API) ConsultantUpdate(c *gin.Context) {
	var consultant models.Consultant
	if err := c.ShouldBindJSON(&consultant); err != nil {
		logger.Log.Error("unable to decode consultant record, error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedConsultant := api.consultants.ConsultantUpdate(consultant)
	c.JSON(http.StatusOK, updatedConsultant)
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
