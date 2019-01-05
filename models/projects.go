package models

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Project struct
type Project struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `gorm:"not null" json:"name"`
	// add reference to a client
}

// ProjectCSV csv struct
type ProjectCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
}

// ProjectManager struct
type ProjectManager struct {
	db *DB
}

// NewProjectManager - Create a Project manager that can be used for retrieving ReportedRecordss
func NewProjectManager(db *DB) (*ProjectManager, error) {

	db.AutoMigrate(&Project{})

	projectMgr := ProjectManager{}

	projectMgr.db = db

	return &projectMgr, nil
}

// ProjectsGetAll - return all records of ReportedRecords
func (db *ProjectManager) ProjectsGetAll() []Project {
	projects := []Project{}
	// Limit(100)
	if err := db.db.Find(&projects); err != nil {
		return projects
	}
	return nil
}

// ProjectSeed - will load data from data file
func (db *ProjectManager) ProjectSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	projectsCSV := []*ConsultantCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &projectsCSV); err != nil {
		fmt.Println(err)
	}
	for _, p := range projectsCSV {
		newP := Project{CreatedAt: p.CreatedAt.Time, Name: p.Name}
		db.db.Create(&newP)
	}

	return len(projectsCSV)
}

// ProjectCount - return all records of Rates
func (db *ProjectManager) ProjectCount() int {
	projects := []Project{}
	var count int
	if err := db.db.Find(&projects).Count(&count); err != nil {
		return count
	}
	return 0
}
