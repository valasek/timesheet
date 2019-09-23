// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package models

import (
	"github.com/valasek/timesheet/server/logger"

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

// NewHolidayManager - Create a Holiday manager that can be used for retrieving ReportedRecordss
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
	logger.Log.Error("unable to retrieve all holidays")
	return nil
}

// HolidaySeed - will load data from data file
func (db *HolidayManager) HolidaySeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err, " Input file: ", file)
	}
	defer csvfile.Close()

	holidayCSV := []*HolidayCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &holidayCSV); err != nil {
		logger.Log.Error(err)
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
	logger.Log.Error("unable to retrieve holidays count")
	return 0
}

// HolidayBackup will backup rates table
func (db *HolidayManager) HolidayBackup(filePath string) (int, error) {
	holidaysFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer holidaysFile.Close()

	holidays := []*Holiday{}
	db.db.Find(&holidays).Where("DeletedAt = ?", nil)
	holidayCSV := []*HolidayCSV{}
	for _, r := range holidays {
		createdAt := DateTime{r.CreatedAt}
		date := Date{r.Date}
		item := HolidayCSV{CreatedAt: createdAt, Date: date, Description: r.Description}
		holidayCSV = append(holidayCSV, &item)
	}

	err = gocsv.MarshalFile(&holidayCSV, holidaysFile)
	if err != nil {
		return 0, err
	}
	return len(holidays), nil
}

// HolidayGenerate generates test data
func (db *HolidayManager) HolidayGenerate(filePath string) (int, error) {
	holidaysFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer holidaysFile.Close()
	type h struct {
		Date        string
		Description string
	}
	hh := []h{
		{Date: "2019-01-01", Description: "New Year’s Day"},
		{Date: "2019-01-21", Description: "Birthday of Martin Luther King, Jr."},
		{Date: "2019-02-18", Description: "Washington’s Birthday"},
		{Date: "2019-05-27", Description: "Memorial Day"},
		{Date: "2019-07-04", Description: "Independence Day"},
		{Date: "2019-09-02", Description: "Labor Day"},
		{Date: "2019-10-14", Description: "Columbus Day"},
		{Date: "2019-11-11", Description: "Veterans Day"},
		{Date: "2019-11-28", Description: "Thanksgiving Day"},
		{Date: "2019-12-25", Description: "Christmas Day"},
	}
	var holidaysCSV []HolidayCSV
	d := DateTime{time.Now()}
	for _, v := range hh {
		dd, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return 0, err
		}
		holidaysCSV = append(holidaysCSV, HolidayCSV{CreatedAt: d, Date: Date{dd}, Description: v.Description})
	}

	err = gocsv.MarshalFile(&holidaysCSV, holidaysFile)
	if err != nil {
		return 0, err
	}
	return len(holidaysCSV), nil
}
