// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"archive/zip"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/models"

	"github.com/spf13/viper"
)

// FileList returs map tables and input file for initial seeding
func FileList() map[string]string {
	list := map[string]string{
		"rates":            filepath.Join(".", "data", viper.GetString("data.rates")),
		"consultants":      filepath.Join(".", "data", viper.GetString("data.consultants")),
		"projects":         filepath.Join(".", "data", viper.GetString("data.projects")),
		"reported_records": filepath.Join(".", "data", viper.GetString("data.reportedRecords")),
		"holidays":         filepath.Join(".", "data", viper.GetString("data.holidays")),
	}
	return list
}

func uploadedFileList() (list map[string]string, err error) {
	list = make(map[string]string)
	files, err := ioutil.ReadDir(filepath.Clean(viper.GetString("uploadFolderTemp")))
	if err != nil {
		return nil, err
	}

	if len(files) != 5 {
		return nil, errors.New("archive should contain 5 files")
	}

	for _, file := range files {
		table := tableFromFilename(file.Name())
		if len(table) > 0 {
			if _, ok := list[table]; ok {
				return nil, errors.New("archive contains same data: " + table)
			}
			list[table] = filepath.Join(".", viper.GetString("uploadFolderTemp"), file.Name())
		} else {
			logger.Log.Warn(file.Name(), " - ignored")
		}
	}

	if len(list) != 5 {
		logger.Log.Error("expected 5 files, got: ", list)
		return nil, errors.New("expected 5 files, got: " + strconv.Itoa(len(list)))
	}

	fmt.Println(list)
	return list, nil
}

func tableFromFilename(file string) string {
	tables := [5]string{"rates", "holidays", "consultants", "reported_records", "projects"}
	for _, table := range tables {
		if strings.Contains(file, table) {
			return table
		}
	}
	return ""
}

func restoreDB(f *os.File) error {
	// delete if target folder exists
	err := os.RemoveAll(viper.GetString("uploadFolderTemp"))
	if err != nil {
		return err
	}
	err = os.Mkdir(viper.GetString("uploadFolderTemp"), os.ModeDir)
	if err != nil {
		return err
	}

	err = unzip(f.Name(), viper.GetString("uploadFolderTemp"))
	if err != nil {
		return err
	}

	files, err := uploadedFileList()
	if err != nil {
		return err
	}

	db := ConnectDB()
	ResetAPI(db)
	err = SeedAPI(db, "all", files)

	return err
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			logger.Log.Error(err)
			panic(err)
		}
	}()

	extractAndWrite := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				logger.Log.Error(err)
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			return errors.New("got folder, zip and upload only csv files")
		}
		os.MkdirAll(filepath.Dir(path), f.Mode())
		ff, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := ff.Close(); err != nil {
				logger.Log.Error(err)
				panic(err)
			}
		}()

		_, err = io.Copy(ff, rc)
		if err != nil {
			return err
		}

		return nil
	}

	for _, f := range r.File {
		err := extractAndWrite(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// ConnectDB connects and pings DB
func ConnectDB() (db *models.DB) {
	switch DBType := viper.GetString("dbType"); DBType {
	case "postgresql":
		// DBhost, DBport, DBuser, DBpassword, DBname, SSLmode, url, port := "", "", "", "", "", "", "", ""
		dbURL := viper.GetString("DATABASE_URL")

		if len(dbURL) == 0 {
			dbURL = "host=" + viper.GetString("postgresql.host") +
				" port=" + viper.GetString("postgresql.port") +
				" user=" + viper.GetString("postgresql.user") +
				" dbname=" + viper.GetString("postgresql.name") +
				" password=" + viper.GetString("postgresql.password") +
				" sslmode=" + viper.GetString("postgresql.SSLMode")
		}
		logger.Log.Info("connecting to DB ", dbURL)
		db = models.NewPostgresDB(dbURL)
		logger.Log.Info("connected to DB ", dbURL)
	case "mysql":
		dbURL := viper.GetString("DATABASE_URL")

		if len(dbURL) == 0 {
			dbURL = viper.GetString("mysql.user") +
				":" + viper.GetString("mysql.password") +
				"@/" + viper.GetString("mysql.dbname") +
				"?charset=utf8&parseTime=True&loc=Local"
		}
		// gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
		logger.Log.Info("connecting to DB ", dbURL)
		db = models.NewMySQLDB(dbURL)
		logger.Log.Info("connected to DB ", dbURL)
	default:
		logger.Log.Error("not able to connect to DB, supported DB types (postgresql, mysql), set: ", DBType)
		os.Exit(1)
	}
	return db
}

func rotateBackupFile(rotation int, folder, baseFileName string) error {
	files, err := ioutil.ReadDir(filepath.Clean(folder))
	if err != nil {
		return err
	}

	oldestTime := time.Now()
	var oldestFile os.FileInfo
	var filteredNames []string
	if len(files) == 0 {
		return nil
	}
	for _, file := range files {
		if strings.Contains(file.Name(), baseFileName) {
			filteredNames = append(filteredNames, file.Name())
			if file.Mode().IsRegular() && file.ModTime().Before(oldestTime) {
				oldestFile = file
				oldestTime = file.ModTime()
			}
		}
	}
	if len(filteredNames) >= rotation {
		err := os.Remove(filepath.Join(folder, oldestFile.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filepath.Base(filename))
	if err != nil {
		msg := "failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func cleanExportedFiles(folder string) error {
	dir, err := os.Open(folder)
	if err != nil {
		return err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, f := range files {
		fName := f.Name()
		fNamePath := filepath.Join(folder, fName)
		os.Remove(fNamePath)
	}
	return nil
}

// exports all data from DB into file timesheet-backup.zip
func export() (fileName string, err error) {
	fileName = "timesheet-backup.zip"
	db := ConnectDB()
	defer db.Close()
	exportFolder := viper.GetString("export.location")

	err = cleanExportedFiles(exportFolder)
	if err != nil {
		return "", err
	}

	BackupAPI(viper.GetInt("backup.rotation"), exportFolder, db)
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile("timesheet-backup.zip", flags, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	files, err := ioutil.ReadDir(exportFolder)
	if err != nil {
		return "", err
	}

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, file := range files {
		err := appendFiles(filepath.Join(exportFolder, file.Name()), zipw)
		if err != nil {
			return "", err
		}
	}

	return fileName, nil
}
