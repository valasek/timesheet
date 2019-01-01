// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"github.com/valasek/timesheet/api"
	"github.com/valasek/timesheet/routes"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
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
		url := viper.GetString("url")
		port := viper.GetString("port")
		db := ConnectDB()
		defer db.Close()
		api := api.CheckAndInitAPI(db)
		r := routes.NewRoutes(api)
		n := negroni.Classic()
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		})
		n.Use(c)
		n.UseHandler(r)
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
