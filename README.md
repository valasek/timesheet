[![GitHub release](https://img.shields.io/github/release-pre/valasek/timesheet.svg)](https://github.com/valasek/timesheet/releases)
[![GitHub issues](https://img.shields.io/github/issues/valasek/timesheet.svg)](https://github.com/valasek/timesheet/issues)


| **Linux & Mac & Windows** |
| :-----------------------: |
| [![Build Status](https://travis-ci.org/valasek/timesheet.svg?branch=master)](https://travis-ci.org/valasek/timesheet) |

# Supporting Simple Timesheet

Simple Timesheet is project with its ongoing development made possible entirely by the support of these awesome backers. If you'd like to join them, please consider:

- [Become a backer or sponsor on Patreon](https://www.patreon.com/valasek)
- [One-time donation via PayPal](https://paypal.me/StanislavValasek)

---

# Simple Timesheet

Self-hosted web application for weekly reporting. Report your consulting hours on projects using selected rates.

# Rationale

Automation is based on a premise, that reporting and billing process includes three separated steps with well-defined data which flows in between:
* **Time reporting** - all company employees
  * covered by this app - data are safe, eliminates manual errors, shows important data for consultants, data entry is as easy as possible
* **Billing preparation** (confirming/editing reported hours for billing) - project manager or administrative person
  * covered partially by this app. If you bill 1:1 with your reporting then it has all you need
  * data are available in the DB which can be exported by this app read or exported by any DB tool
* **Exporting reported and billed hours** in the internal accounting system for the invoicing to the clients and consultants - salaries
  * not covered by this app
  * data are available in the DB which can be exported by this app read or exported by any DB tool

# Requirements

- Linux, Windows or MacOS
- DB connection to PostgreSQL
- Initialize the DB via `timesheet db --clean`
- Import data from csv via `timesheet db --load all`
  - mandatory: consultants, projects, rates, holidays
  - optional: reported_records

Simple timesheet can be deployed using Docker (server image size 24.9 MB, DB image size 312 MB)

# Screenshots

![Home](screenshots/home.png?raw=true "Home")

![Reported Overview](screenshots/reported_overview.png?raw=true "Reported Overview")

![State Holidays](screenshots/holidays.png?raw=true "State Holidays")

![Admnistration](screenshots/administration.png?raw=true "Admnistration")

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

## Fixes

- implement rollback in js store if DB fails 
- report error if yaml config file is not present or DB is misconfigured
- have one app version for FE (package.json) and BE (build.bat)
- get favicon on production build

## Improvements business

- Consistency checks
- Show only available rates per project
- Paginate and sort server-side - using Vuetify data table
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
# create DB docker image
sudo docker run --name pg-docker -e POSTGRES_USER=timesheet -e POSTGRES_PASSWORD=timesheet -e POSTGRES_DB=timesheet -d -p 5432:5432 postgres
# alternatively Mount $HOME/docker/volumes/postgres on the host machine to the container side volume path /var/lib/postgresql/data created inside the container. This ensures that postgres data persists even after the container is removed.
sudo docker run --name pg-docker -e POSTGRES_USER=timesheet -e POSTGRES_PASSWORD=timesheet -e POSTGRES_DB=timesheet -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres

# clone the repository
go get github.com/valasek/timesheet
cd $GOPATH/src/github.com/valasek/timesheet

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run timesheet.go db --clean
go run timesheet.go db --load all
go run timesheet.go server

# open a new terminal and change to the client folder
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run serve

# build for production with minification
npm run build

# replace baseUrl in compiled js files
sed -i -e 's/localhost/www.example.com/g' /client/dist/*.js
```

# Create docker images
```
docker save -o ./postgres.tar postgres:latest
docker save -o ./server.tar timesheet_server:latest
copy to target machine
docker load -i <path to image tar file>
```
