package models

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
	// "github.com/jinzhu/gorm"

	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ReportedRecord struct
type ReportedRecord struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Date        time.Time  `gorm:"not null" json:"date"`
	Hours       float64    `gorm:"not null" json:"hours"`
	Project     string     `gorm:"not null" json:"project"`
	Description string     `gorm:"not null" json:"description"`
	Rate        string     `gorm:"not null" json:"rate"`
	Consultant  string     `gorm:"not null" json:"consultant"`
}

// ReportedRecordCSV csv struct
type ReportedRecordCSV struct {
	CreatedAt   DateTime `csv:"created_at"`
	Date        Date     `csv:"date"`
	Hours       float64  `csv:"hours"`
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
	fmt.Println("failed - get all reported records")
	return nil
}

// ReportedRecordsInMonth - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsInMonth(month string) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("extract(MONTH from date) = ?", month).Find(&reportedRecords); err != nil {
		return reportedRecords
	}
	fmt.Println("failed - get reported records in month", month)
	return nil
}

// ReportedRecordsDelete - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsDelete(id uint64) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("id = ?", id).Delete(&reportedRecords); err != nil {
		return reportedRecords
	}
	fmt.Println("failed - delete reported record for id", id)
	return nil
}

// ReportedRecordAdd -
func (db *ReportedRecordManager) ReportedRecordAdd(newRecord ReportedRecord) ReportedRecord {
	fmt.Println("entering ReportedRecordAdd")
	if err := db.db.Create(&newRecord); err != nil {
		return newRecord
	}
	fmt.Println("failed - add new reported record")
	return ReportedRecord{}
}

// ReportedRecordUpdate -
func (db *ReportedRecordManager) ReportedRecordUpdate(r UpdatedValue) ReportedRecord {
	updateValue := UpdatedValue{
		ID:    r.ID,
		Type:  r.Type,
		Value: r.Value,
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
		fmt.Println("failed - updating reported record, unknown attribute type: ", r.Type, updateValue)
		return ReportedRecord{}
	}
	fmt.Println("failed - updating reported record", updateValue)
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
		newR := ReportedRecord{CreatedAt: r.CreatedAt.Time, Date: r.Date.Time, Hours: r.Hours, Project: r.Project, Description: r.Description, Rate: r.Rate, Consultant: r.Consultant}
		db.db.Create(&newR)
	}

	return len(recordsCSV)
}

// ReportedRecordCount -
func (db *ReportedRecordManager) ReportedRecordCount() int {
	reportedRecords := []ReportedRecord{}
	var count int
	if err := db.db.Find(&reportedRecords).Count(&count); err != nil {
		return count
	}
	fmt.Println("failed - getting reported record count")
	return 0
}

// ReportedRecordBackup will backup rates table
func (db *ReportedRecordManager) ReportedRecordBackup(filePath string) (int, error) {
	reportedRecordsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer reportedRecordsFile.Close()

	reportedRecords := []*ReportedRecord{}
	db.db.Find(&reportedRecords).Where("DeletedAt = ?", nil)
	projectCSV := []*ReportedRecordCSV{}
	for _, r := range reportedRecords {
		createdAt := DateTime{r.CreatedAt}
		date := Date{r.Date}
		item := ReportedRecordCSV{CreatedAt: createdAt, Date: date, Hours: r.Hours, Project: r.Project, Description: r.Description, Consultant: r.Consultant}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, reportedRecordsFile)
	if err != nil {
		return 0, err
	}
	return len(reportedRecords), nil
}