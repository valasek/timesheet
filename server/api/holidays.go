// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
)

// HolidaysGetAll returns list of all
func (api *API) HolidaysGetAll(w http.ResponseWriter, req *http.Request) {
	holidays := api.holidays.HolidaysGetAll()
	json.NewEncoder(w).Encode(holidays)
}
