package models

import (
	"time"
	
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ReportedRecord struct
type ReportedRecord struct {
	gorm.Model  `json:"-"`
	ID          int64  `gorm:"column:Fid; primary_key:yes" json:"_id" `
	Date        time.Time `gorm:"not null" json:"date"`
	Hours       float32  `gorm:"not null" json:"hours"`
	Project     string `gorm:"not null" json:"project"`
	Description string `gorm:"not null" json:"description"`
	Rate        string `gorm:"not null" json:"rate"`
	Consultant  string `gorm:"not null" json:"consultant"`
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
