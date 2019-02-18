[![GitHub release](https://img.shields.io/github/release-pre/valasek/timesheet.svg)](https://github.com/valasek/timesheet/releases)
[![GitHub issues](https://img.shields.io/github/issues/valasek/timesheet.svg)](https://github.com/valasek/timesheet/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/valasek/timesheet)](https://goreportcard.com/report/github.com/valasek/timesheet)


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

# How to start

- Update Application settings , Vacation settings and Warning limits to meet your needs
- Use Backup & Restore to export demo data. Update csv files and import back your projects, rates, consultants and holidays.
- You are good to go 

# Screenshots

![Home](screenshots/home.png?raw=true "Home")

![Reported Overview](screenshots/reported_overview.png?raw=true "Reported Overview")

![State Holidays](screenshots/holidays.png?raw=true "State Holidays")

![Admnistration](screenshots/administration.png?raw=true "Admnistration")

![Help](screenshots/help.png?raw=true "Help")

## Server Commands

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
cd ./server
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

# Push to heroku
https://devcenter.heroku.com/articles/heroku-cli-commands
https://devcenter.heroku.com/articles/container-registry-and-runtime

```
heroku login
heroku container:login

docker-compose -f "docker-compose.yml" up -d --build
docker build --rm -f "Dockerfile" -t timesheet:latest .
heroku container:run --app=timesheet-cloud timesheet -v

heroku container:rm web
heroku container:push web --app timesheet-cloud
heroku container:release web --app timesheet-cloud

heroku run printenv

heroku ps -a timesheet-cloud

heroku logs --app timesheet-cloud

heroku config
```

# TODO/Contributing

I would love your help! Before submitting a PR, please read over the [Contributing](CONTRIBUTING.md) guide.

Here's a couple of areas that could use some love:

    Caching - Evaluation speed could greatly be improved with the help of caching flags/segments/rules/etc in memory.
    Documentation - Typo? Does something not make sense? Could it be worded better? Please help!
    Examples - More examples on how to use Flipt.
    Test Coverage - Would love to get all coverage over 80%.
    Javascripts - I'm no JS wizz, I'm sure the Javascript code in ui/src could be improved/simplified/tested.


# Pro Version

My plan is to soon start working on a Pro Version of Timesheet for enterprise. Along with support, some of the planned features include:

* User management/login
* Permissions
* HTTPS
* Auditing
* Metrics

If you or your organization would like to help beta test a Pro version of Timesheet, please get in touch with me:

    Twitter: @valasek
    Email: valasek at gmail.com
