package models

import (
	"fmt"
	"os"
	"time"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/jinzhu/gorm"

	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ReportedRecord struct
type ReportedRecord struct {
	gorm.Model  `json:"-"`
	ID          int64     `gorm:"column:Fid; primary_key:yes" json:"_id"`
	Date        time.Time `gorm:"not null" json:"date"`
	Hours       float32   `gorm:"not null" json:"hours"`
	Project     string    `gorm:"not null" json:"project"`
	Description string    `gorm:"not null" json:"description"`
	Rate        string    `gorm:"not null" json:"rate"`
	Consultant  string    `gorm:"not null" json:"consultant"`
}

// ReportedRecordCSV csv struct
type ReportedRecordCSV struct {
	ID          uint     `csv:"id"`
	CreatedAt   DateTime `csv:"created_at"`
	Date        Date     `csv:"date"`
	Hours       float32  `csv:"hours"`
	Project     string   `csv:"project"`
	Description string   `csv:"description"`
	Rate        string   `csv:"rate"`
	Consultant  string   `csv:"consultant"`
}

// ReportedRecordManager struct
type ReportedRecordManager struct {
	db *DB
}

// NewReportedRecordManager - Create a ReportedRecords manager that can be used for retrieving ReportedRecordss
func NewReportedRecordManager(db *DB) (*ReportedRecordManager, error) {

	db.AutoMigrate(&ReportedRecord{})

	reportedRecordsMgr := ReportedRecordManager{}

	reportedRecordsMgr.db = db

	return &reportedRecordsMgr, nil
}

// ReportedRecordsGetAll - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsGetAll() []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Find(&reportedRecords); err != nil {
		return reportedRecords
	}
	return nil
}

// ReportedRecordsInMonth - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsInMonth(month string) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("extract(MONTH from date) = ?", month).Find(&reportedRecords); err != nil {
		return reportedRecords
	}
	return nil
}

// ReportedRecordsDelete - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsDelete(id string) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("id = ?", id).Delete(&reportedRecords); err != nil {
		return reportedRecords
	}
	return nil
}

// ReportedRecordUpdate -
func (db *ReportedRecordManager) ReportedRecordUpdate(r UpdatedValue) ReportedRecord {
	updateValue := UpdatedValue{
		ID:        r.ID,
		Type:       r.Type,
		Value:     r.Value,
	}
	reportedRecord := ReportedRecord{}
	// handle attribute types
	switch r.Type {
	case "hours":
		value, err := strconv.ParseFloat(updateValue.Value, 64)
		if err != nil {
			fmt.Println(err)
		}
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, value); err != nil {
			return reportedRecord
		}
	case "description", "rate", "project":
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, updateValue.Value); err != nil {
			return reportedRecord
		}
	case "date":
		value, err := time.Parse("2006-01-02", updateValue.Value)
		if err != nil {
			fmt.Println(err)
		}
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, value); err != nil {
			return reportedRecord
		}
	default:
		fmt.Println("unknown attribute type: ", r.Type)
		return ReportedRecord{}
	}
	return ReportedRecord{}
}

// ReportedRecordSeed - will load data from data file
func (db *ReportedRecordManager) ReportedRecordSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	recordsCSV := []*ReportedRecordCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &recordsCSV); err != nil {
		fmt.Println(err)
	}
	for _, r := range recordsCSV {
		newR := ReportedRecord{gorm.Model{ID: r.ID, CreatedAt: r.CreatedAt.Time}, int64(r.ID), r.Date.Time, r.Hours, r.Project, r.Description, r.Rate, r.Consultant}
		db.db.Create(&newR)
	}

	return len(recordsCSV)
}

// ReportedRecordCount - 
func (db *ReportedRecordManager) ReportedRecordCount() (int) {
	reportedRecords := []ReportedRecord{}
	var count int
	if err := db.db.Find(&reportedRecords).Count(&count); err != nil {
		return count
	}
	return 0
}