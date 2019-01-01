# Timesheet for consultants

![GitHub release](https://img.shields.io/github/release-pre/valasek/timesheet.svg)
[![GitHub issues](https://img.shields.io/github/issues/valasek/timesheet.svg)](https://github.com/valasek/timesheet/issues)
[![GitHub license](https://img.shields.io/github/license/valasek/timesheet.svg)](https://github.com/valasek/timesheet/blob/master/LICENSE)
[![PayPal Donate](https://img.shields.io/badge/donate-PayPal.me-ff69b4.svg)](https://paypal.me/StanislavValasek)

Web application to report consulting hours on projects using selected rates on a weekly bases. Supports export into csv.

# Requirements

- Linux, Windows or macOS
- DB connection to PostgreSQL
- Initial yearly data in csv format (dates should be entered in ISO format YYYY-MM-DD HH:MM:SS):
-- consultants, projects, rates, holidays, initial reported records - optional). 

# Download - alpha release
* [timesheet.exe](https://github.com/valasek/rajce-get/releases/download/v0.0.1/timesheet.exe) (MS Windows 64bit)
* please build from source (Linux 64bit)
* please build from source (Mac OS X 64bit)

# Usage

```
Web based timesheet application with DB persistence.

Application reads DB and server configuration from config.toml, loads default data if DB is empty and launch web GUI.

Usage:
  timesheet [command]

Available Commands:
  db          Loads or cleans DB. See timesheet help db
  help        Help about any command
  routes      Prints the list of all available routes
  server      Starts the server on URL and port defined in config.yaml

Flags:
      --config string   config file (default is ./timesheet.yaml)
  -h, --help            help for timesheet
  -v, --version         Prints application versions

Use "timesheet [command] --help" for more information about a command.
```

# Standing on the shoulders of giants

[Go](https://golang.org/), [Vue](https://vuejs.org/), [Vuetify](https://vuetifyjs.com/en/), [PostgreSQL](https://www.postgresql.org/)

## Backend

- Middleware: [Negroni](https://github.com/urfave/negroni)
- Router: [Gorilla](https://github.com/gorilla/mux)
- Orm: [Gorm](https://github.com/jinzhu/gorm) (with [PostgreSQL](https://www.postgresql.org/) persistence)
- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)
- User management

## Frontend

- [Vue.js](https://vuejs.org/) spa client with webpack
- [Vuetify](https://vuetifyjs.com/en/) - light theme

# Todo

- IMPORTANT
  - create new record
  - duplicate record
- overview
  - show in days
  - edit overtime and total working time per week, month
  - compare weekly reported time against nominal total time
- Export plugin to excel and csv

## Fixes

- NEW RECORD
  - IMPORTANT fix wrong from and to dates
  - do not create new record if no consultant is selected
- Edit records
  - do not save the value id ESC is pressed
  - Do not call backend 2x e.g. for updateDescription 

## Improvements business
- Ability to lock last week
- show only available rates per project
- Paginate and sort server-side - using vuetify data table
- Consistency checks
- Add billing evidence
- Export to csv plugin

## Improvements technical

- email confirmation
- letsencrypt tls
- move remembered state from local storage to the backend once user entity available
- administration of a DB - https://getqor.com/en
https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
- refactor routes as on https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
- connect Vue with go - https://juliensalinas.com/en/golang-API-backend-vuejs-SPA-frontend-docker-modern-application/

/* eslint-disable-line no-console */

# To get started

``` bash
# clone repository
go get github.com/valasek/timesheet
cd $GOPATH/src/github.com/valasek/timesheet

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run timesheet.go

# open a new terminal and change to the client dir
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run serve

# build for production with minification
npm run build
```

# License

[MIT license](./LICENSE.md) - see LICENSE for more details