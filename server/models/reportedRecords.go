// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package models

import (
	"github.com/valasek/timesheet/server/logger"

	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jinzhu/now"

	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ReportedRecord struct
type ReportedRecord struct {
	ID          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Date        time.Time  `gorm:"not null" json:"date"`
	Hours       float64    `gorm:"not null" json:"hours"`
	Project     string     `gorm:"not null" json:"project"`
	Description string     `gorm:"not null" json:"description"`
	Rate        string     `gorm:"not null" json:"rate"`
	Consultant  string     `gorm:"not null" json:"consultant"`
}

// ReportedRecordsSummary struct
type ReportedRecordsSummary struct {
	Consultant string  `json:"consultant"`
	Year       string  `json:"year"`
	Month      string  `json:"month"`
	Project    string  `json:"project"`
	Rate       string  `json:"rate"`
	Hours      float64 `json:"hours"`
}

// ReportedRecordCSV csv struct
type ReportedRecordCSV struct {
	CreatedAt   DateTime `csv:"created_at"`
	Date        Date     `csv:"date"`
	Hours       float64  `csv:"hours"`
	Project     string   `csv:"project"`
	Description string   `csv:"description"`
	Rate        string   `csv:"rate"`
	Consultant  string   `csv:"consultant"`
}

// ReportedRecordManager struct
type ReportedRecordManager struct {
	db *DB
}

// NewReportedRecordManager - Create a ReportedRecords manager that can be used for retrieving ReportedRecordss
func NewReportedRecordManager(db *DB) (*ReportedRecordManager, error) {

	db.AutoMigrate(&ReportedRecord{})

	reportedRecordsMgr := ReportedRecordManager{}

	reportedRecordsMgr.db = db

	return &reportedRecordsMgr, nil
}

// ReportedRecordsGetAll - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsGetAll() []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Find(&reportedRecords); err != nil {
		return reportedRecords
	}
	logger.Log.Error("unable to retieve all reported records")
	return nil
}

// ReportedRecordsInMonth - return records per month
// adds days from previous and next month having the same week as month
func (db *ReportedRecordManager) ReportedRecordsInMonth(year, month, consultant string) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("extract(MONTH from date) = ? AND extract(YEAR from date) = ? AND consultant = ?", month, year, consultant).Find(&reportedRecords); err != nil {
		borderWeeksReportedRecords := []ReportedRecord{}
		days, err := getborderDays(year, month)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("failed - get reported records for month %s, year %s and consultant %s, error: %s", month, year, consultant, err))
			return nil
		}
		if err := db.db.Where("date(date) in (?) AND consultant = ?", days, consultant).Find(&borderWeeksReportedRecords); err != nil {
			reportedRecords = append(reportedRecords, borderWeeksReportedRecords...)
			return reportedRecords
		}
	}
	logger.Log.Error(fmt.Sprintf("failed - get reported records for month %s, year %s and consultant %s", month, year, consultant))
	return nil
}

func getborderDays(year, month string) (days []string, err error) {
	layout := "2006-1-02"
	monthStart, err := time.Parse(layout, year+"-"+month+"-01")
	if err != nil {
		return nil, err
	}
	monthN, err := strconv.Atoi(month)
	if err != nil {
		return nil, err
	}
	yearN, err := strconv.Atoi(year)
	if err != nil {
		return nil, err
	}
	prevMonth := fmt.Sprintf("%02d", monthN-1)
	prevYear, nextYear := strconv.Itoa(yearN), strconv.Itoa(yearN)
	if monthN == 1 {
		prevMonth = "12"
		prevYear = strconv.Itoa(yearN - 1)
	}
	nextMonth := fmt.Sprintf("%02d", monthN+1)
	if monthN == 12 {
		nextMonth = "01"
		nextYear = strconv.Itoa(yearN + 1)
	}
	monday := now.New(monthStart).Monday()
	sunday := now.New((now.New(monthStart).EndOfMonth())).Sunday()
	monthPrevious, err := time.Parse(layout, year+"-"+prevMonth+"-01")
	if err != nil {
		return nil, err
	}

	for day := monday.Day(); day <= now.New(monthPrevious).EndOfMonth().Day(); day++ {
		days = append(days, prevYear+"-"+prevMonth+"-"+fmt.Sprintf("%02d", day))
	}

	if sunday.Month() == monthStart.Month() {
		// fmt.Println("sunday is also end of the month")
	} else {
		for day := 1; day <= sunday.Day(); day++ {
			days = append(days, nextYear+"-"+nextMonth+"-"+fmt.Sprintf("%02d", day))
		}
	}
	return days, nil
}

// ReportedRecordsSummary - return summary records per selected year
func (db *ReportedRecordManager) ReportedRecordsSummary(year string) []ReportedRecordsSummary {
	reportedRecordsSummary := []ReportedRecordsSummary{}
	sql := fmt.Sprintf("select consultant, extract(MONTH from date) as month, project, rate, sum(hours) as hours from reported_records where extract(YEAR from date) = %s and deleted_at is null group by consultant, extract(MONTH from date), project, rate", year)
	if err := db.db.Raw(sql).Scan(&reportedRecordsSummary); err != nil {
		// fmt.Println(reportedRecordsSummary)
		return reportedRecordsSummary
	}
	logger.Log.Error("unable to retrieve reported records summary")
	return nil
}

// ReportedRecordsDelete - return all records of ReportedRecords
func (db *ReportedRecordManager) ReportedRecordsDelete(id uint64) []ReportedRecord {
	reportedRecords := []ReportedRecord{}
	if err := db.db.Where("id = ?", id).Delete(&reportedRecords); err != nil {
		return reportedRecords
	}
	logger.Log.Error("unable to delete reported record for id", id)
	return nil
}

// ReportedRecordAdd -
func (db *ReportedRecordManager) ReportedRecordAdd(newRecord ReportedRecord) ReportedRecord {
	if err := db.db.Create(&newRecord); err != nil {
		return newRecord
	}
	logger.Log.Error("unable to add new reported record", newRecord)
	return ReportedRecord{}
}

// ReportedRecordUpdate -
func (db *ReportedRecordManager) ReportedRecordUpdate(r UpdatedValue) ReportedRecord {
	updateValue := UpdatedValue{
		ID:    r.ID,
		Type:  r.Type,
		Value: r.Value,
	}
	reportedRecord := ReportedRecord{}
	// handle attribute types
	switch r.Type {
	case "hours":
		value, err := strconv.ParseFloat(updateValue.Value, 64)
		if err != nil {
			logger.Log.Error(err)
		}
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, value); err != nil {
			return reportedRecord
		}
	case "description", "rate", "project":
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, updateValue.Value); err != nil {
			return reportedRecord
		}
	case "date":
		value, err := time.Parse("2006-01-02", updateValue.Value)
		if err != nil {
			logger.Log.Error(err)
		}
		if err := db.db.First(&reportedRecord, updateValue.ID).Update(updateValue.Type, value); err != nil {
			return reportedRecord
		}
	default:
		logger.Log.Error("failed - updating reported record, unknown attribute type: ", r.Type, updateValue)
		return ReportedRecord{}
	}
	logger.Log.Error("failed - updating reported record", updateValue)
	return ReportedRecord{}
}

// ReportedRecordSeed - will load data from data file
func (db *ReportedRecordManager) ReportedRecordSeed(file string) int {

	csvfile, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		logger.Log.Error(err, " Input file: ", file)
	}
	defer csvfile.Close()

	recordsCSV := []*ReportedRecordCSV{}
	if err := gocsv.UnmarshalFile(csvfile, &recordsCSV); err != nil {
		logger.Log.Error(err)
	}
	for _, r := range recordsCSV {
		newR := ReportedRecord{CreatedAt: r.CreatedAt.Time, Date: r.Date.Time, Hours: r.Hours, Project: r.Project, Description: r.Description, Rate: r.Rate, Consultant: r.Consultant}
		db.db.Create(&newR)
	}

	return len(recordsCSV)
}

// ReportedRecordCount -
func (db *ReportedRecordManager) ReportedRecordCount() int {
	reportedRecords := []ReportedRecord{}
	var count int
	if err := db.db.Find(&reportedRecords).Count(&count); err != nil {
		return count
	}
	logger.Log.Error("unable to retrieve reported records count")
	return 0
}

// ReportedRecordBackup will backup rates table
func (db *ReportedRecordManager) ReportedRecordBackup(filePath string) (int, error) {
	reportedRecordsFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer reportedRecordsFile.Close()

	reportedRecords := []*ReportedRecord{}
	db.db.Find(&reportedRecords).Where("DeletedAt = ?", nil)
	projectCSV := []*ReportedRecordCSV{}
	for _, r := range reportedRecords {
		createdAt := DateTime{r.CreatedAt}
		date := Date{r.Date}
		item := ReportedRecordCSV{CreatedAt: createdAt, Date: date, Hours: r.Hours, Project: r.Project, Description: r.Description, Rate: r.Rate, Consultant: r.Consultant}
		projectCSV = append(projectCSV, &item)
	}

	err = gocsv.MarshalFile(&projectCSV, reportedRecordsFile)
	if err != nil {
		return 0, err
	}
	return len(reportedRecords), nil
}
