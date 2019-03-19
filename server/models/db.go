// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package models

import (
	"os"
	"time"

	"github.com/valasek/timesheet/server/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // needed because of gorm design
	_ "github.com/jinzhu/gorm/dialects/postgres" // needed because of gorm design
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// UpdatedValue used to pass updated ReportedRecord value type and relevant ID
type UpdatedValue struct {
	ID    int64
	Type  string
	Value string
}

// DateTime used for csv marshalling and unmarshalling
type DateTime struct {
	time.Time
}

// Date used for csv marshalling and unmarshalling
type Date struct {
	time.Time
}

// MarshalCSV Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

// UnmarshalCSV Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

// MarshalCSV Convert the internal date as CSV string
func (date *Date) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

// UnmarshalCSV Convert the CSV string as internal date
func (date *Date) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	return err
}

// NewPostgresDB - postgres database
func NewPostgresDB(dataSourceName string) *DB {

	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		logger.Log.Error("cannot open postgres connection, error: ", err)
		os.Exit(1)
	}

	if err = db.DB().Ping(); err != nil {
		logger.Log.Error("cannot ping postgres, error: ", err)
		os.Exit(1)
	}

	// db.LogMode(true)

	return &DB{db}
}

// NewMySQLDB - mysql database
func NewMySQLDB(dataSourceName string) *DB {

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		logger.Log.Error("cannot open mysql connection, error: ", err)
		os.Exit(1)
	}

	if err = db.DB().Ping(); err != nil {
		// panic("ping", err)
		logger.Log.Error("cannot ping mysql, error: ", err)
		os.Exit(1)
	}

	// db.LogMode(true)

	return &DB{db}
}
