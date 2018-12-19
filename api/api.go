package api

import (
	"github.com/valasek/time-sheet/models"
)

// API -
type API struct {
	users  *models.UserManager
	quotes *models.QuoteManager
	consultants *models.ConsultantManager
	reportedRecords *models.ReportedRecordManager
	projects *models.ProjectManager
	rates *models.RateManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	quotemgr, _ := models.NewQuoteManager(db)
	consultantmgr, _ := models.NewConsultantManager(db)
	reportedrecordsmgr, _ := models.NewReportedRecordManager(db)
	projectsmgr, _ := models.NewProjectManager(db)
	ratesmgr, _ := models.NewRateManager(db)

	return &API{
		users:  usermgr,
		quotes: quotemgr,
		consultants: consultantmgr,
		reportedRecords: reportedrecordsmgr,
		projects: projectsmgr,
		rates: ratesmgr,
	}
}
