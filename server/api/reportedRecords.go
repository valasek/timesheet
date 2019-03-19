// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/models"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ReportedRecordsGetAll returns list of all
func (api *API) ReportedRecordsGetAll(c *gin.Context) {
	reportedRecords := api.reportedRecords.ReportedRecordsGetAll()
	c.JSON(http.StatusOK, reportedRecords)
}

// ReportedRecordsInMonth returns list of all
func (api *API) ReportedRecordsInMonth(c *gin.Context) {
	month := c.Param("month")
	if len(month) < 1 {
		logger.Log.Error("ReportedRecordsInMonth, param 'month' is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReportedRecordsInMonth, param 'month' is missing"})
		return
	}
	year := c.Param("year")
	if len(year) < 1 {
		logger.Log.Error("ReportedRecordsInMonth, param 'year' is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReportedRecordsInMonth, param 'year' is missing"})
		return
	}
	// consultant := vars["consultant"]
	consultant := c.Param("consultant")
	if len(consultant) < 1 {
		logger.Log.Error("ReportedRecordsInMonth, param 'consultant' is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReportedRecordsInMonth, param 'consultant' is missing"})
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsInMonth(year, month, consultant)
	c.JSON(http.StatusOK, reportedRecords)
}

// ReportedRecordsSummary returns list of all
func (api *API) ReportedRecordsSummary(c *gin.Context) {
	year := c.Param("year")
	if len(year) < 1 {
		logger.Log.Error("ReportedRecordsInMonth, param 'year' is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReportedRecordsInMonth, param 'year' is missing"})
		return
	}
	reportedRecordsSummary := api.reportedRecords.ReportedRecordsSummary(year)
	c.JSON(http.StatusOK, reportedRecordsSummary)
}

// ReportedRecordDelete deletes records with given id
func (api *API) ReportedRecordDelete(c *gin.Context) {

	ID := c.Param("id")
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logger.Log.Error("ReportedRecordsDelete, param 'id' is missing:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsDelete(IDn)
	c.JSON(http.StatusOK, reportedRecords)
}

// ReportedRecordUpdate updates record
func (api *API) ReportedRecordUpdate(c *gin.Context) {
	var updatedValue models.UpdatedValue
	if err := c.ShouldBindJSON(&updatedValue); err != nil {
		logger.Log.Error(fmt.Sprintf("unable to decode reported record, error: %s", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reportedRecord := api.reportedRecords.ReportedRecordUpdate(updatedValue)
	c.JSON(http.StatusOK, reportedRecord)
}

// ReportedRecordsAddRecord add new record
func (api *API) ReportedRecordsAddRecord(c *gin.Context) {
	// decoder := json.NewDecoder(req.Body)
	var reportedRecord models.ReportedRecord
	if err := c.ShouldBindJSON(&reportedRecord); err != nil {
		logger.Log.Error("unable to decode reported record, error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedReportedRecord := api.reportedRecords.ReportedRecordAdd(reportedRecord)
	c.JSON(http.StatusOK, addedReportedRecord)
}
