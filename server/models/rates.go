// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package models

import (
	"strings"

	"github.com/valasek/timesheet/server/logger"

	"os"
	"time"

	"github.com/gocarina/gocsv"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Rate struct
type Rate struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `gorm:"not null" json:"name"`
	Type      string     `gorm:"not null" json:"type"`
	// add reference to a client
}

// RateCSV csv struct
type RateCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
	Type      string   `csv:"type"`
}

// RateManager struct
type RateManager struct {
	db *DB
}

// NewRateManager - Create a Rate manager that can be used for retrieving ReportedRecordss
func NewRateManager(db *DB) (*RateManager, error) {

	db.AutoMigrate(&Rate{})

	rateMgr := RateManager{}

	rateMgr.db = db

	return &rateMgr, nil
}

// RatesGetAll - return all records of Rates
func (db *RateManager) RatesGetAll() []Rate {
	rates := []Rate{}
	// Limit(100)
	if err := db.db.Find(&rates); err != nil {
		return rates
	}
	logger.Log.Error("unable to retrieve all rates")
	return nil
}

// RatesGetStatistics - returns table statistics
func (db *RateManager) RatesGetStatistics() EntityOverview {
	table := "rates"
	var total, active int
	if err := db.db.Unscoped().Table("rates").Count(&total); err != nil {
		active = db.RateCount()
	} else {
		logger.Log.Error("failed to retrieve from DB statistics for table ", table)
	}
	return EntityOverview{Name: strings.Title(table), Total: total, Active: active, Disabled: 0, Deleted: total - active}
}

// RateCount - return all records of Rates
func (db *RateManager) RateCount() int {
	rates := []Rate{}
	var count int
	if err := db.db.Find(&rates).Count(&count); err != nil {
		return count
	}
	logger.Log.Error("unable to retrieve rates count")
	return 0
}

// RateSeed - will load data from data file
func (db *RateManager) RateSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err, " Input file: ", file)
	}
	defer csvfile.Close()

	ratesCSV := []*RateCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &ratesCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, r := range ratesCSV {
		newR := Rate{CreatedAt: r.CreatedAt.Time, Name: r.Name, Type: r.Type}
		db.db.Create(&newR)
	}

	return len(ratesCSV)
}

// RateBackup will backup rates table
func (db *RateManager) RateBackup(filePath string) (int, error) {
	ratesFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer ratesFile.Close()

	rates := []*Rate{}
	db.db.Find(&rates).Where("DeletedAt = ?", nil)
	ratesCSV := []*RateCSV{}
	for _, r := range rates {
		createdAt := DateTime{r.CreatedAt}
		item := RateCSV{CreatedAt: createdAt, Name: r.Name, Type: r.Type}
		ratesCSV = append(ratesCSV, &item)
	}

	err = gocsv.MarshalFile(&ratesCSV, ratesFile)
	if err != nil {
		return 0, err
	}
	return len(rates), nil
}

// RateGenerate generates test data
func (db *RateManager) RateGenerate(filePath string) (int, error) {
	ratesFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer ratesFile.Close()

	d := DateTime{time.Now()}
	ratesCSV := []RateCSV{
		{CreatedAt: d, Name: "On-site", Type: "work"},
		{CreatedAt: d, Name: "Standard", Type: "work"},
		{CreatedAt: d, Name: "Standard Weekend", Type: "work"},
		{CreatedAt: d, Name: "Standard Holiday", Type: "work"},
		{CreatedAt: d, Name: "Vacation", Type: "not-work"},
		{CreatedAt: d, Name: "Sick Day", Type: "not-work"},
		{CreatedAt: d, Name: "Personal Day", Type: "not-work"},
		{CreatedAt: d, Name: "Unpaid Leave", Type: "not-work"},
		{CreatedAt: d, Name: "Sick", Type: "not-work"},
	}

	err = gocsv.MarshalFile(&ratesCSV, ratesFile)
	if err != nil {
		return 0, err
	}
	return len(ratesCSV), nil
}
