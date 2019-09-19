// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valasek/timesheet/server/logger"
)

// ProjectsGetAll returns list of all projects
func (api *API) ProjectsGetAll(c *gin.Context) {
	projects := api.projects.ProjectsGetAll()
	c.JSON(http.StatusOK, projects)
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
