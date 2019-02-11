package models

import (
	"github.com/valasek/timesheet/server/logger"
	
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
	Rate      string     `gorm:"not null" json:"rate"`
	// add reference to a client
}

// ProjectCSV csv struct
type ProjectCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
	Rate      string   `csv:"rate"`
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
	logger.Log.Error("unable to retrieve all projects")
	return nil
}

// ProjectSeed - will load data from data file
func (db *ProjectManager) ProjectSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err)
	}
	defer csvfile.Close()

	projectsCSV := []*ProjectCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &projectsCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, p := range projectsCSV {
		newP := Project{CreatedAt: p.CreatedAt.Time, Name: p.Name, Rate: p.Rate}
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
	logger.Log.Error("unable to retrieve projects count")
	return 0
}

// ProjectBackup will backup rates table
func (db *ProjectManager) ProjectBackup(filePath string) (int, error) {
	projectsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer projectsFile.Close()

	projects := []*Project{}
	db.db.Find(&projects).Where("DeletedAt = ?", nil)
	projectCSV := []*ProjectCSV{}
	for _, r := range projects {
		createdAt := DateTime{r.CreatedAt}
		item := ProjectCSV{CreatedAt: createdAt, Name: r.Name, Rate: r.Rate}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, projectsFile)
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}