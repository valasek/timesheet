// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"github.com/valasek/time-sheet/api"
	"github.com/spf13/cobra"
)

var (
	load string
	clean bool
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Loads or cleans DB. See time-sheet help db",
	Long: `Loads or cleans DB.
	
Command first tests connection to DB. If succeeds it will load or clean db and exit.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := ConnectDB()
		defer db.Close()
		if clean {
			api.ResetAPI(db)
		} else if len(load) > 1 {
			api.SeedAPI(db, load)
		} else {
			cmd.Help()
		}
	
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().StringVarP(&load, "load", "l", "", `Truncate DB table/tables and load initial data from files in folder ./data. Options:
all - apply on all tables
rates | consultants | projects | holidays | reported_records - apply only on one selected table`)
	dbCmd.PersistentFlags().BoolVarP(&clean,"clean", "c", false, `Drop and create all required DB tables`)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
