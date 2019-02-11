// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"fmt"
	"github.com/valasek/timesheet/server/api"
	"github.com/valasek/timesheet/server/routes"
	"github.com/valasek/timesheet/server/logger"

	"github.com/robfig/cron"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"github.com/meatballhat/negroni-logrus"
	"github.com/rs/cors"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server on URL and port defined in config.yaml",
	Long: `Starts the server on URL and port defined in config.yaml.
	
Command first tests connection to DB, checks if DB contains at least one record in tables
projects, rates, consultants and holidays. If succeeds it will start server.`,
	Run: func(cmd *cobra.Command, args []string) {

		// prepare the server
		url := viper.GetString("url")
		port := viper.GetString("port")
		db := api.ConnectDB()
		defer db.Close()
		apiInst := api.CheckAndInitAPI(db)
		r := routes.NewRoutes(apiInst)
		n := negroni.New()
		n.Use(negroni.NewRecovery())
		n.Use(negronilogrus.NewMiddlewareFromLogger(logger.Log, "web"))
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		})
		n.Use(c)
		n.UseHandler(r)

		// schedule DB backups
		interval := viper.GetString("backup.interval")
		rotation := viper.GetInt("backup.rotation")
		location := viper.GetString("backup.location")
		scheduler := cron.New()
		switch interval {
			case "weekly":
				scheduler.AddFunc("0 0 0 * * 0", func() { api.BackupAPI(rotation, location, db) })
			case "daily":
				scheduler.AddFunc("0 0 0 * * *", func() { api.BackupAPI(rotation, location, db) })
				// test schedule - every minute
				// scheduler.AddFunc("0 * * * * *", func() { api.BackupAPI(rotation, location, db) })
			default:
				logger.Log.Warning(fmt.Sprintf("miconfigured backup interval;: %s, allowed intervals daily or weekly, falling back to daily", interval))
				// logrus.WithFields(logrus.Fields{"category": "backup"}).Error("config file:      ", viper.ConfigFileUsed())
				interval = "daily"
				scheduler.AddFunc("0 0 0 * * *", func() { api.BackupAPI(rotation, location, db) })
		}
		scheduler.Start()
		defer scheduler.Stop()
		logger.Log.Info(fmt.Sprintf("DB backups scheduled %s, %d backups back kept in location %s", interval, rotation, location))	

		// run the server
		n.Run(url + ":" + port)	
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
