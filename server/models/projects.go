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

// Project struct
type Project struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `gorm:"not null" json:"name"`
	Rate      string     `gorm:"not null" json:"rate"`
	Disabled  bool       `gorm:"not null" json:"disabled"`
}

// ProjectCSV csv struct
type ProjectCSV struct {
	CreatedAt DateTime `csv:"created_at"`
	Name      string   `csv:"name"`
	Rate      string   `csv:"rate"`
	Disabled  bool     `csv:"disabled"`
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

// ProjectToggle - return all records of Rates
func (db *ProjectManager) ProjectToggle(id uint64) Project {
	project := Project{}
	if err := db.db.First(&project, id); err != nil {
		project.Disabled = !project.Disabled
		db.db.Save(&project)
		return project
	}
	logger.Log.Error("unable to toggle project id", id)
	return project
}

// ProjectAdd -
func (db *ProjectManager) ProjectAdd(newRecord Project) Project {
	if err := db.db.Create(&newRecord); err != nil {
		return newRecord
	}
	logger.Log.Error("unable to add new project", newRecord)
	return Project{}
}

// ProjectDelete - deletes the project and all associated records
func (db *ProjectManager) ProjectDelete(id uint64) int64 {

	project := Project{}
	if err := db.db.First(&project, id); err != nil {
		name := project.Name
		if err := db.db.Delete(&project); err != nil {
			var reportedRecordCount int64
			if err := db.db.Model(ReportedRecord{}).Where("project like ?", name).Count(&reportedRecordCount); err != nil {
				if err := db.db.Where("project like ?", name).Delete(&ReportedRecord{}); err != nil {
					return reportedRecordCount
				}
			}
		}
	}
	logger.Log.Error("unable to delete project id", id)
	return 0
}

// ProjectsGetStatistics - returns table statistics
func (db *ProjectManager) ProjectsGetStatistics() EntityOverview {
	table := "projects"
	var total, active, disabled int
	if err := db.db.Unscoped().Table(table).Count(&total); err != nil {
		active = db.ProjectCount()
	} else {
		logger.Log.Error("failed to retrieve from DB statistics (total) for table ", table)
	}
	if err := db.db.Table(table).Where("disabled = true").Count(&disabled); err == nil {
		logger.Log.Error("failed to retrieve from DB statistics (disabled) for table ", table)
	}
	return EntityOverview{Name: strings.Title(table), Total: total, Active: active, Disabled: disabled, Deleted: total - active}
}

// ProjectSeed - will load data from data file
func (db *ProjectManager) ProjectSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err, " Input file: ", file)
	}
	defer csvfile.Close()

	projectsCSV := []*ProjectCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &projectsCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, p := range projectsCSV {
		newP := Project{CreatedAt: p.CreatedAt.Time, Name: p.Name, Rate: p.Rate, Disabled: p.Disabled}
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
		item := ProjectCSV{CreatedAt: createdAt, Name: r.Name, Rate: r.Rate, Disabled: r.Disabled}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, projectsFile)
	if err != nil {
		return 0, err
	}
	return len(projects), nil
}
