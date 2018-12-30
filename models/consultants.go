package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/gocarina/gocsv"


	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Consultant struct
type Consultant struct {
	gorm.Model `json:"-"`
	ID         int64  `gorm:"column:Fid; primary_key:yes" json:"_id" `
	Name       string `gorm:"not null;unique" json:"name"`
}

// ConsultantCSV csv struct
type ConsultantCSV struct {
	ID        uint     `csv:"id`
	CreatedAt DateTime `csv:"created_at"`
	Name      string    `csv:"name"`
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
		newC := Consultant{gorm.Model{ID: c.ID, CreatedAt: c.CreatedAt.Time}, int64(c.ID), c.Name}
		db.db.Create(&newC)
	}

	return len(consultantsCSV)
}

// ConsultantsCount - 
func (db *ConsultantManager) ConsultantCount() (int) {
	consultants := []Consultant{}
	var count int
	if err := db.db.Find(&consultants).Count(&count); err != nil {
		return count
	}
	return 0
}