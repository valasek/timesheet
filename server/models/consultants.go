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

// Consultant struct
type Consultant struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
	Name       string     `gorm:"not null;unique" json:"name"`
	Allocation float64    `gorm:"not null" json:"allocation"`
	Disabled   bool       `gorm:"not null" json:"disabled"`
}

// ConsultantCSV csv struct
type ConsultantCSV struct {
	CreatedAt  DateTime `csv:"created_at"`
	Name       string   `csv:"name"`
	Allocation float64  `csv:"allocation"`
	Disabled   bool     `csv:"disabled"`
}

// ConsultantManager struct
type ConsultantManager struct {
	db *DB
}

// NewConsultantManager - Create a consultant manager that can be used for retrieving consultants
func NewConsultantManager(db *DB) (*ConsultantManager, error) {

	db.AutoMigrate(&Consultant{})

	consultantmgr := ConsultantManager{}

	consultantmgr.db = db

	return &consultantmgr, nil
}

// ConsultantList - return a list of Consultants
func (db *ConsultantManager) ConsultantList() []Consultant {
	consultants := []Consultant{}
	if err := db.db.Find(&consultants); err != nil {
		return consultants
	}
	logger.Log.Error("unable to retrieve all consultants")
	return nil
}

// ConsultantAdd -
func (db *ConsultantManager) ConsultantAdd(newRecord Consultant) Consultant {
	if err := db.db.Create(&newRecord); err != nil {
		return newRecord
	}
	logger.Log.Error("unable to add new consultant", newRecord)
	return Consultant{}
}

// ConsultantDelete - deletes the consultant and all associated records
func (db *ConsultantManager) ConsultantDelete(id uint64) int64 {

	consultant := Consultant{}
	if err := db.db.First(&consultant, id); err != nil {
		name := consultant.Name
		if err := db.db.Delete(&consultant); err != nil {
			var reportedRecordCount int64
			if err := db.db.Model(ReportedRecord{}).Where("consultant like ?", name).Count(&reportedRecordCount); err != nil {
				if err := db.db.Where("consultant like ?", name).Delete(&ReportedRecord{}); err != nil {
					return reportedRecordCount
				}
			}
		}
	}
	logger.Log.Error("unable to delete consultant id", id)
	return 0
}

// ConsultantsGetStatistics - returns table statistics
func (db *ConsultantManager) ConsultantsGetStatistics() EntityOverview {
	table := "consultants"
	var total, active, disabled int
	if err := db.db.Unscoped().Table(table).Count(&total); err != nil {
		active = db.ConsultantCount()
	} else {
		logger.Log.Error("failed to retrieve from DB statistics for table ", table)
	}
	if err := db.db.Table(table).Where("disabled = true").Count(&disabled); err == nil {
		logger.Log.Error("failed to retrieve from DB statistics (disabled) for table ", table)
	}
	return EntityOverview{Name: strings.Title(table), Total: total, Active: active, Disabled: disabled, Deleted: total - active}
}

// ConsultantToggle - return all records of Rates
func (db *ConsultantManager) ConsultantToggle(id uint64) Consultant {
	consultant := Consultant{}
	if err := db.db.First(&consultant, id); err != nil {
		consultant.Disabled = !consultant.Disabled
		db.db.Save(&consultant)
		return consultant
	}
	logger.Log.Error("unable to toggle consultant id", id)
	return consultant
}

// ConsultantSeed - will load data from data file
func (db *ConsultantManager) ConsultantSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err, " Input file: ", file)
	}
	defer csvfile.Close()

	consultantsCSV := []*ConsultantCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &consultantsCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, c := range consultantsCSV {
		newC := Consultant{CreatedAt: c.CreatedAt.Time, Name: c.Name, Allocation: c.Allocation, Disabled: c.Disabled}
		db.db.Create(&newC)
	}

	return len(consultantsCSV)
}

// ConsultantCount -
func (db *ConsultantManager) ConsultantCount() int {
	consultants := []Consultant{}
	var count int
	if err := db.db.Find(&consultants).Count(&count); err != nil {
		return count
	}
	logger.Log.Error("unable to retrieve consultants count")
	return 0
}

// ConsultantBackup will backup rates table
func (db *ConsultantManager) ConsultantBackup(filePath string) (int, error) {
	consultantsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer consultantsFile.Close()

	consultants := []*Consultant{}
	db.db.Find(&consultants).Where("DeletedAt = ?", nil)
	projectCSV := []*ConsultantCSV{}
	for _, r := range consultants {
		createdAt := DateTime{r.CreatedAt}
		item := ConsultantCSV{CreatedAt: createdAt, Name: r.Name, Allocation: r.Allocation, Disabled: r.Disabled}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, consultantsFile)
	if err != nil {
		return 0, err
	}
	return len(consultants), nil
}
