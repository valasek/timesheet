package api

import (
	"fmt"
	"strconv"

	"github.com/valasek/timesheet/models"

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
	IDn, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		fmt.Println("ReportedRecordsDelete, param 'id' is missing:", err)
		return
	}
	reportedRecords := api.reportedRecords.ReportedRecordsDelete(IDn)
	json.NewEncoder(w).Encode(reportedRecords)
}

// ReportedRecordUpdate updates record
func (api *API) ReportedRecordUpdate(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var updatedValue models.UpdatedValue
	err := decoder.Decode(&updatedValue)
	fmt.Println(updatedValue)
	if err != nil {
		fmt.Printf("unable to decode reported record %s, error: %s", req.Body, err)
		return
	}
	// https://www.programming-books.io/essential/go/2f90f3ba9eab401699dbb119feab6665-handling-http-method-accessing-query-strings-request-body
	reportedRecord := api.reportedRecords.ReportedRecordUpdate( updatedValue )
	json.NewEncoder(w).Encode(reportedRecord)
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
	addedReportedRecord := api.reportedRecords.ReportedRecordAdd( reportedRecord )
	json.NewEncoder(w).Encode(addedReportedRecord)
}