package api

import (
	"fmt"
	"github.com/valasek/time-sheet/models"
	"github.com/spf13/viper"
)

// API -
type API struct {
	users  *models.UserManager
	consultants *models.ConsultantManager
	reportedRecords *models.ReportedRecordManager
	projects *models.ProjectManager
	rates *models.RateManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	// usermgr, _ := models.NewUserManager(db)
	consultantmgr, err := models.NewConsultantManager(db)
	if err != nil {
		fmt.Println(err)
	}
	reportedrecordsmgr, err := models.NewReportedRecordManager(db)
	if err != nil {
		fmt.Println(err)
	}
	projectsmgr, err := models.NewProjectManager(db)
	if err != nil {
		fmt.Println(err)
	}
	ratesmgr, err := models.NewRateManager(db)
	if err != nil {
		fmt.Println(err)
	}

	return &API{
		// users:  usermgr,
		consultants: consultantmgr,
		reportedRecords: reportedrecordsmgr,
		projects: projectsmgr,
		rates: ratesmgr,
	}
}

// ResetAPI - drops and creates all empty tables
func ResetAPI(db *models.DB) {

	// db.DropTableIfExists(&models.Users{})
	db.DropTableIfExists(&models.Consultant{})
	db.DropTableIfExists(&models.ReportedRecord{})
	db.DropTableIfExists(&models.Rate{})
	db.DropTableIfExists(&models.Project{})

	fmt.Println("recreated tables:")
	// usermgr, err := models.NewUserManager(db)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("- users")
	_, err := models.NewConsultantManager(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("- consultants")
	_, err = models.NewReportedRecordManager(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("- reported_records")
	models.NewProjectManager(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("- projects")
	models.NewRateManager(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("- rates")
}

// SeedAPI - loads initial data into DB
func SeedAPI(db *models.DB, table string) {
	api := NewAPI(db)
	fmt.Println("Loaded table, # of records, filename:")
	switch table {
	case "rates", "consultants", "projects", "reported_records":
		SeedTable(api, table)
	case "all":
		// users
		SeedTable(api, "rates")
		SeedTable(api, "consultants")
		SeedTable(api, "projects")
		SeedTable(api, "reported_records")
	}
}

// SeedTable -
func SeedTable(api *API, table string) (count int){
	switch table {
	case "rates":
		if api.rates.RateCount() > 0 {
			fmt.Printf("- rates, file %s skipped, table contains %d records\n", viper.GetString("data.rates"), api.rates.RateCount())
			return 0
		}
		count = api.rates.RateSeed("./data./" + viper.GetString("data.rates"))
		fmt.Printf("- rates, %d records, %s\n", count, viper.GetString("data.rates"))
	case "consultants":
		if api.consultants.ConsultantCount() > 0 {
			fmt.Printf("- consultants, file %s skipped, table contains %d records\n", viper.GetString("data.consultants"), api.consultants.ConsultantCount())
			return 0
		}
		count = api.consultants.ConsultantSeed("./data./" + viper.GetString("data.consultants"))
		fmt.Printf("- consultants, %d records, %s\n", count, viper.GetString("data.consultants"))
	case "projects":
		if api.projects.ProjectCount() > 0 {
			fmt.Printf("- projects, file %s skipped, table contains %d records\n", viper.GetString("data.projects"), api.projects.ProjectCount())
			return 0
		}
		count = api.projects.ProjectSeed("./data./" + viper.GetString("data.projects"))
		fmt.Printf("- projects, %d records, %s\n", count, viper.GetString("data.projects"))
	case "reported_records":
		if api.reportedRecords.ReportedRecordCount() > 0 {
			fmt.Printf("- reported_records, file %s skipped, table contains %d records\n", viper.GetString("data.reportedRecords"), api.reportedRecords.ReportedRecordCount())
			return 0
		}
		count = api.reportedRecords.ReportedRecordSeed("./data./" + viper.GetString("data.reportedRecords"))
		fmt.Printf("- reported_records, %d records, %s\n", count, viper.GetString("data.reportedRecords"))
	default:
		fmt.Printf("unknown table to seed: %s\n", table)
	}
	return count
}

// CheckAndInitAPI - loads initial data into DB
func CheckAndInitAPI(db *models.DB) (api *API) {
	fmt.Println("checking DB ...")
	emptyTable := false
	api = NewAPI(db)
	if api.rates.RateCount() == 0 {
		SeedTable(api, "rates")
		emptyTable = true
	}
	if api.consultants.ConsultantCount() == 0 {
		SeedTable(api, "consultants")
		emptyTable = true
	}
	if api.projects.ProjectCount() == 0 {
		SeedTable(api, "projects")
		emptyTable = true
	}
	if emptyTable {
		fmt.Println("loaded missing required data (see tables above)")
	}
	return api
}