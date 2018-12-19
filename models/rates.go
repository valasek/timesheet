package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Project struct
type Rate struct {
	gorm.Model  `json:"-"`
	ID          int64  `gorm:"column:Fid; primary_key:yes" json:"_id" `
	Name     string `gorm:"not null" json:"name"`
	// add reference to a client
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
