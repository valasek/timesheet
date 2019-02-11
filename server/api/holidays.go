package api

import (
	"net/http"
	"encoding/json"
)

// HolidaysGetAll returns list of all 
func (api *API) HolidaysGetAll(w http.ResponseWriter, req *http.Request) {
	holidays := api.holidays.HolidaysGetAll()
	json.NewEncoder(w).Encode(holidays)
}
