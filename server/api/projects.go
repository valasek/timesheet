// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ProjectsGetAll returns list of all projects
func (api *API) ProjectsGetAll(c *gin.Context) {
	projects := api.projects.ProjectsGetAll()
	c.JSON(http.StatusOK, projects)
}
