// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HolidaysGetAll returns list of all
func (api *API) HolidaysGetAll(c *gin.Context) {
	holidays := api.holidays.HolidaysGetAll()
	c.JSON(http.StatusOK, holidays)
}
