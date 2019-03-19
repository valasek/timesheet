// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RatesGetAll returns list of all rates
func (api *API) RatesGetAll(c *gin.Context) {
	rates := api.rates.RatesGetAll()
	c.JSON(http.StatusOK, rates)
}
