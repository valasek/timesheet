// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// RatesGetAll returns list of all rates
func (api *API) RatesGetAll(c *gin.Context) {
	rates := api.rates.RatesGetAll()
	c.JSON(http.StatusOK, rates)
}
