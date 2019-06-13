// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/models"
	"github.com/valasek/timesheet/server/version"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// API -
type API struct {
	consultants     *models.ConsultantManager
	reportedRecords *models.ReportedRecordManager
	projects        *models.ProjectManager
	rates           *models.RateManager
	holidays        *models.HolidayManager
}

// AppSettings -
type AppSettings struct {
	Version              string  `json:"version"`
	DailyWorkingHours    float64 `json:"dailyWorkingHours"`
	DailyWorkingHoursMin float64 `json:"dailyWorkingHoursMin"`
	DailyWorkingHoursMax float64 `json:"dailyWorkingHoursMax"`
	Vacation             string  `json:"vacation"`
	YearlyVacationDays   int64   `json:"yearlyVacationDays"`
	VacationPersonal     string  `json:"vacationPersonal"`
	YearlyPersonalDays   int64   `json:"yearlyPersonalDays"`
	VacationSick         string  `json:"vacationSick"`
	YearlySickDays       int64   `json:"yearlySickDays"`
	IsWorking            string  `json:"isWorking"`
	IsNonWorking         string  `json:"isNonWorking"`
}

// AppSettings returns list of all appliocation and user settings for configuration file
func (api *API) AppSettings(c *gin.Context) {
	settings := AppSettings{
		Version:              version.Version,
		DailyWorkingHours:    viper.GetFloat64("dailyWorkingHours"),
		DailyWorkingHoursMin: viper.GetFloat64("dailyWorkingHoursMin"),
		DailyWorkingHoursMax: viper.GetFloat64("dailyWorkingHoursMax"),
		Vacation:             viper.GetString("vacation"),
		YearlyVacationDays:   viper.GetInt64("yearlyVacationDays"),
		VacationPersonal:     viper.GetString("vacationPersonal"),
		YearlyPersonalDays:   viper.GetInt64("yearlyPersonalDays"),
		VacationSick:         viper.GetString("vacationSick"),
		YearlySickDays:       viper.GetInt64("yearlySickDays"),
		IsWorking:            viper.GetString("isWorking"),
		IsNonWorking:         viper.GetString("isNonWorking"),
	}
	c.JSON(http.StatusOK, settings)
}

// Download -
func (api *API) Download(c *gin.Context) {
	fileName, err := export()
	if err != nil {
		c.String(http.StatusNotFound, "downloading data failed with error: "+err.Error())
		return
	}

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		c.String(http.StatusNotFound, "file not found")
		return
	}

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	file.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := file.Stat()  //Get info from file
	FileSize := FileStat.Size() //Get file size

	//Send the headers
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="timesheet-backup.zip"`,
	}

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	file.Seek(0, 0)
	c.DataFromReader(http.StatusOK, FileSize, FileContentType, file, extraHeaders)
	return
}

// Upload -
func (api *API) Upload(c *gin.Context) {

	// parse and validate file and post parameters
	ffile, err := c.FormFile("timesheet-backup.zip")
	if err != nil {
		logger.Log.Error("unable to upload file, INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}
	file, err := ffile.Open()
	if err != nil {
		logger.Log.Error("unable open file, INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Log.Error("unable to upload file: INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	if (filetype != "application/zip") && (filetype != "application/x-zip-compressed") {
		logger.Log.Error("unable to upload file, INVALID_FILE_TYPE (supported: application/zip, application/x-zip-compressed): ", filetype)
		c.String(http.StatusBadRequest, "INVALID_FILE_TYPE: "+filetype)
		return
	}
	fileName := randToken(12) + ".zip"
	uploadPath := viper.GetString("uploadFolder")
	newPath := filepath.Join(uploadPath, fileName)
	// FIXME check file type
	// fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		logger.Log.Error("unable to upload file, CANT_WRITE_FILE: ", err)
		c.String(http.StatusInternalServerError, "CANT_WRITE_FILE: "+err.Error())
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		logger.Log.Error("unable to upload file. CANT_WRITE_FILE: ", err)
		c.String(http.StatusInternalServerError, "CANT_WRITE_FILE: "+err.Error())
		return
	}
	if err := restoreDB(newFile); err != nil {
		logger.Log.Error("unable to restore: ", err)
		c.String(http.StatusInternalServerError, "CANT_RESTORE: "+err.Error())
		return
	}
	err = os.RemoveAll(viper.GetString("uploadFolder"))
	if err != nil {
		logger.Log.Error("unable to restore: ", err)
		c.String(http.StatusInternalServerError, "CANT_RESTORE: "+err.Error())
		return
	}
	err = os.Mkdir(viper.GetString("uploadFolder"), os.ModeDir)
	if err != nil {
		logger.Log.Error("unable to restore: ", err)
		c.String(http.StatusInternalServerError, "CANT_RESTORE: "+err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", ffile.Filename))
}

// DownloadDocs -
func (api *API) DownloadDocs(c *gin.Context) {

	fileName := filepath.Join("documentation", "documentation.md")
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		c.String(http.StatusOK, fileName+" does not exist")
		return
	}
	fi, err := f.Stat()
	if err != nil {
		c.String(http.StatusOK, fileName+" cannot get file size")
		return
	}
	if fi.Size() == 0 {
		c.String(http.StatusOK, fileName+" does not exist")
		return
	}
	c.File(fileName)
}

// DownloadLogs -
func (api *API) DownloadLogs(c *gin.Context) {

	logLevel := c.Param("logLevel")
	if len(logLevel) < 1 {
		logger.Log.Error("unable to download log files, param 'logLevel' is missing")
		c.String(http.StatusInternalServerError, "unable to download log files, param 'logLevel' is missing")
		return
	}

	file := ""
	switch logLevel {
	case "0":
		file = "info.log"
	case "1":
		file = "error.log"
	default:
		logger.Log.Error("unable to download log files, unknown logLevel: ", logLevel)
	}
	fileName := filepath.Join(viper.GetString("logFolder"), file)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		c.String(http.StatusOK, file+" contains no log entries")
		return
	}
	fi, err := f.Stat()
	if err != nil {
		c.String(http.StatusOK, file+" cannot get file size")
		return
	}
	if fi.Size() == 0 {
		c.String(http.StatusOK, file+" contains no log entries")
		return
	}
	c.File(fileName)
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	consultantmgr, err := models.NewConsultantManager(db)
	if err != nil {
		logger.Log.Error(err)
	}
	reportedrecordsmgr, err := models.NewReportedRecordManager(db)
	if err != nil {
		logger.Log.Error(err)
	}
	projectsmgr, err := models.NewProjectManager(db)
	if err != nil {
		logger.Log.Error(err)
	}
	ratesmgr, err := models.NewRateManager(db)
	if err != nil {
		logger.Log.Error(err)
	}
	holidaysmgr, err := models.NewHolidayManager(db)
	if err != nil {
		logger.Log.Error(err)
	}

	return &API{
		// users:  usermgr,
		consultants:     consultantmgr,
		reportedRecords: reportedrecordsmgr,
		projects:        projectsmgr,
		rates:           ratesmgr,
		holidays:        holidaysmgr,
	}
}

// ResetAPI - drops and creates all empty tables
func ResetAPI(db *models.DB) {

	// db.DropTableIfExists(&models.Users{})
	db.DropTableIfExists(&models.Consultant{})
	db.DropTableIfExists(&models.ReportedRecord{})
	db.DropTableIfExists(&models.Rate{})
	db.DropTableIfExists(&models.Project{})
	db.DropTableIfExists(&models.Holiday{})

	logger.Log.Info("recreated tables:")
	_, err := models.NewConsultantManager(db)
	if err != nil {
		logger.Log.Error("recreated tables:", err)
	}
	logger.Log.Info("- consultants")
	_, err = models.NewReportedRecordManager(db)
	if err != nil {
		logger.Log.Error("recreated tables:", err)
	}
	logger.Log.Info("- reported_records")
	models.NewProjectManager(db)
	if err != nil {
		logger.Log.Error("recreated tables:", err)
	}
	logger.Log.Info("- projects")
	models.NewRateManager(db)
	if err != nil {
		logger.Log.Error("recreated tables:", err)
	}
	logger.Log.Info("- rates")
	models.NewHolidayManager(db)
	if err != nil {
		logger.Log.Error("recreated tables:", err)
	}
	logger.Log.Info("- holidays")
}

// SeedAPI - loads initial data into DB
func SeedAPI(db *models.DB, table string, inputFiles map[string]string) error {
	api := NewAPI(db)
	logger.Log.Info("Loaded table, # of records, filename:")
	switch table {
	case "rates", "consultants", "projects", "reported_records", "holidays":
		SeedTable(api, table, inputFiles[table])
	case "all":
		// users
		SeedTable(api, "rates", inputFiles["rates"])
		SeedTable(api, "consultants", inputFiles["consultants"])
		SeedTable(api, "projects", inputFiles["projects"])
		SeedTable(api, "reported_records", inputFiles["reported_records"])
		SeedTable(api, "holidays", inputFiles["holidays"])
	default:
		logger.Log.Error("unable to seed non-existent table: ", table)
		return errors.New("unable to seed non-existent table: " + table)
	}
	return nil
}

// SeedTable -
func SeedTable(api *API, table, file string) (count int) {
	switch table {
	case "rates":
		if api.rates.RateCount() > 0 {
			logger.Log.Warn(fmt.Sprintf("- rates, file %s skipped, table contains %d records", file, api.rates.RateCount()))
			return 0
		}
		count = api.rates.RateSeed(file)
		logger.Log.Info(fmt.Sprintf("- rates, %d records, %s", count, file))
	case "consultants":
		if api.consultants.ConsultantCount() > 0 {
			logger.Log.Warn(fmt.Sprintf("- consultants, file %s skipped, table contains %d records", file, api.consultants.ConsultantCount()))
			return 0
		}
		count = api.consultants.ConsultantSeed(file)
		logger.Log.Info(fmt.Sprintf("- consultants, %d records, %s", count, file))
	case "projects":
		if api.projects.ProjectCount() > 0 {
			logger.Log.Warn(fmt.Sprintf("- projects, file %s skipped, table contains %d records", file, api.projects.ProjectCount()))
			return 0
		}
		count = api.projects.ProjectSeed(file)
		logger.Log.Info(fmt.Sprintf("- projects, %d records, %s", count, file))
	case "reported_records":
		if api.reportedRecords.ReportedRecordCount() > 0 {
			logger.Log.Warn(fmt.Sprintf("- reported_records, file %s skipped, table contains %d records", file, api.reportedRecords.ReportedRecordCount()))
			return 0
		}
		count = api.reportedRecords.ReportedRecordSeed(file)
		logger.Log.Info(fmt.Sprintf("- reported_records, %d records, %s", count, file))
	case "holidays":
		if api.holidays.HolidayCount() > 0 {
			logger.Log.Warn(fmt.Sprintf("- holidays, file %s skipped, table contains %d records", file, api.holidays.HolidayCount()))
			return 0
		}
		count = api.holidays.HolidaySeed(file)
		logger.Log.Info(fmt.Sprintf("- holidays, %d records, %s", count, file))
	default:
		logger.Log.Warn("unknown table to seed: ", table)
	}
	return count
}

// CheckAndInitAPI - loads initial data into DB
func CheckAndInitAPI(db *models.DB) (api *API) {
	logger.Log.Info("checking DB ...")
	files := FileList()
	emptyTable := false
	api = NewAPI(db)
	if api.rates.RateCount() == 0 {
		SeedTable(api, "rates", files["rates"])
		emptyTable = true
	}
	if api.consultants.ConsultantCount() == 0 {
		SeedTable(api, "consultants", files["consultants"])
		emptyTable = true
	}
	if api.projects.ProjectCount() == 0 {
		SeedTable(api, "projects", files["projects"])
		emptyTable = true
	}
	if api.reportedRecords.ReportedRecordCount() == 0 {
		SeedTable(api, "reported_records", files["reported_records"])
		emptyTable = true
	}
	if api.holidays.HolidayCount() == 0 {
		SeedTable(api, "holidays", files["holidays"])
		emptyTable = true
	}
	if emptyTable {
		logger.Log.Info("loaded missing required data (see tables above)")
	}
	return api
}

// BackupAPI - drops and creates all empty tables
func BackupAPI(rotation int, folder string, db *models.DB) {

	api := NewAPI(db)
	id := time.Now().Format("2006-01-02_150405")

	logger.Log.Info("backuped tables:")
	tableNames := []string{"rates", "projects", "reported_records", "consultants", "holidays"}
	for _, baseFileName := range tableNames {
		err := rotateBackupFile(rotation, folder, baseFileName)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("not able to rotate %s backup files, backups stopped, handle the error: %s", baseFileName, err))
		}

		fileName := baseFileName + "_" + id + ".csv"
		filePath := filepath.Join(folder, fileName)
		n := 0
		switch baseFileName {
		case "projects":
			n, err = api.projects.ProjectBackup(filePath)
		case "rates":
			n, err = api.rates.RateBackup(filePath)
		case "consultants":
			n, err = api.consultants.ConsultantBackup(filePath)
		case "holidays":
			n, err = api.holidays.HolidayBackup(filePath)
		case "reported_records":
			n, err = api.reportedRecords.ReportedRecordBackup(filePath)
		}
		if err != nil {
			logger.Log.Error(fmt.Sprintf("backuped tables: error during %s backup: %s", baseFileName, err))
		} else {
			logger.Log.Info(fmt.Sprintf("- %s, %d records", baseFileName, n))
		}
	}
}
