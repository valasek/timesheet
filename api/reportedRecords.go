package api

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// ReportedRecordsGetAll returns list of all consultants
func (api *API) ReportedRecordsGetAll(w http.ResponseWriter, req *http.Request) {
	reportedRecords := api.reportedRecords.ReportedRecordsGetAll()
	json.NewEncoder(w).Encode(reportedRecords)
	//w.Write([]byte(consultants))
}

// ReportedRecordsInMonth returns list of all consultants
func (api *API) ReportedRecordsInMonth(w http.ResponseWriter, req *http.Request) {
	months, ok := req.URL.Query()["month"]
	if !ok || len(months[0]) < 1 {
		fmt.Println("url param 'month' is missing")
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsInMonth(months[0])
	json.NewEncoder(w).Encode(reportedRecords)
	//w.Write([]byte(consultants))
}