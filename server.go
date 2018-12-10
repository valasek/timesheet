package main

import (
	"github.com/valasek/time-sheet/api"
	"github.com/valasek/time-sheet/models"
	"github.com/valasek/time-sheet/routes"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
)

func main() {
	db := models.NewPostgresDB("host=localhost port=5432 user=time-sheet dbname=time-sheet password=time-sheet sslmode=disable")
	api := api.NewAPI(db)
	r := routes.NewRoutes(api)
	routes.PrintRoutes(r)
	n := negroni.Classic()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	n.Use(c)
	n.UseHandler(r)
	n.Run(":3000")
}
