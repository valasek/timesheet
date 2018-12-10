package api

import (
	"net/http"
	"encoding/json"
)

// ReportedRecordsGet returns list of all consultants
func (api *API) ReportedRecordsGet(w http.ResponseWriter, req *http.Request) {
	reportedRecords := api.reportedRecords.ReportedRecordGet()
	json.NewEncoder(w).Encode(reportedRecords)
	//w.Write([]byte(consultants))
}
