// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package routes

import (
	"github.com/valasek/timesheet/server/api"
	"github.com/valasek/timesheet/server/logger"

	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
)

var w *tabwriter.Writer

func noRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	if (path[1] != "") && (path[1] == "api") {
		c.JSON(http.StatusNotFound, gin.H{"msg": "no route", "body": nil})
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

// Logger provides logrus logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		// after request
		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		message := fmt.Sprintf("[server] | %3d | %12v |%s | %-7s %s %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			c.Errors.String(),
		)
		switch {
		case statusCode >= 400 && statusCode <= 499:
			logger.Log.Warning(message)
		case statusCode >= 500:
			logger.Log.Error(message)
		default:
			logger.Log.Info(message)
		}
	}
}

// SetupRouter builds the routes for the api
func SetupRouter(api *api.API) *gin.Engine {

	gin.DefaultWriter = colorable.NewColorableStdout()
	if viper.GetString("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(Logger())

	// set CORS
	router.Use(cors.Default())

	// no route, bad url
	router.NoRoute(noRoute)

	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))

	a := router.Group("/api")
	{

		// app settings
		a.GET("/settings", api.AppSettings)
		// download all data
		a.GET("/download/data", api.Download)
		// upload all data
		a.POST("/upload/data", api.Upload)
		// download logs
		a.GET("/download/logs/:logLevel", api.DownloadLogs)
		// download docs
		a.GET("/download/docs", api.DownloadDocs)

		// consultants
		a.GET("/consultants", api.ConsultantList)

		// projects
		a.GET("/projects", api.ProjectsGetAll)

		// rates
		a.GET("/rates", api.RatesGetAll)

		// holidays
		a.GET("/holidays", api.HolidaysGetAll)

		// reported records
		a.POST("/reported", api.ReportedRecordsAddRecord)
		a.GET("/reported", api.ReportedRecordsGetAll)
		a.GET("/reported/summary/:year", api.ReportedRecordsSummary)
		a.GET("/reported/year/:year/month/:month/consultant/:consultant", api.ReportedRecordsInMonth)
		a.DELETE("/reported/:id", api.ReportedRecordDelete)
		a.PUT("/reported/:id", api.ReportedRecordUpdate)

	}
	// handle 404 and due to Vue history mode return home page
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "client", "dist", "index.html"))
	})

	return router
}

// PrintRoutes prints all set routes
func PrintRoutes(c *gin.Engine) {
	fmt.Println("gin.DebugPrintRouteFunc()")
}
