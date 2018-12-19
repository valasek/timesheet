package api

import (
	"net/http"
	"encoding/json"
)

// RatesGetAll returns list of all rates
func (api *API) RatesGetAll(w http.ResponseWriter, req *http.Request) {
	rates := api.rates.RatesGetAll()
	json.NewEncoder(w).Encode(rates)
}
