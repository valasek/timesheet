package api

import (
	"github.com/valasek/time-sheet/models"
)

// API -
type API struct {
	users  *models.UserManager
	quotes *models.QuoteManager
	consultants *models.ConsultantManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserManager(db)
	quotemgr, _ := models.NewQuoteManager(db)
	consultantmgr, _ := models.NewConsultantManager(db)

	return &API{
		users:  usermgr,
		quotes: quotemgr,
		consultants: consultantmgr,
	}
}
