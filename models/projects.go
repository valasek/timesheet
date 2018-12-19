package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Project struct
type Project struct {
	gorm.Model  `json:"-"`
	ID          int64  `gorm:"column:Fid; primary_key:yes" json:"_id" `
	Name     string `gorm:"not null" json:"name"`
	// add reference to a client
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
