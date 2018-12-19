package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// ReportedRecordsGetAll returns list of all 
func (api *API) ReportedRecordsGetAll(w http.ResponseWriter, req *http.Request) {
	reportedRecords := api.reportedRecords.ReportedRecordsGetAll()
	json.NewEncoder(w).Encode(reportedRecords)
}

// ReportedRecordsInMonth returns list of all 
func (api *API) ReportedRecordsInMonth(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	month := vars["month"]
	if len(month) < 1 {
		fmt.Println("ReportedRecordsInMonth, param 'month' is missing")
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsInMonth(month)
	json.NewEncoder(w).Encode(reportedRecords)
}

// ReportedRecordsDelete deletes records with given id
func (api *API) ReportedRecordsDelete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]
	if len(ID) < 1 {
		fmt.Println("ReportedRecordsDelete, param 'id' is missing")
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsDelete(ID)
	json.NewEncoder(w).Encode(reportedRecords)
}
