// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(c *gin.Context) {
	consultants := api.consultants.ConsultantList()
	c.JSON(http.StatusOK, consultants)
}
