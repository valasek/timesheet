package api

import (
	"fmt"

	"github.com/valasek/time-sheet/models"

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

// ReportedRecordDelete deletes records with given id
func (api *API) ReportedRecordDelete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]
	if len(ID) < 1 {
		fmt.Println("ReportedRecordsDelete, param 'id' is missing")
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsDelete(ID)
	json.NewEncoder(w).Encode(reportedRecords)
}

// ReportedRecordUpdate updates record
func (api *API) ReportedRecordUpdate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ID := vars["id"]
	if len(ID) < 1 {
		fmt.Println("ReportedRecordUpdate, param 'id' is missing")
		return
	}
	// https://www.programming-books.io/essential/go/2f90f3ba9eab401699dbb119feab6665-handling-http-method-accessing-query-strings-request-body

	//reportedRecord := api.reportedRecords.ReportedRecordUpdate( req.Body )
	// json.NewEncoder(w).Encode(reportedRecord)
}

// ReportedRecordsAddRecord add new record
func (api *API) ReportedRecordsAddRecord(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var reportedRecord models.ReportedRecord
	err := decoder.Decode(&reportedRecord)
	if err != nil {
		fmt.Println("unable to decode reported record, error: ", err)
		return
	}
	fmt.Println(reportedRecord)

	fmt.Println("ReportedRecordsAddRecord")
	fmt.Println(req.Body)

	id := "1234" 
	json.NewEncoder(w).Encode(id)
}