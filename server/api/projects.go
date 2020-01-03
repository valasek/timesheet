// Copyright © 2018-2020 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/models"
)

// ProjectsGetAll returns list of all projects
func (api *API) ProjectsGetAll(c *gin.Context) {
	projects := api.projects.ProjectsGetAll()
	c.JSON(http.StatusOK, projects)
}

// ProjectAdd add new record
func (api *API) ProjectAdd(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		logger.Log.Error("unable to decode project record, error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addedProject := api.projects.ProjectAdd(project)
	c.JSON(http.StatusOK, addedProject)
}

// ProjectUpdate add new record
func (api *API) ProjectUpdate(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		logger.Log.Error("unable to decode project record, error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProject := api.projects.ProjectUpdate(project)
	c.JSON(http.StatusOK, updatedProject)
}

// ProjectDelete deletes the project and all associated records
func (api *API) ProjectDelete(c *gin.Context) {
	ID := c.Param("id")
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logger.Log.Error("ProjectDelete, param 'id' is missing:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reportedRecords := api.projects.ProjectDelete(IDn)
	c.JSON(http.StatusOK, reportedRecords)
}

// ProjectToggle switch projects visibility
func (api *API) ProjectToggle(c *gin.Context) {
	ID := c.Param("id")
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		logger.Log.Error("ProjectToggle, param 'id' is missing:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project := api.projects.ProjectToggle(IDn)
	c.JSON(http.StatusOK, project)
}
