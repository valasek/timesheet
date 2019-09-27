// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"github.com/valasek/timesheet/server/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	load     string
	generate bool
	clean    bool
	backup   bool
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Initiate, load. backup DB or generate demo data. See timesheet help db",
	Long: `Initiate, load, backup DB or generate demo data.

Command first tests connection to DB. If succeeds it will initiate, load, backup db or generate demo data and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := api.ConnectDB()
		defer db.Close()
		if clean {
			api.ResetAPI(db)
		} else {
			if generate {
				api.GenerateAPI(viper.GetString("export.location"), db)
			} else {
				if len(load) > 2 {
					// ignore the error, is already logged
					api.SeedAPI(db, load, api.FileList())
				} else {
					if backup {
						api.BackupAPI(viper.GetInt("backup.rotation"), viper.GetString("backup.location"), db)
					} else {
						cmd.Help()
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, `Generate demo data and save them into ./data folder`)
	dbCmd.PersistentFlags().StringVarP(&load, "load", "l", "", `Truncate DB table/tables and load initial data from files in folder ./data. Options:
all - load all tables
rates | consultants | projects | holidays | reported_records - load selected table`)
	dbCmd.PersistentFlags().BoolVarP(&clean, "clean", "c", false, `Drop and create all required DB tables`)
	dbCmd.PersistentFlags().BoolVarP(&backup, "backup", "b", false, `Backup all DB tables in the format used by db --load command`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
