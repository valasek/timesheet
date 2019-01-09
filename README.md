# Timesheet

[![GitHub release](https://img.shields.io/github/release-pre/valasek/timesheet.svg)](https://github.com/valasek/timesheet/releases)
[![GitHub issues](https://img.shields.io/github/issues/valasek/timesheet.svg)](https://github.com/valasek/timesheet/issues)
[![GitHub license](https://img.shields.io/github/license/valasek/timesheet.svg)](https://github.com/valasek/timesheet/blob/master/LICENSE)
[![PayPal Donate](https://img.shields.io/badge/donate-PayPal.me-ff69b4.svg)](https://paypal.me/StanislavValasek)

Self-hosted web application for weekly reporting. Report your consulting hours on projects using selected rates.

# Rationale

Automation is based on a premise, that reporting and billing process includes three separated steps with well-defined data which flows in between:
* **Time reporting** - all company employees
  * covered by this app - eliminates manual errors, shows important data for consultants, data entry is as easy as possible
* **Billing preparation** (confirming/editing reported hours for billing) - project manager or administrative person
  * covered partially by this app. If you bill 1:1 with your reporting then it has all you need
  * data are available in the DB which can be exported by this app read or exported by any DB tool
* **Exporting reported and billed hours** in the internal accounting system for the invoicing to the clients and consultants - salaries
  * not covered by this app
  * data are available in the DB which can be exported by this app read or exported by any DB tool

# Requirements

- Linux, Windows or MacOS
- DB connection to PostgreSQL - configure in timesheet.yaml
- Initialize the DB using command `timesheet db --clean`
- Import initial data in csv format using command `timesheet db --load all`
  - (dates should be entered in ISO format YYYY-MM-DD HH:MM:SS)
  - consultants, projects, rates, holidays, initial reported records - optional

# Screenshots

![Main screen](screenshots/main_screen.png?raw=true "Main screen")

# Download

* [Releases](https://github.com/valasek/timesheet/releases/) (MS Windows 64bit, Linux 64bit, Mac OS X 64bit)

# Usage

```
Web based timesheet application with DB persistence.

Application reads DB and server configuration from config.toml, loads default data if DB is empty and launch web GUI.

Usage:
  timesheet [command]

Available Commands:
  db          Initiate, load or backup DB. See timesheet help db
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

[Go](https://golang.org/), [Gorm](https://gorm.io/), [Vue](https://vuejs.org/), [axios](https://github.com/axios/axios), [Vuetify](https://vuetifyjs.com/en/), [PostgreSQL](https://www.postgresql.org/)

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

Always show full week (then shows month data correctly)
Auto retrieve weekly data from backend when the week is changed
Overview
  - show in month and week correct working and non-working time
  - edit overtime and total working time per week, month
  - compare weekly reported time against nominal total time
Ability to lock last week

## Fixes

- implement rollback in js store if DB fails
- Edit records
  - do not save the value id ESC is pressed

## Improvements business
- Add validations on entered data
- Consistency checks
- Show only available rates per project
- Paginate and sort server-side - using Vuetify data table
- Exports
- Add support for billing - assigning reported hours on projects

## Improvements technical

- email confirmation
- letsencrypt tls
- move remembered state from local storage to the backend once user entity available
- administration of a DB - https://getqor.com/en
https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
- refactor routes as on https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
- connect Vue with go - https://juliensalinas.com/en/golang-API-backend-vuejs-SPA-frontend-docker-modern-application/

# To get started

``` bash
# clone repository
go get github.com/valasek/timesheet
cd $GOPATH/src/github.com/valasek/timesheet

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run timesheet.go db --clean
go run timesheet.go db --load all
go run timesheet.go server

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