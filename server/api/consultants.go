// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
)

// ConsultantList returns list of all consultants
func (api *API) ConsultantList(w http.ResponseWriter, req *http.Request) {
	consultants := api.consultants.ConsultantList()
	json.NewEncoder(w).Encode(consultants)
	//w.Write([]byte(consultants))
}
