package api

import (
	"encoding/json"
	"net/http"
)

// RatesGetAll returns list of all rates
func (api *API) RatesGetAll(w http.ResponseWriter, req *http.Request) {
	rates := api.rates.RatesGetAll()
	json.NewEncoder(w).Encode(rates)
}
