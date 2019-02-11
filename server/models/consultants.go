package models

import (
	"github.com/valasek/timesheet/server/logger"
	
	"os"
	"time"

	"github.com/gocarina/gocsv"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Consultant struct
type Consultant struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `gorm:"not null;unique" json:"name"`
}

// ConsultantCSV csv struct
type ConsultantCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
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

// ConsultantSeed - will load data from data file
func (db *ConsultantManager) ConsultantSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err)
	}
	defer csvfile.Close()

	consultantsCSV := []*ConsultantCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &consultantsCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, c := range consultantsCSV {
		newC := Consultant{CreatedAt: c.CreatedAt.Time, Name: c.Name}
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
		item := ConsultantCSV{CreatedAt: createdAt, Name: r.Name}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, consultantsFile)
	if err != nil {
		return 0, err
	}
	return len(consultants), nil
}