package models

import (
	"fmt"
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
	// add reference to a client
}

// RateCSV csv struct
type RateCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
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
	return nil
}

// RateCount - return all records of Rates
func (db *RateManager) RateCount() int {
	rates := []Rate{}
	var count int
	if err := db.db.Find(&rates).Count(&count); err != nil {
		return count
	}
	return 0
}

// RateSeed - will load data from data file
func (db *RateManager) RateSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	ratesCSV := []*ConsultantCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &ratesCSV); err != nil {
		fmt.Println(err)
	}
	for _, r := range ratesCSV {
		newR := Rate{CreatedAt: r.CreatedAt.Time, Name: r.Name}
		db.db.Create(&newR)
	}

	return len(ratesCSV)
}
