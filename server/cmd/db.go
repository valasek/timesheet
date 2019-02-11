// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"github.com/valasek/timesheet/server/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	load   string
	clean  bool
	backup bool
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Initiate, load or backup DB. See timesheet help db",
	Long: `Initiate, load,or backup DB.
	
Command first tests connection to DB. If succeeds it will initiate, load or backup db and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := api.ConnectDB()
		defer db.Close()
		if clean {
			api.ResetAPI(db)
		} else {
			if len(load) > 1 {
				api.SeedAPI(db, load)
			} else {
				if backup {
					api.BackupAPI(viper.GetInt("backup.rotation"), viper.GetString("backup.location"), db)
				} else {
					cmd.Help()
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().StringVarP(&load, "load", "l", "", `Truncate DB table/tables and load initial data from files in folder ./data. Options:
all - load all tables
rates | consultants | projects | holidays | reported_records - load selected table`)
	dbCmd.PersistentFlags().BoolVarP(&clean, "clean", "c", false, `Drop and create all required DB tables`)
	dbCmd.PersistentFlags().BoolVarP(&backup, "backup", "b", false, `Backup all DB tables in the format used by db --load command`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
