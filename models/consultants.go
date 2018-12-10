package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Consultant struct
type Consultant struct {
	gorm.Model `json:"-"`
	ID         int64  `gorm:"column:Fid; primary_key:yes" json:"_id" `
	Name       string `gorm:"not null;unique" json:"name"`
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
