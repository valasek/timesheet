// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(c *gin.Context) {
	consultants := api.consultants.ConsultantList()
	c.JSON(http.StatusOK, consultants)
}
