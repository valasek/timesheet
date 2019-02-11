package api

import (
	"net/http"
	"encoding/json"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(w http.ResponseWriter, req *http.Request) {
	consultants := api.consultants.ConsultantList()
	json.NewEncoder(w).Encode(consultants)
	//w.Write([]byte(consultants))
}
