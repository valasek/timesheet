// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package routes

import (
	"github.com/valasek/timesheet/server/api"

	"fmt"
	"net/http"
	"path/filepath"
	"text/tabwriter"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/gin-contrib/cors"
	// "github.com/thinkerou/favicon"
	"github.com/mattn/go-colorable"
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

// SetupRouter builds the routes for the api
func SetupRouter(api *api.API) *gin.Engine {

	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// router.Use(favicon.New(filepath.Join("..", "client", "dist", "favicon.png")))

	// set CORS
	router.Use(cors.Default())

	// no route, bad url
	router.NoRoute(noRoute)

	// FIXME
	// TODO
	// https://github.com/ndabAP/vue-go-example
	// https://jonathanmh.com/creating-simple-markdown-blog-go-gin/

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