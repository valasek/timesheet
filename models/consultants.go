package models

import (
	"fmt"
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
	return nil
}

// ConsultantSeed - will load data from data file
func (db *ConsultantManager) ConsultantSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	consultantsCSV := []*ConsultantCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &consultantsCSV); err != nil {
		fmt.Println(err)
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
	return 0
}