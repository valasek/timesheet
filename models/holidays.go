package models

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Holiday struct
type Holiday struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Date        time.Time  `gorm:"not null" json:"date"`
	Description string     `gorm:"not null" json:"description"`
	// add reference to a client
}

// HolidayCSV csv struct
type HolidayCSV struct {
	CreatedAt   DateTime `csv:"created_at"`
	Date        Date     `csv:"date"`
	Description string   `csv:"description"`
}

// HolidayManager struct
type HolidayManager struct {
	db *DB
}

// NewHolidayManager - Create a Project manager that can be used for retrieving ReportedRecordss
func NewHolidayManager(db *DB) (*HolidayManager, error) {

	db.AutoMigrate(&Holiday{})

	holidayMgr := HolidayManager{}

	holidayMgr.db = db

	return &holidayMgr, nil
}

// HolidaysGetAll - return all records of ReportedRecords
func (db *HolidayManager) HolidaysGetAll() []Holiday {
	holidays := []Holiday{}
	// Limit(100)
	if err := db.db.Find(&holidays); err != nil {
		return holidays
	}
	return nil
}

// HolidaySeed - will load data from data file
func (db *HolidayManager) HolidaySeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	holidayCSV := []*HolidayCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &holidayCSV); err != nil {
		fmt.Println(err)
	}
	for _, h := range holidayCSV {
		newH := Holiday{CreatedAt: h.CreatedAt.Time, Date: h.Date.Time, Description: h.Description}
		db.db.Create(&newH)
	}

	return len(holidayCSV)
}

// HolidayCount - return all records of Holidays
func (db *HolidayManager) HolidayCount() int {
	holidays := []Holiday{}
	var count int
	if err := db.db.Find(&holidays).Count(&count); err != nil {
		return count
	}
	return 0
}
